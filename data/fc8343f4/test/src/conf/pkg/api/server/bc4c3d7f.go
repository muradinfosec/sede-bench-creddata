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

package dynamiccertificates

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/util/dump"
)

var serverKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA13f50PPWuR/InxLIoJjHdNSG+jVUd25CY7ZL2J023X2BAY+1
D4CFMp4x3pVHSsn77AVzG87/F8w/UM3MaR310Wh598J615d+uNnTNB8pXJiv1T6S
e80lWw2UHJ/+5UF3JkyLWGc6W4cu3hFziSXk37nK8eXOQd0a1e53ZzRUsfRWtXzs
QFWFipZNU0Q5fGZRtfYUBwUBzAdNBqeq3WyeUraYrvz3GW23kGgZ54E46BR0Y1M9
HPqd/vTpLW2XJDWCSXxwo540QHSvGUUjf5dAtyw3t41H4UhR0WJdXJTedzgkG4aw
58S1Q7Blpz/kS0e8ZymdgXb8ipBmPtTxAVXLTqhpeVbticxWvwPadPJjeq7b5D+Z
OVOQH6ul98hYaQTVHr1D++WhJZBqWBypNCdcWEyKd6tzAnjgbY0Lpa0uP/OTxnh8
O1x+2ENCwze4f/RK2BQQa4mmgRStIGLvowh31iCasdd8We6jx3tK+5VyNNO6YezF
euY55ITA8+oJiSWtzABVXZ87t1eGh6Xslf+seUeHbvgz0fVwicbj4zeBDkXa7mAI
qHXF1IuFb9yy/xhg6fDgWUcjqgDubdG+iRfEcifQLPm5lSMWou+mABWt07U/FmJz
QopNmqT1OxIr6+9hmZOhy+R3EaVj0uIB059Zmhbld27OlhwgTqkrrJzYwzx05A3y
Pd3f17PqwbAT3ZTB9cp9uB6w6szH9Ua4jY01pBa0p7rdER1qDUHNbiICW6NaoaXN
QZ/6jXFypMRuDuXKN/fbkNmlzlb1t+fX/mzOtPh34ZG6D5x24HCUYhc5ROElTZoI
iBSmQ7/cjdGmoUdwLOfhm/gYnMJ7nPcU4dZ+Sfi2bGSa5QiRt8f3Cz8nBBJG3U6X
gOpRMG3PB7q5e9s0FIa8YOyLfVXUopLxV8lp9tMXHuUfD9dpFR6qfFDJbYb/3HMG
PDYlc8HDXAhQX27Fqa9uMbd9RY/kiSe3TWpZwjtwX5n3vRxm18G01sAyEz1Gv8A9
YVPMrngahtYf8E+8LOGRuqZCVwzVzYahc+itLqJQxrslpvqnpHjBmdP8xCsW+JEP
5PwAhFUlJzPFuI84uFWcWGX58D2xzDySUDeZhsIpBxznVaPCCfUVLmMq1+5E5lpA
nvben5xA7xUhs3jypRaeHFKmJbzABhENaK4ZJtFQZ7U6mxrXoryy+NU6wiiMDQEe
FQOR2rQBBC3H7LnuN2TjoQUmSJaln3y4kTOyF+4ENAJkvm+GifOxfU1UNwYa7ZE5
LUSX7vO9sSMfCgniODToA5MT/rylx0rBKpiCMSj87hEPZnp8iStWoz11Lf4SOmdU
JU5SbbUUZeFYBqcNAPAhHSA6awT1V5IoKqKAv1RSCUJ6pdPA9foKINLKQ+OVVQZA
ydRPYKELJUtHskQNdt34ot/EIiuj7aIBmdkizPf6btDE4WF10yJfefTkbBG7Ef1R
YGLqCXfOvu/PnRbCANML+powbPD/DcoeVbKrgpmzbsCtOtuwSAWbDKU2qky/iz0Z
gC4KcZkWSQLEnuE6DfpgxDzoFnGMHkEa62CmYeShNa0Yxvclbvt1
-----END RSA PRIVATE KEY-----`)

var serverCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDQDCCAiigAwIBAgIJANWw74P5KJk2MA0GCSqGSIb3DQEBCwUAMDQxMjAwBgNV
BAMMKWdlbmVyaWNfd2ViaG9va19hZG1pc3Npb25fcGx1Z2luX3Rlc3RzX2NhMCAX
DTE3MTExNjAwMDUzOVoYDzIyOTEwOTAxMDAwNTM5WjAjMSEwHwYDVQQDExh3ZWJo
b29rLXRlc3QuZGVmYXVsdC5zdmMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQDXd/nQ89a5H8ifEsigmMd01Ib6NVR3bkJjtkvYnTbdfYEBj7UzqOQtHoLa
dIVmefny5uIHvj93WD8WDVPB3jX2JHrXkDTXd/6o6jIXHcsUfFTVLp6/bZ+Anqe0
r/7hAPkzA2A7APyTWM3ZbEeo1afXogXhOJ1u/wz0DflgcB21gNho4kKTONXO3NHD
XLpspFqSkxfEfKVDJaYAoMnYZJtFNsa2OvsmLnhYF8bjeT3i07lfwrhUZvP+7Gsp
7UgUwc06WuNHjfx1s5e6ySzH0QioMD1rjYneqOvk0pKrMIhuAEWXqq7jlXcDtx1E
j+wnYbVqqVYheHZ8BCJoVAAQGs9/AgMBAAGjZDBiMAkGA1UdEwQCMAAwCwYDVR0P
BAQDAgXgMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATApBgNVHREEIjAg
hwR/AAABghh3ZWJob29rLXRlc3QuZGVmYXVsdC5zdmMwDQYJKoZIhvcNAQELBQAD
ggEBAD/GKSPNyQuAOw/jsYZesb+RMedbkzs18sSwlxAJQMUrrXwlVdHrA8q5WhE6
ABLqU1b8lQ8AWun07R8k5tqTmNvCARrAPRUqls/ryER+3Y9YEcxEaTc3jKNZFLbc
T6YtcnkdhxsiO136wtiuatpYL91RgCmuSpR8+7jEHhuFU01iaASu7ypFrUzrKHTF
bKwiLRQi1cMzVcLErq5CDEKiKhUkoDucyARFszrGt9vNIl/YCcBOkcNvM3c05Hn3
M++C29JwS3Hwbubg6WO3wjFjoEhpCwU6qRYUz3MRp4tHO4kxKXx+oQnUiFnR7vW0
YkNtGc1RUDHwecCTFpJtPb7Yu/E=
-----END CERTIFICATE-----`)

