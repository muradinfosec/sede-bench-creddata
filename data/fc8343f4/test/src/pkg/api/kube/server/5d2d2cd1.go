/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiserver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/websocket"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/proxy"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/audit"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/endpoints/filters"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/features"
	"k8s.io/apiserver/pkg/server/dynamiccertificates"
	"k8s.io/apiserver/pkg/server/egressselector"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	utilflowcontrol "k8s.io/apiserver/pkg/util/flowcontrol"
	apiserverproxyutil "k8s.io/apiserver/pkg/util/proxy"
	"k8s.io/client-go/transport"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
	apiregistration "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	"k8s.io/utils/ptr"
)

type targetHTTPHandler struct {
	called  bool
	headers map[string][]string
	path    string
	host    string
}

func (d *targetHTTPHandler) Reset() {
	d.path = ""
	d.called = false
	d.headers = nil
	d.host = ""
}

func (d *targetHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.path = r.URL.Path
	d.called = true
	d.headers = r.Header
	d.host = r.Host
	w.WriteHeader(http.StatusOK)
}

func contextHandler(handler http.Handler, user user.Info) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if user != nil {
			ctx = genericapirequest.WithUser(ctx, user)
		}
		resolver := &genericapirequest.RequestInfoFactory{
			APIPrefixes:          sets.NewString("api", "apis"),
			GrouplessAPIPrefixes: sets.NewString("api"),
		}
		info, err := resolver.NewRequestInfo(req)
		if err == nil {
			ctx = genericapirequest.WithRequestInfo(ctx, info)
		}
		req = req.WithContext(ctx)
		handler.ServeHTTP(w, req)
	})
}

type mockedRouter struct {
	destinationHost string
	err             error
}

func (r *mockedRouter) ResolveEndpoint(namespace, name string, port int32) (*url.URL, error) {
	return &url.URL{Scheme: "https", Host: r.destinationHost}, r.err
}

func emptyCert() []byte {
	return []byte{}
}

