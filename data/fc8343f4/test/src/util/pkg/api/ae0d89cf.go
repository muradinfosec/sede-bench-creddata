/*
Copyright 2015 The Kubernetes Authors.

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

package spdy

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/armon/go-socks5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"k8s.io/apimachinery/pkg/util/httpstream"
	utilnettesting "k8s.io/apimachinery/pkg/util/net/testing"
)

type serverHandlerConfig struct {
	shouldError      bool
	statusCode       int
	connectionHeader string
	upgradeHeader    string
}

func serverHandler(t *testing.T, config serverHandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if config.shouldError {
			if e, a := httpstream.HeaderUpgrade, req.Header.Get(httpstream.HeaderConnection); e != a {
				t.Fatalf("expected connection=upgrade header, got '%s", a)
			}

			w.Header().Set(httpstream.HeaderConnection, config.connectionHeader)
			w.Header().Set(httpstream.HeaderUpgrade, config.upgradeHeader)
			w.WriteHeader(config.statusCode)

			return
		}

		streamCh := make(chan httpstream.Stream)

		responseUpgrader := NewResponseUpgrader()
		spdyConn := responseUpgrader.UpgradeResponse(w, req, func(s httpstream.Stream, replySent <-chan struct{}) error {
			streamCh <- s
			return nil
		})
		if spdyConn == nil {
			t.Fatal("unexpected nil spdyConn")
		}
		defer spdyConn.Close()

		stream := <-streamCh
		io.Copy(stream, stream)
	}
}

type serverFunc func(http.Handler) *httptest.Server

func httpsServerInvalidHostname(t *testing.T) serverFunc {
	return func(h http.Handler) *httptest.Server {
		cert, err := tls.X509KeyPair(exampleCert, exampleKey)
		if err != nil {
			t.Errorf("https (invalid hostname): proxy_test: %v", err)
		}
		ts := httptest.NewUnstartedServer(h)
		ts.TLS = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
		ts.StartTLS()
		return ts
	}
}

func httpsServerValidHostname(t *testing.T) serverFunc {
	return func(h http.Handler) *httptest.Server {
		cert, err := tls.X509KeyPair(localhostCert, localhostKey)
		if err != nil {
			t.Errorf("https (valid hostname): proxy_test: %v", err)
		}
		ts := httptest.NewUnstartedServer(h)
		ts.TLS = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
		ts.StartTLS()
		return ts
	}
}

func localhostCertPool(t *testing.T) *x509.CertPool {
	localhostPool := x509.NewCertPool()

	if !localhostPool.AppendCertsFromPEM(localhostCert) {
		t.Errorf("error setting up localhostCert pool")
	}
	return localhostPool
}

// be sure to unset environment variable https_proxy (if exported) before testing, otherwise the testing will fail unexpectedly.
func TestRoundTripAndNewConnection(t *testing.T) {
	localhostPool := localhostCertPool(t)

	testCases := map[string]struct {
		serverFunc             func(http.Handler) *httptest.Server
		proxyServerFunc        func(http.Handler) *httptest.Server
		proxyAuth              *url.Userinfo
		clientTLS              *tls.Config
		serverConnectionHeader string
		serverUpgradeHeader    string
		serverStatusCode       int
		shouldError            bool
	}{
		"no headers": {
			serverFunc:             httptest.NewServer,
			serverConnectionHeader: "",
			serverUpgradeHeader:    "",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            true,
		},
		"no upgrade header": {
			serverFunc:             httptest.NewServer,
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            true,
		},
		"no connection header": {
			serverFunc:             httptest.NewServer,
			serverConnectionHeader: "",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            true,
		},
		"no switching protocol status code": {
			serverFunc:             httptest.NewServer,
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusForbidden,
			shouldError:            true,
		},
		"http": {
			serverFunc:             httptest.NewServer,
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"https (invalid hostname + InsecureSkipVerify)": {
			serverFunc:             httpsServerInvalidHostname(t),
			clientTLS:              &tls.Config{InsecureSkipVerify: true},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"https (invalid hostname + hostname verification)": {
			serverFunc:             httpsServerInvalidHostname(t),
			clientTLS:              &tls.Config{InsecureSkipVerify: false},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            true,
		},
		"https (valid hostname + RootCAs)": {
			serverFunc:             httpsServerValidHostname(t),
			clientTLS:              &tls.Config{RootCAs: localhostPool},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied http->http": {
			serverFunc:             httptest.NewServer,
			proxyServerFunc:        httptest.NewServer,
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied https (invalid hostname + InsecureSkipVerify) -> http": {
			serverFunc:             httptest.NewServer,
			proxyServerFunc:        httpsServerInvalidHostname(t),
			clientTLS:              &tls.Config{InsecureSkipVerify: true},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied https with auth (invalid hostname + InsecureSkipVerify) -> http": {
			serverFunc:             httptest.NewServer,
			proxyServerFunc:        httpsServerInvalidHostname(t),
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			clientTLS:              &tls.Config{InsecureSkipVerify: true},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied https (invalid hostname + hostname verification) -> http": {
			serverFunc:             httptest.NewServer,
			proxyServerFunc:        httpsServerInvalidHostname(t),
			clientTLS:              &tls.Config{InsecureSkipVerify: false},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            true, // fails because the client doesn't trust the proxy
		},
		"proxied https (valid hostname + RootCAs) -> http": {
			serverFunc:             httptest.NewServer,
			proxyServerFunc:        httpsServerValidHostname(t),
			clientTLS:              &tls.Config{RootCAs: localhostPool},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied https with auth (valid hostname + RootCAs) -> http": {
			serverFunc:             httptest.NewServer,
			proxyServerFunc:        httpsServerValidHostname(t),
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			clientTLS:              &tls.Config{RootCAs: localhostPool},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied https (invalid hostname + InsecureSkipVerify) -> https (invalid hostname)": {
			serverFunc:             httpsServerInvalidHostname(t),
			proxyServerFunc:        httpsServerInvalidHostname(t),
			clientTLS:              &tls.Config{InsecureSkipVerify: true},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false, // works because the test proxy ignores TLS errors
		},
		"proxied https with auth (invalid hostname + InsecureSkipVerify) -> https (invalid hostname)": {
			serverFunc:             httpsServerInvalidHostname(t),
			proxyServerFunc:        httpsServerInvalidHostname(t),
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			clientTLS:              &tls.Config{InsecureSkipVerify: true},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false, // works because the test proxy ignores TLS errors
		},
		"proxied https (invalid hostname + hostname verification) -> https (invalid hostname)": {
			serverFunc:             httpsServerInvalidHostname(t),
			proxyServerFunc:        httpsServerInvalidHostname(t),
			clientTLS:              &tls.Config{InsecureSkipVerify: false},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            true, // fails because the client doesn't trust the proxy
		},
		"proxied https (valid hostname + RootCAs) -> https (valid hostname + RootCAs)": {
			serverFunc:             httpsServerValidHostname(t),
			proxyServerFunc:        httpsServerValidHostname(t),
			clientTLS:              &tls.Config{RootCAs: localhostPool},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied https with auth (valid hostname + RootCAs) -> https (valid hostname + RootCAs)": {
			serverFunc:             httpsServerValidHostname(t),
			proxyServerFunc:        httpsServerValidHostname(t),
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			clientTLS:              &tls.Config{RootCAs: localhostPool},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied valid https, proxy auth with chars that percent escape -> valid https": {
			serverFunc:             httpsServerValidHostname(t),
			proxyServerFunc:        httpsServerValidHostname(t),
			proxyAuth:              url.UserPassword("proxy user", "proxypasswd%"),
			clientTLS:              &tls.Config{RootCAs: localhostPool},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
	}

	for k, testCase := range testCases {
		t.Run(k, func(t *testing.T) {
			server := testCase.serverFunc(serverHandler(
				t, serverHandlerConfig{
					shouldError:      testCase.shouldError,
					statusCode:       testCase.serverStatusCode,
					connectionHeader: testCase.serverConnectionHeader,
					upgradeHeader:    testCase.serverUpgradeHeader,
				},
			))
			defer server.Close()
			t.Logf("Server URL: %v", server.URL)

			serverURL, err := url.Parse(server.URL)
			if err != nil {
				t.Fatalf("error creating request: %s", err)
			}
			req, err := http.NewRequest("GET", server.URL, nil)
			if err != nil {
				t.Fatalf("error creating request: %s", err)
			}

			spdyTransport, err := NewRoundTripper(testCase.clientTLS)
			if err != nil {
				t.Fatalf("error creating SpdyRoundTripper: %v", err)
			}

			var proxierCalled bool
			var proxyCalledWithHost string
			var proxyCalledWithAuth bool
			var proxyCalledWithAuthHeader string
			if testCase.proxyServerFunc != nil {
				proxyHandler := utilnettesting.NewHTTPProxyHandler(t, func(req *http.Request) bool {
					proxyCalledWithHost = req.Host

					proxyAuthHeaderName := "Proxy-Authorization"
					_, proxyCalledWithAuth = req.Header[proxyAuthHeaderName]
					proxyCalledWithAuthHeader = req.Header.Get(proxyAuthHeaderName)
					return true
				})
				defer proxyHandler.Wait()

				proxy := testCase.proxyServerFunc(proxyHandler)
				defer proxy.Close()

				t.Logf("Proxy URL: %v", proxy.URL)

				spdyTransport.proxier = func(proxierReq *http.Request) (*url.URL, error) {
					proxierCalled = true
					proxyURL, err := url.Parse(proxy.URL)
					if err != nil {
						return nil, err
					}
					proxyURL.User = testCase.proxyAuth
					return proxyURL, nil
				}
			}

			client := &http.Client{Transport: spdyTransport}

			resp, err := client.Do(req)
			var conn httpstream.Connection
			if err == nil {
				conn, err = spdyTransport.NewConnection(resp)
			}
			haveErr := err != nil
			if e, a := testCase.shouldError, haveErr; e != a {
				t.Fatalf("shouldError=%t, got %t: %v", e, a, err)
			}
			if testCase.shouldError {
				return
			}
			defer conn.Close()

			if resp.StatusCode != http.StatusSwitchingProtocols {
				t.Fatalf("expected http 101 switching protocols, got %d", resp.StatusCode)
			}

			stream, err := conn.CreateStream(http.Header{})
			if err != nil {
				t.Fatalf("error creating client stream: %s", err)
			}

			n, err := stream.Write([]byte("hello"))
			if err != nil {
				t.Fatalf("error writing to stream: %s", err)
			}
			if n != 5 {
				t.Fatalf("expected to write 5 bytes, but actually wrote %d", n)
			}

			b := make([]byte, 5)
			n, err = stream.Read(b)
			if err != nil {
				t.Fatalf("error reading from stream: %s", err)
			}
			if n != 5 {
				t.Fatalf("expected to read 5 bytes, but actually read %d", n)
			}
			if e, a := "hello", string(b[0:n]); e != a {
				t.Fatalf("expected '%s', got '%s'", e, a)
			}

			if testCase.proxyServerFunc != nil {
				if !proxierCalled {
					t.Fatal("expected to use a proxy but proxier in SpdyRoundTripper wasn't called")
				}
				if proxyCalledWithHost != serverURL.Host {
					t.Fatalf("expected to see a call to the proxy for backend %q, got %q", serverURL.Host, proxyCalledWithHost)
				}
			}

			if testCase.proxyAuth != nil {
				expectedUsername := testCase.proxyAuth.Username()
				expectedPassword, _ := testCase.proxyAuth.Password()
				username, password, ok := (&http.Request{Header: http.Header{"Authorization": []string{proxyCalledWithAuthHeader}}}).BasicAuth()
				if !ok {
					t.Fatalf("invalid proxy auth header %s", proxyCalledWithAuthHeader)
				}
				if username != expectedUsername || password != expectedPassword {
					t.Fatalf("expected proxy auth \"%s:%s\", got \"%s:%s\"", expectedUsername, expectedPassword, username, password)
				}
			} else if proxyCalledWithAuth {
				t.Fatalf("proxy authorization unexpected, got %q", proxyCalledWithAuthHeader)
			}
		})
	}
}

// Tests SpdyRoundTripper constructors
func TestRoundTripConstuctor(t *testing.T) {
	testCases := map[string]struct {
		tlsConfig         *tls.Config
		proxier           func(req *http.Request) (*url.URL, error)
		upgradeTransport  http.RoundTripper
		expectedTLSConfig *tls.Config
		errMsg            string
	}{
		"Basic TLSConfig; no error": {
			tlsConfig:         &tls.Config{InsecureSkipVerify: true},
			expectedTLSConfig: &tls.Config{InsecureSkipVerify: true},
			upgradeTransport:  nil,
		},
		"Basic TLSConfig and Proxier: no error": {
			tlsConfig:         &tls.Config{InsecureSkipVerify: true},
			proxier:           func(req *http.Request) (*url.URL, error) { return nil, nil },
			expectedTLSConfig: &tls.Config{InsecureSkipVerify: true},
			upgradeTransport:  nil,
		},
		"TLSConfig with UpgradeTransport: error": {
			tlsConfig:         &tls.Config{InsecureSkipVerify: true},
			upgradeTransport:  &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
			expectedTLSConfig: &tls.Config{InsecureSkipVerify: true},
			errMsg:            "SpdyRoundTripper: UpgradeTransport is mutually exclusive to TLSConfig or Proxier",
		},
		"Proxier with UpgradeTransport: error": {
			proxier:           func(req *http.Request) (*url.URL, error) { return nil, nil },
			upgradeTransport:  &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
			expectedTLSConfig: &tls.Config{InsecureSkipVerify: true},
			errMsg:            "SpdyRoundTripper: UpgradeTransport is mutually exclusive to TLSConfig or Proxier",
		},
		"Only UpgradeTransport: no error": {
			upgradeTransport:  &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
			expectedTLSConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			spdyRoundTripper, err := NewRoundTripperWithConfig(
				RoundTripperConfig{
					TLS:              testCase.tlsConfig,
					Proxier:          testCase.proxier,
					UpgradeTransport: testCase.upgradeTransport,
				},
			)
			if testCase.errMsg != "" {
				if err == nil {
					t.Fatalf("expected error but received none")
				}
				if !strings.Contains(err.Error(), testCase.errMsg) {
					t.Fatalf("expected error message (%s), got (%s)", err.Error(), testCase.errMsg)
				}
			}
			if testCase.errMsg == "" {
				if err != nil {
					t.Fatalf("unexpected error received: %v", err)
				}
				actualTLSConfig := spdyRoundTripper.TLSClientConfig()
				if !reflect.DeepEqual(testCase.expectedTLSConfig, actualTLSConfig) {
					t.Errorf("expected TLSConfig (%v), got (%v)",
						testCase.expectedTLSConfig, actualTLSConfig)
				}
			}
		})
	}
}

type Interceptor struct {
	Authorization       socks5.AuthContext
	proxyCalledWithHost *string
}

func (i *Interceptor) GetAuthContext() (int, map[string]string) {
	return int(i.Authorization.Method), i.Authorization.Payload
}

func (i *Interceptor) Rewrite(ctx context.Context, req *socks5.Request) (context.Context, *socks5.AddrSpec) {
	*i.proxyCalledWithHost = req.DestAddr.Address()
	i.Authorization = socks5.AuthContext(*req.AuthContext)
	return ctx, req.DestAddr
}

// be sure to unset environment variable https_proxy (if exported) before testing, otherwise the testing will fail unexpectedly.
func TestRoundTripSocks5AndNewConnection(t *testing.T) {
	localhostPool := localhostCertPool(t)

	socks5Server := func(creds *socks5.StaticCredentials, interceptor *Interceptor) *socks5.Server {
		var conf *socks5.Config
		if creds != nil {
			authenticator := socks5.UserPassAuthenticator{Credentials: creds}
			conf = &socks5.Config{
				AuthMethods: []socks5.Authenticator{authenticator},
				Rewriter:    interceptor,
			}
		} else {
			conf = &socks5.Config{Rewriter: interceptor}
		}

		ts, err := socks5.New(conf)
		if err != nil {
			t.Errorf("failed to create sock5 server: %v", err)
		}
		return ts
	}

	testCases := map[string]struct {
		clientTLS              *tls.Config
		proxyAuth              *url.Userinfo
		serverConnectionHeader string
		serverFunc             serverFunc
		serverStatusCode       int
		serverUpgradeHeader    string
		shouldError            bool
	}{
		"proxied without auth -> http": {
			serverFunc:             httptest.NewServer,
			serverConnectionHeader: "Upgrade",
			serverStatusCode:       http.StatusSwitchingProtocols,
			serverUpgradeHeader:    "SPDY/3.1",
			shouldError:            false,
		},
		"proxied with invalid auth -> http": {
			serverFunc:             httptest.NewServer,
			proxyAuth:              url.UserPassword("invalid", "auth"),
			serverConnectionHeader: "Upgrade",
			serverStatusCode:       http.StatusSwitchingProtocols,
			serverUpgradeHeader:    "SPDY/3.1",
			shouldError:            true,
		},
		"proxied with valid auth -> http": {
			serverFunc:             httptest.NewServer,
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			serverConnectionHeader: "Upgrade",
			serverStatusCode:       http.StatusSwitchingProtocols,
			serverUpgradeHeader:    "SPDY/3.1",
			shouldError:            false,
		},
		"proxied with valid auth -> https (invalid hostname + InsecureSkipVerify)": {
			serverFunc:             httpsServerInvalidHostname(t),
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			clientTLS:              &tls.Config{InsecureSkipVerify: true},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
		"proxied with valid auth -> https (invalid hostname + hostname verification)": {
			serverFunc:             httpsServerInvalidHostname(t),
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			clientTLS:              &tls.Config{InsecureSkipVerify: false},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            true,
		},
		"proxied with valid auth -> https (valid hostname + RootCAs)": {
			serverFunc:             httpsServerValidHostname(t),
			proxyAuth:              url.UserPassword("proxyuser", "proxypasswd"),
			clientTLS:              &tls.Config{RootCAs: localhostPool},
			serverConnectionHeader: "Upgrade",
			serverUpgradeHeader:    "SPDY/3.1",
			serverStatusCode:       http.StatusSwitchingProtocols,
			shouldError:            false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			server := testCase.serverFunc(serverHandler(
				t, serverHandlerConfig{
					shouldError:      testCase.shouldError,
					statusCode:       testCase.serverStatusCode,
					connectionHeader: testCase.serverConnectionHeader,
					upgradeHeader:    testCase.serverUpgradeHeader,
				},
			))
			defer server.Close()

			req, err := http.NewRequest("GET", server.URL, nil)
			if err != nil {
				t.Fatalf("error creating request: %s", err)
			}

			spdyTransport, err := NewRoundTripper(testCase.clientTLS)
			if err != nil {
				t.Fatalf("error creating SpdyRoundTripper: %v", err)
			}
			var proxierCalled bool
			var proxyCalledWithHost string

			interceptor := &Interceptor{proxyCalledWithHost: &proxyCalledWithHost}

			proxyHandler := socks5Server(nil, interceptor)

			if testCase.proxyAuth != nil {
				proxyHandler = socks5Server(&socks5.StaticCredentials{
					"proxyuser": "proxypasswd", // Socks5 server static credentials when client authentication is expected
				}, interceptor)
			}

			closed := make(chan struct{})
			isClosed := func() bool {
				select {
				case <-closed:
					return true
				default:
					return false
				}
			}

			l, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				t.Fatalf("socks5Server: proxy_test: Listen: %v", err)
			}
			defer l.Close()

			go func(shoulderror bool) {
				conn, err := l.Accept()
				if err != nil {
					if isClosed() {
						return
					}

					t.Errorf("error accepting connection: %s", err)
				}

				if err := proxyHandler.ServeConn(conn); err != nil && !shoulderror {
					// If the connection request is closed before the channel is closed
					// the test will fail with a ServeConn error. Since the test only return
					// early if expects shouldError=true, the channel is closed at the end of
					// the test, just before all the deferred connections Close() are executed.
					if isClosed() {
						return
					}

					t.Errorf("ServeConn error: %s", err)
				}
			}(testCase.shouldError)
			spdyTransport.proxier = func(proxierReq *http.Request) (*url.URL, error) {
				proxierCalled = true
				return &url.URL{
					Scheme: "socks5",
					Host:   net.JoinHostPort("127.0.0.1", strconv.Itoa(l.Addr().(*net.TCPAddr).Port)),
					User:   testCase.proxyAuth,
				}, nil
			}

			client := &http.Client{Transport: spdyTransport}

			resp, err := client.Do(req)
			haveErr := err != nil
			if e, a := testCase.shouldError, haveErr; e != a {
				t.Fatalf("shouldError=%t, got %t: %v", e, a, err)
			}
			if testCase.shouldError {
				return
			}

			conn, err := spdyTransport.NewConnection(resp)
			haveErr = err != nil
			if e, a := testCase.shouldError, haveErr; e != a {
				t.Fatalf("shouldError=%t, got %t: %v", e, a, err)
			}
			if testCase.shouldError {
				return
			}

			defer conn.Close()

			if resp.StatusCode != http.StatusSwitchingProtocols {
				t.Fatalf("expected http 101 switching protocols, got %d", resp.StatusCode)
			}

			stream, err := conn.CreateStream(http.Header{})
			if err != nil {
				t.Fatalf("error creating client stream: %s", err)
			}

			n, err := stream.Write([]byte("hello"))
			if err != nil {
				t.Fatalf("error writing to stream: %s", err)
			}
			if n != 5 {
				t.Fatalf("expected to write 5 bytes, but actually wrote %d", n)
			}

			b := make([]byte, 5)
			n, err = stream.Read(b)
			if err != nil {
				t.Fatalf("error reading from stream: %s", err)
			}
			if n != 5 {
				t.Fatalf("expected to read 5 bytes, but actually read %d", n)
			}
			if e, a := "hello", string(b[0:n]); e != a {
				t.Fatalf("expected '%s', got '%s'", e, a)
			}

			if !proxierCalled {
				t.Fatal("xpected to use a proxy but proxier in SpdyRoundTripper wasn't called")
			}

			serverURL, err := url.Parse(server.URL)
			if err != nil {
				t.Fatalf("error creating request: %s", err)
			}
			if proxyCalledWithHost != serverURL.Host {
				t.Fatalf("expected to see a call to the proxy for backend %q, got %q", serverURL.Host, proxyCalledWithHost)
			}

			authMethod, authUser := interceptor.GetAuthContext()

			if testCase.proxyAuth != nil {
				expectedSocks5AuthMethod := 2
				expectedSocks5AuthUser := "proxyuser"

				if expectedSocks5AuthMethod != authMethod {
					t.Fatalf("socks5 Proxy authorization unexpected, got %d, expected %d", authMethod, expectedSocks5AuthMethod)
				}

				if expectedSocks5AuthUser != authUser["Username"] {
					t.Fatalf("socks5 Proxy authorization user unexpected, got %q, expected %q", authUser["Username"], expectedSocks5AuthUser)
				}
			} else {
				if authMethod != 0 {
					t.Fatalf("proxy authentication method unexpected, got %d", authMethod)
				}
				if len(authUser) != 0 {
					t.Fatalf("unexpected proxy user: %v", authUser)
				}
			}

			// The channel must be closed before any of the connections are closed
			close(closed)
		})
	}
}

func TestRoundTripPassesContextToDialer(t *testing.T) {
	urls := []string{"http://127.0.0.1:1233/", "https://127.0.0.1:1233/"}
	for _, u := range urls {
		t.Run(u, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
			require.NoError(t, err)
			spdyTransport, err := NewRoundTripper(&tls.Config{})
			if err != nil {
				t.Fatalf("error creating SpdyRoundTripper: %v", err)
			}
			_, err = spdyTransport.Dial(req)
			assert.EqualError(t, err, "dial tcp 127.0.0.1:1233: operation was canceled")
		})
	}
}

// exampleCert was generated from crypto/tls/generate_cert.go with the following command:
//
//	go run generate_cert.go  --rsa-bits 2048 --host example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var exampleCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDADCCAeigAwIBAgIQVHG3Fn9SdWayyLOZKCW1vzANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYw
MDAwWjASMRAwDgYDVQQKEwdBY21lIENvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEArTCu9fiIclNgDdWHphewM+JW55dCb5yYGlJgCBvwbOx547M9p+tn
zm9QOhsdZDHDZsG9tqnWxE2Nc1HpIJyOlfYsOoonpEoG/Ep6nnK91ngj0bn/JlNy
+i/bwU4r97MOukvnOIQez9/D9jAJaOX2+b8/d4lRz9BsqiwJyg+ynZ5tVVYj7aMi
vXnd6HOnJmtqutOtr3beucJnkd6XbwRkLUcAYATT+ZihOWRbTuKqhCg6zGkJOoUG
f8sX61JjoilxiURA//ftGVbdTCU3DrmGmardp5NNOHbumMYU8Vhmqgx1Bqxb+9he
7G42uW5YWYK/GqJzgVPjjlB2dOGj9KrEWQIDAQABo1AwTjAOBgNVHQ8BAf8EBAMC
AqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zAWBgNVHREE
DzANggtleGFtcGxlLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAig4AIi9xWs1+pLES
eeGGdSDoclplFpcbXANnsYYFyLf+8pcWgVi2bOmb2gXMbHFkB07MA82wRJAUTaA+
2iNXVQMhPCoA7J6ADUbww9doJX2S9HGyArhiV/MhHtE8txzMn2EKNLdhhk3N9rmV
x/qRbWAY1U2z4BpdrAR87Fe81Nlj7h45csW9K+eS+NgXipiNTIfEShKgCFM8EdxL
1WXg7r9AvYV3TNDPWTjLsm1rQzzZQ7Uvcf6deWiNodZd8MOT/BFLclDPTK6cF2Hr
UU4dq6G4kCwMSxWE4cM3HlZ4u1dyIt47VbkP0rtvkBCXx36y+NXYA5lzntchNFZP
uvEQdw==
-----END CERTIFICATE-----`)

var exampleKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEArTCu9fiIclNgDdWHphewM+JW55dCb5yYGlJgCBvwbOx547M9
f+sVSA0iZpMPAuttUrH8zKzmvs6xB2SSIukNDbgqVMjXhvdU/EQ8rOq72KVB1eU/
boYt+A/PCq0R30hXjygNRigIj3/s1jQvGkl6+p4/U5aHH3gvTfRzOl+KxV2qsIJq
6rDIrEtk2FQPIhPxojuev9uhKToLFt1JgOTSftiKYkgE+KKxDUTYBxiVEFw5gHFs
SrrmW8fC07fEveFJJwjK//sITACRInr3sElJwESOu1GNKzLsIOLa5nVcJuI9Dcii
+1bO0x23ZW6eqsV/lTTMbvkwxUD9RDRC2ichKdfMMEIsHdgGrJElm4dMRprv7NtW
bhbxr9eYSptCcS5JDO+grVoJHD5DI5eDb398nNt6dHj4Epk4ZUw7MFRJTi7zCGL8
mkV0XrXOpXvD/cBMDNJXcfNMbxT86n7Fa28i+Iw8EBTs+3K6xpzqiN6lthT436x1
lmJykVmlKEbZPArrYz3yilD10E8iGKK3NBQf/GRDbBAqZn2JQYSS6F1/MxAIegMc
rQdozcBBlwDJWZoPVbj1+eHjmS2GzoAYf3HutHRE1uj/F82azwfrGQVQ3lXPxthe
OVZ88RBlSa5KKuONnpGxPiRW83lH8CJH3pszLHoibnOC72NpJ4wwYiviw5rip/sc
srpDPCFWMzkfuHryUJCd0fQI/MxRwDLUnD4tVH2NXir3I+ZEehKJkmUrTz1+6F6e
gZa0RA8737qHqXclWHM/OqcHp/5pQmtj3C9OtOSvBy5O9hOk974Iw9MHDWxC92qU
gaExRzkoMiNzdUC/GFpG5GEt9NldMlolRyM38mbAckyfMHks8p0/WE6iJaorOL35
JTIDc0f/fdGEEUWFz7KV87NUC3ysm9ctQX/qWKuUTCjF3sRYLFVsABPA2yZLSSsV
p1HnnABiucd8ajKSJFLGejkjLqyF9tLP/Ho635l0VafuiogwAY204SmWjTj/lM0P
k2RhOYRPNvTpttirn4GK09ghzcQrZZ0N5Ch6rvcXZYJiDUUks0nnOFS0AhUt88Tf
cWMoQbDi9/tnYU4ZqXdYvesyNQ9o604WggO2hNbKcWiawjKsziAexmIY8uTSdcIK
CYdD5waSpuILBWFHbVncGaZMrEi3BYdRASLgm0n8eibM2cxQw+TwYgd9KhuzQhAC
4BWo/Gy2bmNz6xnH2FXw7aEXEsvD6gw4/mT16wNoVAr9nXAj2mfWC5ph0NKIs7P9
M/br4xmRAWyYtdOERXVVYLBbXWxj7e9dMVUN0RwR+JcyNLVWa3EGLCM+yIqSxl98
Q3Ih5CHrbSRT5ExExFRz5iipD8thS61ZSnmNQluvjWusnLpZN5sURxv3v8zOiQqx
Dx31p0YXWKBkGYbDnKUbuSp0IDAIo6L1sQGi1qk4RagDGmedWgYmRRaC/kTDgOA3
ZxusBLMi9XkJ1YBxkpjGsSCZmXjp50vjTtn0Zo5GpWM0gzgHr/sT0wwFdNNrx3rX
fD6GRi86YK6gKTOz8/osUPaPYNhiEOZDUoyERt/1YFTz9nPrGwbr8TV=
-----END RSA PRIVATE KEY-----`)

// localhostCert was generated from crypto/tls/generate_cert.go with the following command:
//
//	go run generate_cert.go  --rsa-bits 2048 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var localhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDGTCCAgGgAwIBAgIRALL5AZcefF4kkYV1SEG6YrMwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALQ/FHcyVwdFHxARbbD2KBtDUT7Eni+8ioNdjtGcmtXqBv45EC1C
JOqqGJTroFGJ6Q9kQIZ9FqH5IJR2fOOJD9kOTueG4Vt1JY1rj1Kbpjefu8XleZ5L
SBwIWVnN/lEsEbuKmj7N2gLt5AH3zMZiBI1mg1u9Z5ZZHYbCiTpBrwsq6cTlvR9g
dyo1YkM5hRESCzsrL0aUByoo0qRMD8ZsgANJwgsiO0/M6idbxDwv1BnGwGmRYvOE
Hxpy3v0Jg7GJYrvnpnifJTs4nw91N5X9pXxR7FFzi/6HTYDWRljvTb0w6XciKYAz
bWZ0+cJr5F7wB7ovlbm7HrQIR7z7EIIu2d8CAwEAAaNoMGYwDgYDVR0PAQH/BAQD
AgKkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA8GA1UdEwEB/wQFMAMBAf8wLgYDVR0R
BCcwJYILZXhhbXBsZS5jb22HBH8AAAGHEAAAAAAAAAAAAAAAAAAAAAEwDQYJKoZI
hvcNAQELBQADggEBAFPPWopNEJtIA2VFAQcqN6uJK+JVFOnjGRoCrM6Xgzdm0wxY
XCGjsxY5dl+V7KzdGqu858rCaq5osEBqypBpYAnS9C38VyCDA1vPS1PsN8SYv48z
DyBwj+7R2qar0ADBhnhWxvYO9M72lN/wuCqFKYMeFSnJdQLv3AsrrHe9lYqOa36s
8wxSwVTFTYXBzljPEnSaaJMPqFD8JXaZK1ryJPkO5OsCNQNGtatNiWAf3DcmwHAT
MGYMzP0u4nw47aRz9shB8w+taPKHx2BVwE1m/yp3nHVioOjXqA1fwRQVGclCJSH1
D2iq3hWVHRENgjTjANBPICLo9AZ4JfN6PH19mnU=
-----END CERTIFICATE-----`)

// localhostKey is the private key for localhostCert.
var localhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAtD8UdzJXB0UfEBFtsPYoG0NRPsSeL7yKg12O0Zya1eoG/jkQ
evwx4SAasnrHGmRDU5uahD6fzWwhxzX715HK5B0J00GYE7kMmAKAtyfRN1+9aqY8
jrfglSKJKK3+PkkJO2YnhM7thz5oGAGZwWFkMGvcF58nogHbwWArwHFMoAGxqnc3
t5L5TpPPpLOtTlfrDlxRsDCZbNxTboKDurMNw0AWWoY7E6pEJ2MALB/OfPPdtIQO
84jBwNLI/TfBFKPOO+hUVs2pEOpDw9h1en2aKPLCbvPk/muEMlcitS0FjUcWDXlE
YVHNMAU5aMYPRIUCfD+bisRRQIHBzShDoc6B1GdnYVgQerDLhBvi07DGZLLw0szn
UoOfjXMpYaUPdAHKvsp5SgTZHcbSO0N376RCEPoEujOCvRCMZUtssJ9a+Fe+ogoC
gUBG5bJvb6telu7JUqPFmz4dieU3BrfLzTqE8C62y2GRY73KLUJbpRTCkQAAcSZ3
VYlKJrfmiiP9OKl4AiuHXtpo9mm5Hgq47UcQlE+T+eNTB4a0qzORfxizPGkF/fvh
SvMhlH/dyeA9KHsZ9Ce2e2cZfQpc8xVAz3MXiaOn8vQub7Fj3uof4sbv2+Q8ut52
VG+x5pX5+n+1slcK2MYOppj/LbxmbAhutxw8kywUZHKjUA6F/5arncYjbqosuVR4
sAfe02MksdzPciBTV52PlYq7iRJzDAosO5Eaux9E48qBYR1wbWbsXScNwGllH454
tDVo8+T0JVE8USRDnQjiqigrt7lLjBlLPk/yPPdgdK3zvJTHSf0ajLPJyPVKTmyP
xvOQYx9xCnmyHN39zIutiP9ZI1jIbCgqweVUnHw6z44jdLGiZ1Y2pLMJALNZ0MNX
9BffTvZ1RVy2Y3Qk6DOQ0fn2ZatEi5/5FsOax7Lo2tdHqYAMxD8ff+4FIlcjnj0+
4NpOleVpi7SAlz/MX7uZIq0ko/F29wUaPHShclnEvUalN05Z+glTRljdaUHxgoia
RQ89msSehpBC7ydAVQeJkslPtrHqeNpS2Vmqus0ZeMEWgKHbDmlz6TAyddR5wK3/
c4f2wYbHoPAYNIjq5gb1Zy5oWXScBrswq405EFnkyavdaZKVjvdcVjNMwKq4nIwF
GneN6NIGhuxpdP5e3+95HKCr9PGKVPM2M05woAzw+cCn6xoKqXLJrRJuTDhDGdoj
HREMP1sUcFGr6eTa/DEBDqNnvV9GuPKbG8xtcpZbgKFRejGr+EjK7kWWPQTKMTXk
uCMuWAEgW5G58WCESYVkFuf94JlmBjoeHfXRm4N12xWS3bouYZhXwItuD+Eoq6E4
RsyLhTc5hCE7Rs/TEvynHspWZdPG94QfewzaN8OCJ3DteJ78VrQEEn9u78AEiNkq
7S7TbLcbZs5jz3bA8ZXshiBYrhCDbesJTn2VM4/Bes66EKBvks4M+RPb1dxowdVA
9UekC5e7B/SccZ9oniSUrBoXCDpmqKFE0x1KFIwTHKlL0b239xXC8xPBZ4Fnxnr1
kx7HVntH51re/fVM7WhBRF92lpAr4Vw3iWICdw2O3jnfuMTQD9z=
-----END RSA PRIVATE KEY-----
`)
