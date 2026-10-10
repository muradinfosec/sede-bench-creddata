/*
Copyright 2019 The Kubernetes Authors.

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

package conversion

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// StartConversionWebhookServer starts an http server with the provided handler and returns the WebhookClientConfig
// needed to configure a CRD to use this conversion webhook as its converter.
func StartConversionWebhookServer(handler http.Handler) (func(), *apiextensionsv1.WebhookClientConfig, error) {
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(localhostCert) {
		return nil, nil, fmt.Errorf("failed to append Cert from PEM")
	}
	cert, err := tls.X509KeyPair(localhostCert, localhostKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build cert with error: %+v", err)
	}

	webhookMux := http.NewServeMux()
	webhookMux.Handle("/convert", handler)
	webhookServer := httptest.NewUnstartedServer(webhookMux)
	webhookServer.TLS = &tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{cert},
	}
	webhookServer.StartTLS()
	endpoint := webhookServer.URL + "/convert"
	webhookConfig := &apiextensionsv1.WebhookClientConfig{
		CABundle: localhostCert,
		URL:      &endpoint,
	}

	// StartTLS returns immediately, there is a small chance of a race to avoid.
	if err := wait.PollUntilContextTimeout(context.Background(), time.Millisecond*100, wait.ForeverTestTimeout, true, func(ctx context.Context) (done bool, err error) {
		_, getErr := webhookServer.Client().Get(webhookServer.URL) // even a 404 is fine
		return getErr == nil, nil
	}); err != nil {
		webhookServer.Close()
		return nil, nil, err
	}

	return webhookServer.Close, webhookConfig, nil
}

// V1Beta1ReviewConverterFunc converts an entire ConversionReview.
type V1Beta1ReviewConverterFunc func(review *apiextensionsv1beta1.ConversionReview) (*apiextensionsv1beta1.ConversionReview, error)

// V1ReviewConverterFunc converts an entire ConversionReview.
type V1ReviewConverterFunc func(review *apiextensionsv1.ConversionReview) (*apiextensionsv1.ConversionReview, error)

// NewReviewWebhookHandler creates a handler that delegates the review conversion to the provided ReviewConverterFunc.
func NewReviewWebhookHandler(t *testing.T, v1beta1ConverterFunc V1Beta1ReviewConverterFunc, v1ConverterFunc V1ReviewConverterFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			t.Errorf("contentType=%s, expect application/json", contentType)
			return
		}

		typeMeta := &metav1.TypeMeta{}
		if err := json.Unmarshal(data, typeMeta); err != nil {
			t.Errorf("Fail to deserialize object: %s with error: %v", string(data), err)
			http.Error(w, err.Error(), 400)
			return
		}

		var response runtime.Object

		switch typeMeta.GroupVersionKind() {
		case apiextensionsv1.SchemeGroupVersion.WithKind("ConversionReview"):
			review := &apiextensionsv1.ConversionReview{}
			if err := json.Unmarshal(data, review); err != nil {
				t.Errorf("Fail to deserialize object: %s with error: %v", string(data), err)
				http.Error(w, err.Error(), 400)
				return
			}

			if v1ConverterFunc == nil {
				http.Error(w, "Cannot handle v1 ConversionReview", 422)
				return
			}
			response, err = v1ConverterFunc(review)
			if err != nil {
				t.Errorf("Error converting review: %v", err)
				http.Error(w, err.Error(), 500)
				return
			}

		case apiextensionsv1beta1.SchemeGroupVersion.WithKind("ConversionReview"):
			review := &apiextensionsv1beta1.ConversionReview{}
			if err := json.Unmarshal(data, review); err != nil {
				t.Errorf("Fail to deserialize object: %s with error: %v", string(data), err)
				http.Error(w, err.Error(), 400)
				return
			}

			if v1beta1ConverterFunc == nil {
				http.Error(w, "Cannot handle v1beta1 ConversionReview", 422)
				return
			}
			response, err = v1beta1ConverterFunc(review)
			if err != nil {
				t.Errorf("Error converting review: %v", err)
				http.Error(w, err.Error(), 500)
				return
			}

		default:
			err := fmt.Errorf("unrecognized request kind: %v", typeMeta.GroupVersionKind())
			t.Error(err)
			http.Error(w, err.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("Marshal of response failed with error: %v", err)
		}
	})
}

// ObjectConverterFunc converts a single custom resource to the desiredAPIVersion and returns it or returns an error.
type ObjectConverterFunc func(desiredAPIVersion string, customResource runtime.RawExtension) (runtime.RawExtension, error)

// NewObjectConverterWebhookHandler creates a handler that delegates custom resource conversion to the provided ConverterFunc.
func NewObjectConverterWebhookHandler(t *testing.T, converterFunc ObjectConverterFunc) http.Handler {
	return NewReviewWebhookHandler(t, func(review *apiextensionsv1beta1.ConversionReview) (*apiextensionsv1beta1.ConversionReview, error) {
		converted := []runtime.RawExtension{}
		errMsgs := []string{}
		for _, obj := range review.Request.Objects {
			convertedObj, err := converterFunc(review.Request.DesiredAPIVersion, obj)
			if err != nil {
				errMsgs = append(errMsgs, err.Error())
			}

			converted = append(converted, convertedObj)
		}

		review.Response = &apiextensionsv1beta1.ConversionResponse{
			UID:              review.Request.UID,
			ConvertedObjects: converted,
		}
		if len(errMsgs) == 0 {
			review.Response.Result = metav1.Status{Status: "Success"}
		} else {
			review.Response.Result = metav1.Status{Status: "Failure", Message: strings.Join(errMsgs, ", ")}
		}
		return review, nil
	}, func(review *apiextensionsv1.ConversionReview) (*apiextensionsv1.ConversionReview, error) {
		converted := []runtime.RawExtension{}
		errMsgs := []string{}
		for _, obj := range review.Request.Objects {
			convertedObj, err := converterFunc(review.Request.DesiredAPIVersion, obj)
			if err != nil {
				errMsgs = append(errMsgs, err.Error())
			}

			converted = append(converted, convertedObj)
		}

		review.Response = &apiextensionsv1.ConversionResponse{
			UID:              review.Request.UID,
			ConvertedObjects: converted,
		}
		if len(errMsgs) == 0 {
			review.Response.Result = metav1.Status{Status: "Success"}
		} else {
			review.Response.Result = metav1.Status{Status: "Failure", Message: strings.Join(errMsgs, ", ")}
		}
		return review, nil
	})
}

// localhostCert was generated from crypto/tls/generate_cert.go with the following command:
//
//	go run generate_cert.go  --rsa-bits 2048 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var localhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDGDCCAgCgAwIBAgIQTKCKn99d5HhQVCLln2Q+eTANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYw
MDAwWjASMRAwDgYDVQQKEwdBY21lIENvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEA1Z5/aTwqY706M34tn60l8ZHkanWDl8mM1pYf4Q7qg3zA9XqWLX6S
4rTYDYCb4stEasC72lQnbEWHbthiQE76zubP8WOFHdvGR3mjAvHWz4FxvLOTheZ+
3iDUrl6Aj9UIsYqzmpBJAoY4+vGGf+xHvuukHrVcFqR9ZuBdZuJ/HbbjUyuNr3X9
erNIr5Ha17gVzf17SNbYgNrX9gbCeEB8Z9Ox7dVuJhLDkpF0T/B5Zld3BjyUVY/T
cukU4dTVp6isbWPvCMRCZCCOpb+qIhxEjJ0n6tnPt8nf9lvDl4SWMl6X1bH+2EFa
a8R06G0QI+XhwPyjXUyCR8QEOZPCR5wyqQIDAQABo2gwZjAOBgNVHQ8BAf8EBAMC
AqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zAuBgNVHREE
JzAlggtleGFtcGxlLmNvbYcEfwAAAYcQAAAAAAAAAAAAAAAAAAAAATANBgkqhkiG
9w0BAQsFAAOCAQEAThqgJ/AFqaANsOp48lojDZfZBFxJQ3A4zfR/MgggUoQ9cP3V
rxuKAFWQjze1EZc7J9iO1WvH98lOGVNRY/t2VIrVoSsBiALP86Eew9WucP60tbv2
8/zsBDSfEo9Wl+Q/gwdEh8dgciUKROvCm76EgAwPGicMAgRsxXgwXHhS5e8nnbIE
Ewaqvb5dY++6kh0Oz+adtNT5OqOwXTIRI67WuEe6/B3Z4LNVPQDIj7ZUJGNw8e6L
F4nkUthwlKx4yEJHZBRuFPnO7Z81jNKuwL276+mczRH7piI6z9uyMV/JbEsOIxyL
W6CzB7pZ9Nj1YLpgzc1r6oONHLokMJJIz/IvkQ==
-----END CERTIFICATE-----`)

// localhostKey is the private key for localhostCert.
var localhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA1Z5/aTwqY706M34tn60l8ZHkanWDl8mM1pYf4Q7qg3zA9XqW
zg5j5zlBYVii6wstJPd77YloVGyMkUrQtG36DUyT4gucsGmbT4tFueRyP3xKhWtM
Efy+1GXsEr4vx4etufbGjoEZSAk6+zzzC+QflfmmyZGNNts1wovorCZ/kZfnusLh
i3f2Onkgg8vn39gHgr73fDItZomU1kwFTrB7Q1Cs7sTbfoKPqKQ0V/o0Ixg5pZEH
Yp/xqhrE3qKQT1aMSPIJATHwtldhXt+oNMiSRf0k2SZRs8fA2FvHR9zQLU6O7Ue+
3dzAJ3t70u8jo+eSKekhAdVgk7tXjHHhW1swLXNxbAaHQemdndMryq7nilm2iBhV
CeOPSRkYjsJ9eCzlyRvTa6p0J0qns4S5x1r+u2pka1gUs1JJdhRGUqzbAe62gOCa
paSoVhNp48nEoYzWQxp1lMQXidKpH63n/ZSLs8WaH6VVyi+oZE/cTKSMPYvMtmMJ
T72cgdoK9F34yr5yAifaV0qCFiYcCb4CozMb4wOH9/69BvXxgi0/RH/wOFI72ELS
xW0bBfNmYVJTrAHsY3940DeR72WPMAg1JeYAxw/CTwJv1z1xIAN4jfO0jt3Caiut
xfZTeWuDu6s3mhECK9kEQISuJGYd382J851/wekushlcOERW51z7aYRS+LwOEXiF
ikcV4G5NGezn7NBm4AoBk/biynMuy+aN92rzBrkIbROqVFXFIPU6zh4l905CB8Ht
9OVTXsbXcfFhTcSVvOHXWeqdQVjjKCPvIGwNEMJFfrt/Pjb0SZJ2nCnJAxHN41Yh
1eq8Batkf1sVq/bhz5w6ujjbAjYLvS3QQZZpiKV8lfrYCjLnHE0W8W7gQtVt8PHn
MIbhCObUFMUoxF9xR9dgxUSagfP4g7ScHTwlowlj7O3cyzA568sAQJ95qrOp2T/8
qVC2Z5YO0qUQYOzb5LaT8ZhcPqH9WLqqrTwrsuR6T7Zi3GVc1+lF+1oMIy79QOC2
RozYCNlkcRv4kCQb5v4hyG/DUtEH23a/H95xW/wfMtKE52fLB7PylVy/8bedPVgX
okT2Yo6T+rOIcc6AzEaPOvb1qgtRkUJSsfDk9W2AEd5sZWyGV0pT82Ao3+UT6moL
ZP8dz0LDbFFd5N+mVavOSYruW2+aDKwv7uD/QlyMREGEJnN1//MCdcYPnU8jCgGX
qNZ+9olntjQO11bsHqIgGKOXwYAHMYbrT7kAjEEFKi775EDCAk9l85zyJf9iuw5g
Z9uiSy8ZwsLFFjBGL6koFNwQW7Hpayb0yWCWuqfkbcE8xrT8c2QqpJyAvyLY0qBg
GQdjlpY5aL1/piNXrnPaSZ6pwfZkv0ZCEPUVraqsuRxKSkw35x+4DhdoKIQ2WVh0
IUuyTXnuXExHZ77rMaJgQWUEj8AbnucJ8GMDwZQn2iUSXHpWLNy0H5PsJzvq+ULY
XiaVXdlxXUU+uEJdaQrvX3dzlpm6d3vj3AHu2NbEh4Zto+MRm6p0HAn3eig4wmOj
sa+SbFZZr0CFYLD3p1I8xei8EzkNFVEIva38j07O0ud0shPD2n6R
-----END RSA PRIVATE KEY-----`)