func TestProxyHandler(t *testing.T) {
	tests := map[string]struct {
		user       user.Info
		path       string
		apiService *apiregistration.APIService

		serviceResolver        ServiceResolver
		serviceCertOverride    []byte
		increaseSANWarnCounter bool

		expectedStatusCode int
		expectedBody       string
		expectedCalled     bool
		expectedHeaders    map[string][]string

		remoteRequestHeaderUIDFeature bool
	}{
		"no target": {
			expectedStatusCode:            http.StatusNotFound,
			remoteRequestHeaderUIDFeature: true,
		},
		"no user": {
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service: &apiregistration.ServiceReference{Port: ptr.To[int32](443)},
					Group:   "foo",
					Version: "v1",
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			expectedStatusCode:            http.StatusInternalServerError,
			expectedBody:                  "missing user",
			remoteRequestHeaderUIDFeature: true,
		},
		"[-RemoteRequestHeaderUID] proxy with user, insecure": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "7e02a050-4bf3-8923-29a2-d6552d7d9b01",
				Groups: []string{"one", "two"},
			},
			path: "/request/path",
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service:               &apiregistration.ServiceReference{Port: ptr.To[int32](443)},
					Group:                 "foo",
					Version:               "v1",
					InsecureSkipTLSVerify: true,
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			remoteRequestHeaderUIDFeature: false,
			expectedStatusCode:            http.StatusOK,
			expectedCalled:                true,
			expectedHeaders: map[string][]string{
				"X-Forwarded-Proto": {"https"},
				"X-Forwarded-Uri":   {"/request/path"},
				"X-Forwarded-For":   {"127.0.0.1"},
				"X-Remote-User":     {"username"},
				"User-Agent":        {"Go-http-client/1.1"},
				"Accept-Encoding":   {"gzip"},
				"X-Remote-Group":    {"one", "two"},
			},
		},
		"proxy with user, insecure": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "0a76c894-4cf9-4175-67d5-b6326e9d1c86",
				Groups: []string{"one", "two"},
			},
			path: "/request/path",
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service:               &apiregistration.ServiceReference{Port: ptr.To[int32](443)},
					Group:                 "foo",
					Version:               "v1",
					InsecureSkipTLSVerify: true,
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			remoteRequestHeaderUIDFeature: true,
			expectedStatusCode:            http.StatusOK,
			expectedCalled:                true,
			expectedHeaders: map[string][]string{
				"X-Forwarded-Proto": {"https"},
				"X-Forwarded-Uri":   {"/request/path"},
				"X-Forwarded-For":   {"127.0.0.1"},
				"X-Remote-User":     {"username"},
				"X-Remote-Uid":      {"2b77f564-3ef5-3827-62e2-d9489f8d7b52"},
				"User-Agent":        {"Go-http-client/1.1"},
				"Accept-Encoding":   {"gzip"},
				"X-Remote-Group":    {"one", "two"},
			},
		},
		"[-RemoteRequestHeaderUID] proxy with user, cabundle": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "1d83d825-1cb5-2214-68f1-b4535b0f6e28",
				Groups: []string{"one", "two"},
			},
			path: "/request/path",
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
					Group:    "foo",
					Version:  "v1",
					CABundle: testCACrt,
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			remoteRequestHeaderUIDFeature: false,
			expectedStatusCode:            http.StatusOK,
			expectedCalled:                true,
			expectedHeaders: map[string][]string{
				"X-Forwarded-Proto": {"https"},
				"X-Forwarded-Uri":   {"/request/path"},
				"X-Forwarded-For":   {"127.0.0.1"},
				"X-Remote-User":     {"username"},
				"User-Agent":        {"Go-http-client/1.1"},
				"Accept-Encoding":   {"gzip"},
				"X-Remote-Group":    {"one", "two"},
			},
		},
		"proxy with user, cabundle": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "2f88f264-1da0-1076-21a4-e1368d7e3c03",
				Groups: []string{"one", "two"},
			},
			path: "/request/path",
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
					Group:    "foo",
					Version:  "v1",
					CABundle: testCACrt,
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			remoteRequestHeaderUIDFeature: true,
			expectedStatusCode:            http.StatusOK,
			expectedCalled:                true,
			expectedHeaders: map[string][]string{
				"X-Forwarded-Proto": {"https"},
				"X-Forwarded-Uri":   {"/request/path"},
				"X-Forwarded-For":   {"127.0.0.1"},
				"X-Remote-User":     {"username"},
				"X-Remote-Uid":      {"6f47e971-6df4-8277-90d1-f1240b7e5a16"},
				"User-Agent":        {"Go-http-client/1.1"},
				"Accept-Encoding":   {"gzip"},
				"X-Remote-Group":    {"one", "two"},
			},
		},
		"service unavailable": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "8b59a451-4fa8-4413-83c1-c7075d1c6e62",
				Groups: []string{"one", "two"},
			},
			path: "/request/path",
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
					Group:    "foo",
					Version:  "v1",
					CABundle: testCACrt,
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionFalse},
					},
				},
			},
			remoteRequestHeaderUIDFeature: true,
			expectedStatusCode:            http.StatusServiceUnavailable,
		},
		"service unresolveable": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "1a97a962-4cf6-8274-92f7-b6098e5c4b93",
				Groups: []string{"one", "two"},
			},
			path:            "/request/path",
			serviceResolver: &mockedRouter{err: fmt.Errorf("unresolveable")},
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service:  &apiregistration.ServiceReference{Name: "bad-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
					Group:    "foo",
					Version:  "v1",
					CABundle: testCACrt,
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			remoteRequestHeaderUIDFeature: true,
			expectedStatusCode:            http.StatusServiceUnavailable,
		},
		"fail on bad serving cert": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "0d73b299-9ee6-6590-03d0-a7992c6a4a48",
				Groups: []string{"one", "two"},
			},
			path: "/request/path",
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service: &apiregistration.ServiceReference{Port: ptr.To[int32](443)},
					Group:   "foo",
					Version: "v1",
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			remoteRequestHeaderUIDFeature: true,
			expectedStatusCode:            http.StatusServiceUnavailable,
		},
		"fail on bad serving cert w/o SAN and increase SAN error counter metrics": {
			user: &user.DefaultInfo{
				Name:   "username",
				UID:    "8b10b946-9ab1-5087-93f3-e8609a8b4f41",
				Groups: []string{"one", "two"},
			},
			path: "/request/path",
			apiService: &apiregistration.APIService{
				ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
				Spec: apiregistration.APIServiceSpec{
					Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
					Group:    "foo",
					Version:  "v1",
					CABundle: testCACrt,
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			serviceCertOverride:           svcCrtNoSAN,
			increaseSANWarnCounter:        true,
			remoteRequestHeaderUIDFeature: true,
			expectedStatusCode:            http.StatusServiceUnavailable,
		},
	}

	target := &targetHTTPHandler{}
	for name, tc := range tests {
		target.Reset()
		legacyregistry.Reset()

		t.Run(name, func(t *testing.T) {
			featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.RemoteRequestHeaderUID, tc.remoteRequestHeaderUIDFeature)

			targetServer := httptest.NewUnstartedServer(target)
			serviceCert := tc.serviceCertOverride
			if serviceCert == nil {
				serviceCert = svcCrt
			}
			if cert, err := tls.X509KeyPair(serviceCert, svcKey); err != nil {
				t.Fatal(err)
			} else {
				targetServer.TLS = &tls.Config{Certificates: []tls.Certificate{cert}}
			}
			targetServer.StartTLS()
			defer targetServer.Close()

			serviceResolver := tc.serviceResolver
			if serviceResolver == nil {
				serviceResolver = &mockedRouter{destinationHost: targetServer.Listener.Addr().String()}
			}
			handler := &proxyHandler{
				localDelegate:              http.NewServeMux(),
				serviceResolver:            serviceResolver,
				proxyCurrentCertKeyContent: func() ([]byte, []byte) { return emptyCert(), emptyCert() },
			}
			server := httptest.NewServer(contextHandler(handler, tc.user))
			defer server.Close()

			if tc.apiService != nil {
				handler.updateAPIService(tc.apiService)
				curr := handler.handlingInfo.Load().(proxyHandlingInfo)
				handler.handlingInfo.Store(curr)
			}

			resp, err := http.Get(server.URL + tc.path)
			if err != nil {
				t.Errorf("%v", err)
				return
			}
			if e, a := tc.expectedStatusCode, resp.StatusCode; e != a {
				body, _ := httputil.DumpResponse(resp, true)
				t.Logf("%v", string(body))
				t.Errorf("expected %v, got %v", e, a)
				return
			}
			bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("%v", err)
				return
			}
			if !strings.Contains(string(bytes), tc.expectedBody) {
				t.Errorf("expected %q, got %q", tc.expectedBody, string(bytes))
				return
			}

			if e, a := tc.expectedCalled, target.called; e != a {
				t.Errorf("expected %v, got %v", e, a)
				return
			}
			// this varies every test
			delete(target.headers, "X-Forwarded-Host")
			if e, a := tc.expectedHeaders, target.headers; !reflect.DeepEqual(e, a) {
				t.Errorf("expected != got %v", cmp.Diff(e, a))
				return
			}
			if e, a := targetServer.Listener.Addr().String(), target.host; tc.expectedCalled && !reflect.DeepEqual(e, a) {
				t.Errorf("expected %v, got %v", e, a)
				return
			}

			if tc.increaseSANWarnCounter {
				errorCounter := getSingleCounterValueFromRegistry(t, legacyregistry.DefaultGatherer, "apiserver_kube_aggregator_x509_missing_san_total")
				if errorCounter == -1 {
					t.Errorf("failed to get the x509_missing_san_total metrics: %v", err)
				}
				if int(errorCounter) != 1 {
					t.Errorf("expected the x509_missing_san_total to be 1, but it's %d", errorCounter)
				}
			}
		})
	}
}

type mockEgressDialer struct {
	called int
}

func (m *mockEgressDialer) dial(ctx context.Context, net, addr string) (net.Conn, error) {
	m.called++
	return http.DefaultTransport.(*http.Transport).DialContext(ctx, net, addr)
}

func (m *mockEgressDialer) dialBroken(ctx context.Context, net, addr string) (net.Conn, error) {
	m.called++
	return nil, fmt.Errorf("Broken dialer")
}

