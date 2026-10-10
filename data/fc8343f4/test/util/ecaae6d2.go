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

package utils

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"k8s.io/api/admission/v1beta1"
)

// NewAdmissionWebhookServer sets up a webhook server with TLS enabled, returns URL and Close function
// for the server
func NewAdmissionWebhookServer(handler http.Handler) (string, func(), error) {
	// set up webhook server
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(LocalhostCert) {
		return "", nil, fmt.Errorf("Failed to append Cert from PEM")
	}
	cert, err := tls.X509KeyPair(LocalhostCert, LocalhostKey)
	if err != nil {
		return "", nil, fmt.Errorf("Failed to build cert with error: %+v", err)
	}
	webhookServer := httptest.NewUnstartedServer(handler)
	webhookServer.TLS = &tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{cert},
	}
	webhookServer.StartTLS()
	return webhookServer.URL, webhookServer.Close, nil
}

// AdmissionWebhookHandler creates a HandlerFunc that decodes/encodes AdmissionReview and performs
// given admit function
func AdmissionWebhookHandler(t *testing.T, admit func(*v1beta1.AdmissionReview) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		review := v1beta1.AdmissionReview{}
		if err := json.Unmarshal(data, &review); err != nil {
			t.Errorf("Fail to deserialize object: %s with error: %v", string(data), err)
			http.Error(w, err.Error(), 400)
			return
		}

		if err := admit(&review); err != nil {
			t.Errorf("%v", err)
			http.Error(w, err.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(review); err != nil {
			t.Errorf("Marshal of response failed with error: %v", err)
		}
	}
}

// LocalhostCert was generated from crypto/tls/generate_cert.go with the following command:
//
//	go run generate_cert.go  --rsa-bits 2048 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
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

// LocalhostKey is the private key for LocalhostCert.
var LocalhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA1Z5/aTwqY706M34tn60l8ZHkanWDl8mM1pYf4Q7qg3zA9XqW
GE8G3YooALxT1vGmbUb64xWLwvGSUWCpPy04gmAA5mGmxCqop1xTLrVLE0CXJBvU
xkO+4ypVuQ3vI5CmJMtNeWWQZBE8+xveQ+bvklsITLvJGjr5lkdiLBr/NAIbRlxw
z9s9qljQm3qk01QtwT79fOXXenEt3WhoJhv2C5rD7IKQtHesUVQ9A/O6Euh7RebO
CE/FZner0yJXY0eOjdjqEWbdnKaEZM+vDFcIxz1Y8PogK5jk8BBUD1APur7r4id+
4UYwa8I20B9xT+ujstoogAfIL1hvsuchO1rERwqWdMBOnQbkOoyjsW2pdrm6ORpM
mYmmXCEngNO6YnSuCPclD8j5R9ltB0Y7F0o+g8Hwh3rOe7pdwQMnDeCrku21jBHD
FjFfSiZU79cPnhMkdWO0sCtjDiEVY01k/mzse3Rsk4kOPv+gXQ/tcIQgXgieOBse
a54HOEnD3x30by9GHisqb1CtBZfEWm3wPtVY9tiP8/49HNBjUk7/MV/NqAs78vxS
ei9IrebzYClpNUYMJ2310xGw13BnsbH1LvGFNc/zeaZr5A8Ytrb1Swa5Ew6KpwcV
fwEriOvyC4N0nCuct0FGpdgApiHB925N231/RqbZxbXVLehG73N1rkkI+qOdwUyW
nwgK8z8yZDxJ1BQt4eOHJ/hRIVwFP+xW51bovMtHARgZjdyJGwT2sH0K790MY7me
1wjYxKofzQmFuckyQHXmreBJkDMYJhJqcNRGAAtHldB/cRB1qAL7KoMMeKqT77Ll
1bl7wBJau2cmJ/bRS0y2cuRsUxWHuK1CYbOjbPZ3HpWxFZDsTA9Y2g9weOKJ1zQK
OqVkMGrZJXtSCn8jc1JFnvXgHjU3F7jBYrOetMsO4z8umrI038szSK30WooX7e/3
Fya7f7em7saRsCcH3Pur8jFMeJe7nMWgrNcyBmL3y6eW5uca5+rr+9Iing87upK5
JpORzADVYTE7FTAd8c8gNL/xvLbE72t/V56KO/DlHfOO63JWR2QEBAD/4YXVvZrs
GeU1uY9Z+RnnmA1OcLHtVmO7hqmaSZLXSVee8Q3qVO3jZSywx1KN34WC1+YY0XJE
tf3fG1oniYGm7k+LaiawlkDiG4+TknLo0Yn/nzKxVeFgwAX2//lYhgvUYV2lOvkr
Hxw+8qzWavwb40gkhYipKWQyzwQiiFKZd3IHBzzRIy525uzUdj8l02JRos3qWi7W
p9rwIz4BombsRYfEp4JtaAqyO4mROTf5eCeuvmqUhDx0KRC8n6TgFnvUYCwE4ldK
PDcjZtS4aV6/ypaqtHBURU2PenzZR0skTaYRiZKDoJMilfh53Q+9lRFgZAv4ehA7
EJAsfcPSWSiYh29xzHQiJnscD9RMRtbL3WGcxxtt6KXbfRidGiJ9e8wpPnRf+Vav
TqGHDxpmmlU+slIQhDfAJ3vBOsM3E8Kn9kjC2oWjN1XcG+ceR2L4NFz4HMY9ushs
Ot+qWIzsI4WxflC9T1i5kkn8gLjttjhWMS81P85X5lT4vuKr4b4a
-----END RSA PRIVATE KEY-----`)