func TestNewStaticCertKeyContent(t *testing.T) {
	testCertProvider, err := NewStaticSNICertKeyContent("test-cert", serverCert, serverKey, "foo")
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name        string
		clientCA    CAContentProvider
		servingCert CertKeyContentProvider
		sniCerts    []SNICertKeyContentProvider

		expected    *dynamicCertificateContent
		expectedErr string
	}{
		{
			name:        "filled",
			clientCA:    &staticCAContent{name: "test-ca", caBundle: &caBundleAndVerifier{caBundle: []byte("content-1")}},
			servingCert: testCertProvider,
			sniCerts:    []SNICertKeyContentProvider{testCertProvider},
			expected: &dynamicCertificateContent{
				clientCA: caBundleContent{caBundle: []byte("content-1")},
				// ignore sni names for serving cert
				servingCert: certKeyContent{cert: serverCert, key: serverKey},
				sniCerts:    []sniCertKeyContent{{certKeyContent: certKeyContent{cert: serverCert, key: serverKey}, sniNames: []string{"foo"}}},
			},
		},
		{
			name:     "nil",
			expected: &dynamicCertificateContent{clientCA: caBundleContent{}, servingCert: certKeyContent{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &DynamicServingCertificateController{
				clientCA:    test.clientCA,
				servingCert: test.servingCert,
				sniCerts:    test.sniCerts,
			}
			actual, err := c.newTLSContent()
			if !reflect.DeepEqual(actual, test.expected) {
				t.Error(dump.Pretty(actual))
			}
			switch {
			case err == nil && len(test.expectedErr) == 0:
			case err == nil && len(test.expectedErr) != 0:
				t.Errorf("missing %q", test.expectedErr)
			case err != nil && len(test.expectedErr) == 0:
				t.Error(err)
			case err != nil && err.Error() != test.expectedErr:
				t.Error(err)
			}
		})
	}
}