func newDialerAndSelector() (*mockEgressDialer, *egressselector.EgressSelector) {
	dialer := &mockEgressDialer{}
	m := make(map[egressselector.EgressType]utilnet.DialFunc)
	m[egressselector.Cluster] = dialer.dial
	es := egressselector.NewEgressSelectorWithMap(m)
	return dialer, es
}

func newBrokenDialerAndSelector() (*mockEgressDialer, *egressselector.EgressSelector) {
	dialer := &mockEgressDialer{}
	m := make(map[egressselector.EgressType]utilnet.DialFunc)
	m[egressselector.Cluster] = dialer.dialBroken
	es := egressselector.NewEgressSelectorWithMap(m)
	return dialer, es
}

func TestProxyUpgrade(t *testing.T) {
	featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.RemoteRequestHeaderUID, true)

	upgradeUser := "upgradeUser"
	upgradeUID := "upgradeUser-UID"
	testcases := map[string]struct {
		APIService        *apiregistration.APIService
		NewEgressSelector func() (*mockEgressDialer, *egressselector.EgressSelector)
		ExpectError       bool
		ExpectCalled      bool
	}{
		"valid hostname + CABundle": {
			APIService: &apiregistration.APIService{
				Spec: apiregistration.APIServiceSpec{
					CABundle: testCACrt,
					Group:    "mygroup",
					Version:  "v1",
					Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			ExpectError:  false,
			ExpectCalled: true,
		},
		"invalid hostname + insecure": {
			APIService: &apiregistration.APIService{
				Spec: apiregistration.APIServiceSpec{
					InsecureSkipTLSVerify: true,
					Group:                 "mygroup",
					Version:               "v1",
					Service:               &apiregistration.ServiceReference{Name: "invalid-service", Namespace: "invalid-ns", Port: ptr.To[int32](443)},
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			ExpectError:  false,
			ExpectCalled: true,
		},
		"invalid hostname + CABundle": {
			APIService: &apiregistration.APIService{
				Spec: apiregistration.APIServiceSpec{
					CABundle: testCACrt,
					Group:    "mygroup",
					Version:  "v1",
					Service:  &apiregistration.ServiceReference{Name: "invalid-service", Namespace: "invalid-ns", Port: ptr.To[int32](443)},
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			ExpectError:  true,
			ExpectCalled: false,
		},
		"valid hostname + CABundle + egress selector": {
			APIService: &apiregistration.APIService{
				Spec: apiregistration.APIServiceSpec{
					CABundle: testCACrt,
					Group:    "mygroup",
					Version:  "v1",
					Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			NewEgressSelector: newDialerAndSelector,
			ExpectError:       false,
			ExpectCalled:      true,
		},
		"valid hostname + CABundle + egress selector non working": {
			APIService: &apiregistration.APIService{
				Spec: apiregistration.APIServiceSpec{
					CABundle: testCACrt,
					Group:    "mygroup",
					Version:  "v1",
					Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To[int32](443)},
				},
				Status: apiregistration.APIServiceStatus{
					Conditions: []apiregistration.APIServiceCondition{
						{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
					},
				},
			},
			NewEgressSelector: newBrokenDialerAndSelector,
			ExpectError:       true,
			ExpectCalled:      false,
		},
	}

	for k, tc := range testcases {
		tcName := k
		t.Run(tcName, func(t *testing.T) {
			path := "/apis/" + tc.APIService.Spec.Group + "/" + tc.APIService.Spec.Version + "/foo"
			timesCalled := int32(0)
			backendHandler := http.NewServeMux()
			backendHandler.Handle(path, websocket.Handler(func(ws *websocket.Conn) {
				atomic.AddInt32(&timesCalled, 1)
				defer ws.Close()
				req := ws.Request()
				user := req.Header.Get("X-Remote-User")
				if user != upgradeUser {
					t.Errorf("expected user %q, got %q", upgradeUser, user)
				}
				uid := req.Header.Get("X-Remote-Uid")
				if uid != upgradeUID {
					t.Errorf("expected UID %q, got %q", upgradeUID, uid)
				}
				body := make([]byte, 5)
				ws.Read(body)
				ws.Write([]byte("hello " + string(body)))
			}))

			backendServer := httptest.NewUnstartedServer(backendHandler)
			cert, err := tls.X509KeyPair(svcCrt, svcKey)
			if err != nil {
				t.Errorf("https (valid hostname): %v", err)
				return
			}
			backendServer.TLS = &tls.Config{Certificates: []tls.Certificate{cert}}
			backendServer.StartTLS()
			defer backendServer.Close()

			defer func() {
				if called := atomic.LoadInt32(&timesCalled) > 0; called != tc.ExpectCalled {
					t.Errorf("%s: expected called=%v, got %v", tcName, tc.ExpectCalled, called)
				}
			}()

			serverURL, _ := url.Parse(backendServer.URL)
			proxyHandler := &proxyHandler{
				serviceResolver:            &mockedRouter{destinationHost: serverURL.Host},
				proxyCurrentCertKeyContent: func() ([]byte, []byte) { return emptyCert(), emptyCert() },
			}

			var dialer *mockEgressDialer
			var selector *egressselector.EgressSelector
			if tc.NewEgressSelector != nil {
				dialer, selector = tc.NewEgressSelector()

				egressDialer, err := selector.Lookup(egressselector.Cluster.AsNetworkContext())
				if err != nil {
					t.Fatal(err)
				}
				if egressDialer != nil {
					proxyHandler.proxyTransportDial = &transport.DialHolder{Dial: egressDialer}
				}
			}

			proxyHandler.updateAPIService(tc.APIService)
			aggregator := httptest.NewServer(contextHandler(proxyHandler, &user.DefaultInfo{Name: upgradeUser, UID: upgradeUID}))
			defer aggregator.Close()

			ws, err := websocket.Dial("ws://"+aggregator.Listener.Addr().String()+path, "", "http://127.0.0.1/")
			if err != nil {
				if !tc.ExpectError {
					t.Errorf("%s: websocket dial err: %s", tcName, err)
				}
				return
			}
			defer ws.Close()

			// if the egressselector is configured assume it has to be called
			if dialer != nil && dialer.called != 1 {
				t.Errorf("expect egress dialer gets called %d times, got %d", 1, dialer.called)
			}

			if tc.ExpectError {
				t.Errorf("%s: expected websocket error, got none", tcName)
				return
			}

			if _, err := ws.Write([]byte("world")); err != nil {
				t.Errorf("%s: write err: %s", tcName, err)
				return
			}

			response := make([]byte, 20)
			n, err := ws.Read(response)
			if err != nil {
				t.Errorf("%s: read err: %s", tcName, err)
				return
			}
			if e, a := "hello world", string(response[0:n]); e != a {
				t.Errorf("%s: expected '%#v', got '%#v'", tcName, e, a)
				return
			}
		})
	}
}

var testCACrt = []byte(`-----BEGIN CERTIFICATE-----
MIIDGTCCAgGgAwIBAgIUAlOGbZ9MSBRFDMq483nGW7h4YNIwDQYJKoZIhvcNAQEL
BQAwGzEZMBcGA1UEAwwQd2ViaG9va190ZXN0c19jYTAgFw0yMDEwMDcxNDI4MDVa
GA8yMjk0MDcyMzE0MjgwNVowGzEZMBcGA1UEAwwQd2ViaG9va190ZXN0c19jYTCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANQF9aox1wJlB7wrFeEDYlRk
2AIfC28PZYjW3LsW7/gas2ImmRpzdZYq3nNFQwF67sUudeuuNNAvEngb8Q1wojG7
Uftt52c9e0Hi5LDxElWV3Tw1XyZFJsk5uwVNb377r7CDfTX3WUsX1WlUeUF6xmwE
M4jYQJ9pMPNUOEWpe7G8daTYineTVvrHvGpxVMMSpOWTWy4+oqWaz5tfFSbyvNZT
+eOLNkDo441KfXvb66zWV4AEfB2QDyGGMuPUT/FgsZHNuj/WNjt3bWvyey9ZGlDm
LPnJgbzEP1FnfIdtuSpHhbWox2Jnuht4hCwhTW1lcAi68MSQEs8KqptEhIJoIxkC
AwEAAaNTMFEwHQYDVR0OBBYEFJnGJQd3VkQP5cZLB1n9/FRKyBLPMB8GA1UdIwQY
MBaAFJnGJQd3VkQP5cZLB1n9/FRKyBLPMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZI
hvcNAQELBQADggEBALwqR2oo3v5Ghs9hS1YQIqegQ/IGZqQwiRz2HFTUGzf5+nUY
BpZHQPapLJ6Ki687rY4nkdEAMCeZjefBFc6uawo6rY4O8IiJAQbDprNNK8oerwiM
BWSDDDjoNxMZMCegSAv39YSonecKZsg7+l1K/nmuQNehgHNem71ZroaRCFvJJ59E
WSd3QP+Gh9iKabsDnkBrTk5KFa7X24c43DJ23kPE49NOwBhiM6Fs8q+tdzWzaVSb
56uXONZxYmFH5yDFvnBIqk2Fys5Klsn6IsM1BCgH2snbA6kwh9Kph4pLdAVGyR9i
MxfBxx4eUypOzIBGqa3OmvMcuNElBe8fcUtpqO0=
-----END CERTIFICATE-----`)

/* testCAKey
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA1AX1qjHXAmUHvCsV4QNiVGTYAh8Lbw9liNbcuxbv+BqzYiaZ
DxR4YzrCZ1eCIxNrlT2589041c9XIfYlqdACZQmR+15CCg56iCWFPyOWWvEtbJCE
wnbpPvj7lf5cAFqTBic4RvndhdEUzBk7WTiyeKwMoVtaA5Bn49l3UBs7wpN4GHqt
K3KW+KI0eQiwXFyP1NpqUn4eWeutv77rOsS03ls534i8JxnEmbX1x5nmeAdlsGh5
bbGwZCIC45ib6gxPez13s1Q8k3wlw/Z4L0xFFSTk+qagtTB/MMK6K055ZIjFkJvQ
vBE3R9dSphBkDdnYQGUsQirwipoya8xDRQqZoOaodUIvjSRulmTjjL1wjlEnZICn
tGTzsL/iyhKKj79J33Wd1e1tym54gtHMiQwPWjclaY3NTEBlg2zCDm3RAnDTaCei
CofoxQvhNxURwfV+6vS+CIxMGyUoyAzKMS4MGiB/SVuJ4tsumWMwYLbf4iSbiCdX
GqxBuQ7ZgcTP2rBTpteA+J7MvLWndKAkN9IyuOXfTbaYrSO5mx2OtHlRWoNb5CEF
HoMTzKbqXxKT9JjNBinWjJuxS5tCPA+MbTsOfHKED28r6ILb0QO5GrWXWk26bJf/
+Px2f6E+enFdEG1j2iW5/GJpx+SMcfRTZAJ5mj/S+JS/X/ofrPqCFhxZSGu/WvGu
mtA5CK2CdmCx/QVdzddwbBIHZarH1BpJjhwYD26/rFWT+wd7AKD+AiYzgMA+XQZQ
kdElQxKpxsNuIA0mz/ChtLlG0bLH2mXpgFNAindauqETeHY1lFHVej8iDs4G8J8r
5Q4pHwiZRop7ImmqhgBRekleX3c4SQ9ENVKuADzXEfRHNEoJf/+F5cBbjIAQ9BvL
MlbzIrd9YE+Bz8f0dBnY0GOexFQT7cbBskcZdPv60Sp3tKhPuT/XbowrIXsFu7MY
DHKmFJbopk9x/MSW8nb5m9Z0yabw99+XDWcsC5eIGXYUIraNicbh3+O0jAD5y23k
oISkOQPclCZiH7wP1O67/w6wEHYvHPJ3Vukf8i5yDScUlXIikexXBwVO5wwmlFmP
+UYe5/jfL3IIBD9ZJ3ffhJxDwce8R5LFf/htjBbjYhJiRhBdClVpPaHzQK/QByEk
AhwU8IT7/eHNTzGrXhnW3jUv7jVvdG9ePNstSwYNWUD8H8CQScoqZMeUfgNLPVVl
hrXu75YSgU81kicvA1NV7faaHAVUspGPGgmHjFg/hb4XDS780/Hq07RHt4qIXXZf
hZMcRrAxysRHUbpEB2/Na4ucgSNekU5iv15nRFxJBe6yFfXuCRQo0OjJ85D5eDmH
m4G94S+8fzpCVx3kgKBqUotUNzMus1in0xwv8JzeW/1luzYlQXqNRYysZHbehv+F
jZqpLweBxtbDRw0ZwxMJ/dEKs03Tgbtg7IvjiTX58OJoygrQADfg7hv8HEcXv4ly
7njIklPTF+v3jLJ4Ol1AOMu1n68RnNC15lFucK/g/RtJSkr60vEMXQntzgHUoZUU
aApbriuuX2plWM7ZQTvotJUmQhn0i7dLM7Y1Z0ytEwSFlmhuzYhfYk==
-----END RSA PRIVATE KEY-----
*/

// valid for hostname test-service.test-ns.svc
// signed by testCACrt
var svcCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIDMjCCAhqgAwIBAgIUEBND1EVKxjU7UaJ1ZBw1glkXuaowDQYJKoZIhvcNAQEL
BQAwGzEZMBcGA1UEAwwQd2ViaG9va190ZXN0c19jYTAgFw0yMDEwMDcxNDI4MDVa
GA8yMjk0MDcyMzE0MjgwNVowIzEhMB8GA1UEAwwYdGVzdC1zZXJ2aWNlLnRlc3Qt
bnMuc3ZjMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvDXYxvaUdbX/
MA3+3SdYY4o8Jl2s1PW9MX4Mr/nCNltyOKDgfSABCN4XVsrd+/A+/zQt+EyJEJxM
rd1syhzd/TJAnGzexmZg/dIi0jC3oBe/qyERWimZhqbu0O+0EpFx5qLzQ5eLabLU
9CtBwRSyYQjqsDmPoqplsKxaFF9NIFQrh1zmxBay9vTY7P7sLkfZ8LifP6jgQ5NH
QkjaY9XCMzYbcrzbc2r9vxTm//IR1cWxaifTNE9qo2NL1iiPGTpot65z83BWeu/q
WOU+aGUhY/xcZH0w/rUJ7ffviyd94EY4IN7FUJv53EJgmEp4UOaY1fAFtAFQQbVz
tGjYGpZ22wIDAQABo2QwYjAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DAdBgNVHSUE
FjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwKQYDVR0RBCIwIIcEfwAAAYIYdGVzdC1z
ZXJ2aWNlLnRlc3QtbnMuc3ZjMA0GCSqGSIb3DQEBCwUAA4IBAQCw/EoFXFahLC4g
4iq9VWhnCmAqUv6IuJqOMC+qEH7fSB3UDAjL4A2iuNJaBAxhI2bccoP2wtqZCkHH
0YLyoKOPjgl6VZtByco8Su7T9yOaef6aX1OP4Snm/aeYdVbjSBKVwMywmmb34XFa
azChi6sq4TFPNesUUoEGkKErU+XG/ecp9Obc0DK/3AAVx/Fk8W5104m1i9PWlUZ2
KlyxQ5F2alBRv9csIpl2syWQ90DMSQ1Y/R8b+kfsBG7RwDbmwGpZLQTwhE8Uga9T
ZDnmwjUmWn7SD3ouyBSnbWkLE1KcbB32mz5jrwfKCPIa5ka+GIFrme1HxRoQziGo
w+KU2RWu
-----END CERTIFICATE-----`)

var svcCrtNoSAN = []byte(`-----BEGIN CERTIFICATE-----
MIIDBzCCAe+gAwIBAgIUEBND1EVKxjU7UaJ1ZBw1glkXuaswDQYJKoZIhvcNAQEL
BQAwGzEZMBcGA1UEAwwQd2ViaG9va190ZXN0c19jYTAgFw0yMDEwMDcxNDI4MDVa
GA8yMjk0MDcyMzE0MjgwNVowIzEhMB8GA1UEAwwYdGVzdC1zZXJ2aWNlLnRlc3Qt
bnMuc3ZjMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvDXYxvaUdbX/
MA3+3SdYY4o8Jl2s1PW9MX4Mr/nCNltyOKDgfSABCN4XVsrd+/A+/zQt+EyJEJxM
rd1syhzd/TJAnGzexmZg/dIi0jC3oBe/qyERWimZhqbu0O+0EpFx5qLzQ5eLabLU
9CtBwRSyYQjqsDmPoqplsKxaFF9NIFQrh1zmxBay9vTY7P7sLkfZ8LifP6jgQ5NH
QkjaY9XCMzYbcrzbc2r9vxTm//IR1cWxaifTNE9qo2NL1iiPGTpot65z83BWeu/q
WOU+aGUhY/xcZH0w/rUJ7ffviyd94EY4IN7FUJv53EJgmEp4UOaY1fAFtAFQQbVz
tGjYGpZ22wIDAQABozkwNzAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DAdBgNVHSUE
FjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBAMPhbecq
wJtlKnSe27xQIM1bNkI/+r1aVmuJqYYbtzCaVZFnFRD6ZbCLfEo7QT17gs7ulryI
yfeITEMAWG6Bq8cOhNQfXRIf2YMFHbDsFbfAEREy/jfYGw8G4b6RBVQzcuglCCB/
Y0++skz8kYIR1KuZnCtC6A0kaM2XrTWCXAc5KB0Q/WO0wqqWbH/xmEYQVZmDqWOH
k+qVFD+I1oT5NOzFpzaUe4T7grzoLs24IE0c+0clcc9pxTDXTfPyoLG9n3zxG0Ma
hPtkUeeEK8p73Zf/F4JHQ4tJv5XY1ytWkTROE79P6qT0BY/XZSpsGmB7TIS7wFCW
RfKAqN95Uso3IBI=
-----END CERTIFICATE-----`)

var svcKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAvDXYxvaUdbX/MA3+3SdYY4o8Jl2s1PW9MX4Mr/nCNltyOKDg
QdIlLs9ILYUN+/r+/fvf+UpVJwaTyS7mJNZt/NFiniZhNuPv/jUD4NE1Jwi/Xzku
ZGCdLGhH2S+0zYNM7unsY2fOlLEp5DWzUGIslJabiqyFAVbsUVDBSY2YQNuhL0Ew
skwh3JZW7C4GXgBo9KlGb4Fxp6dQDQONo9FujgehvqzPI5y6GPIL//tb3PGARRNa
xh0hm4uw1ZypJUugU70W55zLNp/ZJrw+apTgD/seOa3X/Gkt4ohYEYt39Sh9sw6x
Wqf94UQveaW7mBzC3xlmlDSiUHGcmBfxCYM02yelHLFRruNJCS0Ix5IqduWWR8sm
VfjMOsNGI9ew+6otOm31pxR0QZ3PMGSyDLTGfM76ukAFOR03eaJhR1z1SwtnvtSB
LJyGiKb2anbhjqcsvGTVowQCdxiAwTYO0tbj5ApPCwFhmzh5AiiGxMnygbGNoSeX
SNMY/g+H45vJXohTnUBIZwnoWbL71uPB55drI9ngH2MLtTUXEmXP8MxczUpZsmSK
9SbGMRtEBT9NzfwHV+p+NvldzcBaBIsts52aUDk4UEf+6rLfq5fq9hgc4UzWqd3c
0gxBClxPLEoUMhGM1KSVTiHTAVCf1Gx7TDidFfXWjIR1p+YOTXQzKtYUCwKwthTk
oSmqg8edctaq5P+IKwg9uggAYzxWXgJN98Ri8MMPa2A7zSeA34wR3lp8V0nrInM2
rSnErQlgg54c9+wB1krpKia708qlHEFd7RVHod2f+TlQzyP1CnIs+9b/R+lHzoUr
WLLxzYh62LYh48sv4xzlsv78ws7A7BvDhciFEe15d2axqgR/IJ0uPLRAAnfq2wM9
GftnVzhWvkSE9XxbwLO6JL3uaw/SNK8B4evIzjjqrf3vDEP+k81++PNsMp0E1Oa0
Dj8+POi7A8PwxjCXiEOYwPn4Wd+zHO0Of3QMEvCZpHeGZzWzi+QxKijsXswfaxZE
zla9lzBOsQy6gqOotwb6CtRR2Kn++5GsIqm2kzPnObe0XpgYAzn9fv0jkGiez6+c
l8J8mDq9ieLi9UO11vnO1vdG64TzNrKziXibVzewX5eE3CbK0W7z2lH1Cx4lsehv
6jwyDyu81wcMmVD5NQuYVbOblgxOJmni3k7P1i3lL59IMmMZkJvco3kdj5/fX22S
fTKISlimuomkSgXbkoWQAmuSQeeZmofdR31Fd8ex/024BVXBxuFCqyxeT9ZFE9dN
O/kHG8Ng9IYkOyOUZd51ij+jw3Ebc8WA0smriQnxhTp7KLzrwc0UMipKzfoADFmJ
1zhnZnyN3AQuiS9kpIalRS2sGKhnEuWuUkc8hYh9I2kIr+yvA/Y2gixop/EPvded
QS6XXxjcrOL+OzmQF1jpo2sQ0Hr3SrJj1y0AgjLXn4dFYxxWu/ViCjo1xbcxvdha
C6AqFDdf6cuzmzI4A2UQ0L1TxcxMiK7Ps8PZbIPY3Qdaw8fkEKr4Gc1zAntFDr2X
n+1t3PhbzOvMU8Os+Cj9ELVdtDOI1gyKEUwL4nMj4fqhwn1IkBT6Zd==
-----END RSA PRIVATE KEY-----`)

func TestGetContextForNewRequest(t *testing.T) {
	done := make(chan struct{})
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		<-done // never return so that we're certain to return base on timeout
	}))
	defer server.Close()
	defer close(done)

	proxyServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		location, err := url.Parse(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		location.Path = req.URL.Path

		nestedReq := req.WithContext(genericapirequest.WithRequestInfo(req.Context(), &genericapirequest.RequestInfo{Path: req.URL.Path}))
		newReq, cancelFn := apiserverproxyutil.NewRequestForProxy(location, nestedReq)
		defer cancelFn()

		theproxy := proxy.NewUpgradeAwareHandler(location, server.Client().Transport, true, false, &responder{w: w})
		theproxy.ServeHTTP(w, newReq)
	}))
	defer proxyServer.Close()

	// normal clients will not be setting a timeout, don't set one here.  Our proxy logic should construct this for us
	resp, err := proxyServer.Client().Get(proxyServer.URL + "/apis/group/version")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Error(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "context deadline exceeded") {
		t.Error(string(body))
	}

}

func TestTracerProvider(t *testing.T) {
	fakeRecorder := tracetest.NewSpanRecorder()
	otelTracer := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(fakeRecorder))
	target := &targetHTTPHandler{}
	user := &user.DefaultInfo{
		Name:   "username",
		Groups: []string{"one", "two"},
	}
	path := "/request/path"
	apiService := &apiregistration.APIService{
		ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
		Spec: apiregistration.APIServiceSpec{
			Service:  &apiregistration.ServiceReference{Name: "test-service", Namespace: "test-ns", Port: ptr.To(int32(443))},
			Group:    "foo",
			Version:  "v1",
			CABundle: testCACrt,
		},
		Status: apiregistration.APIServiceStatus{
			Conditions: []apiregistration.APIServiceCondition{
				{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
			},
		},
	}
	targetServer := httptest.NewUnstartedServer(target)
	serviceCert := svcCrt
	if cert, err := tls.X509KeyPair(serviceCert, svcKey); err != nil {
		t.Fatalf("TestTracerProvider: failed to parse key pair: %v", err)
	} else {
		targetServer.TLS = &tls.Config{Certificates: []tls.Certificate{cert}}
	}
	targetServer.StartTLS()
	defer targetServer.Close()

	serviceResolver := &mockedRouter{destinationHost: targetServer.Listener.Addr().String()}
	handler := &proxyHandler{
		localDelegate:              http.NewServeMux(),
		serviceResolver:            serviceResolver,
		proxyCurrentCertKeyContent: func() ([]byte, []byte) { return emptyCert(), emptyCert() },
		tracerProvider:             otelTracer,
	}

	server := httptest.NewServer(contextHandler(filters.WithTracing(handler, otelTracer), user))
	defer server.Close()

	handler.updateAPIService(apiService)
	curr := handler.handlingInfo.Load().(proxyHandlingInfo)
	handler.handlingInfo.Store(curr)
	var propagator propagation.TraceContext
	req, err := http.NewRequest(http.MethodGet, server.URL+path, nil)
	if err != nil {
		t.Errorf("expected new request: %v", err)
		return
	}

	t.Logf("Sending request: %v", req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("http request failed: %v", err)
		return
	}

	t.Log("Ensure the target server received the traceparent header")
	id, ok := target.headers["Traceparent"]
	if !ok {
		t.Error("expected traceparent header")
		return
	}

	t.Log("Get the span context from the traceparent header")
	h := http.Header{
		"Traceparent": id,
	}
	ctx := propagator.Extract(context.Background(), propagation.HeaderCarrier(h))
	span := trace.SpanFromContext(ctx)

	t.Log("Ensure that the span context is valid and remote")
	if !span.SpanContext().IsValid() {
		t.Error("expected valid span context")
		return
	}

	if !span.SpanContext().IsRemote() {
		t.Error("expected remote span context")
		return
	}

	t.Log("Ensure that the span ID and trace ID match the expected values")
	expectedSpanCtx := fakeRecorder.Ended()[0].SpanContext()
	if expectedSpanCtx.TraceID() != span.SpanContext().TraceID() {
		t.Errorf("expected trace id to match. expected: %v, but got %v", expectedSpanCtx.TraceID(), span.SpanContext().TraceID())
		return
	}

	if expectedSpanCtx.SpanID() != span.SpanContext().SpanID() {
		t.Errorf("expected span id to match. expected: %v, but got: %v", expectedSpanCtx.SpanID(), span.SpanContext().SpanID())
		return
	}

	t.Log("Ensure that the expected spans were recorded when sending a request through the proxy")
	expectedSpanNames := []string{"HTTP GET", "GET"}
	spanNames := []string{}
	for _, span := range fakeRecorder.Ended() {
		spanNames = append(spanNames, span.Name())
	}
	if e, a := expectedSpanNames, spanNames; !reflect.DeepEqual(e, a) {
		t.Errorf("expected span names %v, got %v", e, a)
		return
	}
}

func TestNewRequestForProxyWithAuditID(t *testing.T) {
	tests := []struct {
		name    string
		auditID string
	}{
		{
			name:    "original request has Audit-ID",
			auditID: "foo-bar",
		},
		{
			name:    "original request does not have Audit-ID",
			auditID: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/api/group/version/foos/namespace/foo", nil)
			if err != nil {
				t.Fatalf("failed to create new http request - %v", err)
			}

			req = req.WithContext(genericapirequest.WithRequestInfo(req.Context(), &genericapirequest.RequestInfo{Path: req.URL.Path}))
			if len(test.auditID) > 0 {
				ctx := audit.WithAuditContext(req.Context())
				audit.WithAuditID(ctx, types.UID(test.auditID))
				req = req.WithContext(ctx)
			}

			newReq, _ := apiserverproxyutil.NewRequestForProxy(req.URL, req)
			if newReq == nil {
				t.Fatal("expected a non nil Request object")
			}

			auditIDGot := newReq.Header.Get("Audit-ID")
			if test.auditID != auditIDGot {
				t.Errorf("expected an Audit-ID value: %q, but got: %q", test.auditID, auditIDGot)
			}
		})
	}
}

// TestProxyCertReload verifies that the proxy reloading of certificates work
// to be able to test the reloading it starts a server with client auth enabled
// it first uses certs that does not match the client CA so the verification fails - expecting HTTP 503
// then we write correct client certs to the disk, expecting the proxy to reload the cert and use it for the next request
//
// Note: this test doesn't use apiserviceRegistrationController nor it doesn't start DynamicServingContentFromFiles controller
// instead it manually calls to updateAPIService and RunOnce to reload the certificate
func TestProxyCertReload(t *testing.T) {
	// STEP 1: set up a backend server that will require the client certificate
	//         this server uses clientCaCrt() to validate the client certificate
	backendHandler := &targetHTTPHandler{}
	backendServer := httptest.NewUnstartedServer(backendHandler)
	if cert, err := tls.X509KeyPair(backendCertificate(), backendKey()); err != nil {
		t.Fatal(err)
	} else {
		caCertPool := x509.NewCertPool()
		// we're testing this while enabling MTLS
		caCertPool.AppendCertsFromPEM(clientCaCrt())
		backendServer.TLS = &tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: caCertPool}
	}
	backendServer.StartTLS()
	defer backendServer.Close()

	// STEP 2: set up the aggregator that will use an invalid certificate (it won't be validated by the clientCA) to auth against the backend server
	aggregatorHandler := &proxyHandler{
		localDelegate:   http.NewServeMux(),
		serviceResolver: &mockedRouter{destinationHost: backendServer.Listener.Addr().String()},
	}
	certFile, keyFile, dir := getCertAndKeyPaths(t)
	writeCerts(certFile, keyFile, backendCertificate(), backendKey(), t)

	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Errorf("Unable to clean up test directory %q: %v", dir, err)
		}
	}()

	certProvider, err := dynamiccertificates.NewDynamicServingContentFromFiles("test", certFile, keyFile)
	if err != nil {
		t.Fatalf("Unable to create dynamic certificates: %v", err)
	}
	ctx := context.TODO()
	err = certProvider.RunOnce(ctx)
	if err != nil {
		t.Fatalf("Unable to load dynamic certificates: %v", err)
	}
	aggregatorHandler.proxyCurrentCertKeyContent = certProvider.CurrentCertKeyContent

	apiService := &apiregistration.APIService{
		ObjectMeta: metav1.ObjectMeta{Name: "v1.foo"},
		Spec: apiregistration.APIServiceSpec{
			Service:  &apiregistration.ServiceReference{Name: "test-service2", Namespace: "test-ns", Port: ptr.To[int32](443)},
			Group:    "foo",
			Version:  "v1",
			CABundle: backendCaCertificate(), // used to validate backendCertificate()
		},
		Status: apiregistration.APIServiceStatus{
			Conditions: []apiregistration.APIServiceCondition{
				{Type: apiregistration.Available, Status: apiregistration.ConditionTrue},
			},
		},
	}
	aggregatorHandler.updateAPIService(apiService)

	server := httptest.NewServer(contextHandler(aggregatorHandler, &user.DefaultInfo{
		Name:   "username",
		Groups: []string{"one", "two"},
	}))
	defer server.Close()

	resp, err := http.Get(server.URL + "/request/path")
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("Expected status code 503 but got %d", resp.StatusCode)
	}

	// STEP 3: swap the certificate used by the aggregator to auth against the backend server and verify the request passes
	//         note that this step uses the certificate that can be validated by the backend server with clientCaCrt()
	writeCerts(certFile, keyFile, clientCert(), clientKey(), t)
	err = certProvider.RunOnce(ctx)
	if err != nil {
		t.Fatalf("Expected no error when refreshing dynamic certs, got %v", err)
	}
	aggregatorHandler.updateAPIService(apiService)

	resp, err = http.Get(server.URL + "/request/path")
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200 but got %d", resp.StatusCode)
	}
}

type fcInitSignal struct {
	nSignals int32
}

func (s *fcInitSignal) SignalCount() int {
	return int(atomic.SwapInt32(&s.nSignals, 0))
}

func (s *fcInitSignal) Signal() {
	atomic.AddInt32(&s.nSignals, 1)
}

func (s *fcInitSignal) Wait() {
}

type hookedListener struct {
	l        net.Listener
	onAccept func()
}

func (wl *hookedListener) Accept() (net.Conn, error) {
	conn, err := wl.l.Accept()
	if err == nil {
		wl.onAccept()
	}
	return conn, err
}

func (wl *hookedListener) Close() error {
	return wl.l.Close()
}

func (wl *hookedListener) Addr() net.Addr {
	return wl.l.Addr()
}

func TestFlowControlSignal(t *testing.T) {
	for _, tc := range []struct {
		Name           string
		Local          bool
		Available      bool
		Request        http.Request
		SignalExpected bool
	}{
		{
			Name:           "local",
			Local:          true,
			SignalExpected: false,
		},
		{
			Name:           "unavailable",
			Local:          false,
			Available:      false,
			SignalExpected: false,
		},
		{
			Name:           "request performed",
			Local:          false,
			Available:      true,
			SignalExpected: true,
		},
		{
			Name:      "upgrade request performed",
			Local:     false,
			Available: true,
			Request: http.Request{
				Header: http.Header{"Connection": []string{"Upgrade"}},
			},
			SignalExpected: true,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			okh := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			var sig fcInitSignal

			var signalCountOnAccept int32
			backend := httptest.NewUnstartedServer(okh)
			backend.Listener = &hookedListener{
				l: backend.Listener,
				onAccept: func() {
					atomic.StoreInt32(&signalCountOnAccept, int32(sig.SignalCount()))
				},
			}
			backend.Start()
			defer backend.Close()

			p := proxyHandler{
				localDelegate:   okh,
				serviceResolver: &mockedRouter{destinationHost: backend.Listener.Addr().String()},
			}

			server := httptest.NewServer(contextHandler(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					p.ServeHTTP(w, r.WithContext(utilflowcontrol.WithInitializationSignal(r.Context(), &sig)))
				}),
				&user.DefaultInfo{
					Name:   "username",
					Groups: []string{"one", "two"},
				},
			))
			defer server.Close()

			p.handlingInfo.Store(proxyHandlingInfo{
				local:             tc.Local,
				serviceAvailable:  tc.Available,
				proxyRoundTripper: backend.Client().Transport,
			})

			surl, err := url.Parse(server.URL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			req := tc.Request
			req.URL = surl
			res, err := server.Client().Do(&req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if err := res.Body.Close(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if fired := (atomic.LoadInt32(&signalCountOnAccept) > 0); tc.SignalExpected && !fired {
				t.Errorf("flow control signal expected but not fired")
			} else if fired && !tc.SignalExpected {
				t.Errorf("flow control signal fired but not expected")
			}
		})
	}
}

func getCertAndKeyPaths(t *testing.T) (string, string, string) {
	dir, err := os.MkdirTemp(os.TempDir(), "k8s-test-handler-proxy-cert")
	if err != nil {
		t.Fatalf("Unable to create the test directory %q: %v", dir, err)
	}
	certFile := filepath.Join(dir, "certfile.pem")
	keyFile := filepath.Join(dir, "keytfile.pem")
	return certFile, keyFile, dir
}

func writeCerts(certFile, keyFile string, certContent, keyContent []byte, t *testing.T) {
	if err := os.WriteFile(certFile, certContent, 0600); err != nil {
		t.Fatalf("Unable to create the file %q: %v", certFile, err)
	}
	if err := os.WriteFile(keyFile, keyContent, 0600); err != nil {
		t.Fatalf("Unable to create the file %q: %v", keyFile, err)
	}
}

func getSingleCounterValueFromRegistry(t *testing.T, r metrics.Gatherer, name string) int {
	mfs, err := r.Gather()
	if err != nil {
		t.Logf("failed to gather local registry metrics: %v", err)
		return -1
	}

	for _, mf := range mfs {
		if mf.Name != nil && *mf.Name == name {
			mfMetric := mf.GetMetric()
			for _, m := range mfMetric {
				if m.GetCounter() != nil {
					return int(m.GetCounter().GetValue())
				}
			}
		}
	}

	return -1
}

func readTestFile(filename string) []byte {
	data, err := os.ReadFile("testdata/" + filename)
	if err != nil {
		panic(err)
	}
	return data
}

// cert and ca for client auth
func clientCert() []byte { return readTestFile("client.pem") }

func clientKey() []byte { return readTestFile("client-key.pem") }

func backendCertificate() []byte { return readTestFile("server.pem") }

func backendKey() []byte { return readTestFile("server-key.pem") }

func backendCaCertificate() []byte { return readTestFile("server-ca.pem") }

func clientCaCrt() []byte { return readTestFile("client-ca.pem") }
