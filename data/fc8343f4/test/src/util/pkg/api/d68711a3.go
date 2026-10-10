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

package proxy

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	utilnet "k8s.io/apimachinery/pkg/util/net"
)

func TestDialURL(t *testing.T) {
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(localhostCert) {
		t.Fatal("error setting up localhostCert pool")
	}

	cert, err := tls.X509KeyPair(localhostCert, localhostKey)
	if err != nil {
		t.Fatal(err)
	}
	var d net.Dialer

	testcases := map[string]struct {
		TLSConfig   *tls.Config
		Dial        utilnet.DialFunc
		ExpectError string
		ExpectProto string
	}{
		"insecure": {
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		},
		"secure, no roots": {
			TLSConfig:   &tls.Config{InsecureSkipVerify: false},
			ExpectError: "unknown authority|not trusted",
		},
		"secure with roots": {
			TLSConfig: &tls.Config{InsecureSkipVerify: false, RootCAs: roots},
		},
		"secure with mismatched server": {
			TLSConfig:   &tls.Config{InsecureSkipVerify: false, RootCAs: roots, ServerName: "bogus.com"},
			ExpectError: "not bogus.com",
		},
		"secure with matched server": {
			TLSConfig: &tls.Config{InsecureSkipVerify: false, RootCAs: roots, ServerName: "example.com"},
		},

		"insecure, custom dial": {
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
			Dial:      d.DialContext,
		},
		"secure, no roots, custom dial": {
			TLSConfig:   &tls.Config{InsecureSkipVerify: false},
			Dial:        d.DialContext,
			ExpectError: "unknown authority|not trusted",
		},
		"secure with roots, custom dial": {
			TLSConfig: &tls.Config{InsecureSkipVerify: false, RootCAs: roots},
			Dial:      d.DialContext,
		},
		"secure with mismatched server, custom dial": {
			TLSConfig:   &tls.Config{InsecureSkipVerify: false, RootCAs: roots, ServerName: "bogus.com"},
			Dial:        d.DialContext,
			ExpectError: "not bogus.com",
		},
		"secure with matched server, custom dial": {
			TLSConfig: &tls.Config{InsecureSkipVerify: false, RootCAs: roots, ServerName: "example.com"},
			Dial:      d.DialContext,
		},
		"ensure we use http2 if specified": {
			TLSConfig:   &tls.Config{InsecureSkipVerify: false, RootCAs: roots, ServerName: "example.com", NextProtos: []string{"http2"}},
			Dial:        d.DialContext,
			ExpectProto: "http2",
		},
		"ensure we use http/1.1 if unspecified": {
			TLSConfig:   &tls.Config{InsecureSkipVerify: false, RootCAs: roots, ServerName: "example.com"},
			Dial:        d.DialContext,
			ExpectProto: "http/1.1",
		},
		"ensure we use http/1.1 if available": {
			TLSConfig:   &tls.Config{InsecureSkipVerify: false, RootCAs: roots, ServerName: "example.com", NextProtos: []string{"http2", "http/1.1"}},
			Dial:        d.DialContext,
			ExpectProto: "http/1.1",
		},
	}

	for k, tc := range testcases {
		func() {
			ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {}))
			defer ts.Close()
			ts.TLS = &tls.Config{Certificates: []tls.Certificate{cert}, NextProtos: []string{"http2", "http/1.1"}}
			ts.StartTLS()

			// Make a copy of the config
			tlsConfigCopy := tc.TLSConfig.Clone()
			// Clone() mutates the receiver (!), so also call it on the copy
			tlsConfigCopy.Clone()
			transport := &http.Transport{
				DialContext:     tc.Dial,
				TLSClientConfig: tlsConfigCopy,
			}

			extractedDial, err := utilnet.DialerFor(transport)
			if err != nil {
				t.Fatal(err)
			}
			if fmt.Sprintf("%p", extractedDial) != fmt.Sprintf("%p", tc.Dial) {
				t.Fatalf("%s: Unexpected dial", k)
			}

			extractedTLSConfig, err := utilnet.TLSClientConfig(transport)
			if err != nil {
				t.Fatal(err)
			}
			if extractedTLSConfig == nil {
				t.Fatalf("%s: Expected tlsConfig", k)
			}

			u, _ := url.Parse(ts.URL)
			_, p, _ := net.SplitHostPort(u.Host)
			u.Host = net.JoinHostPort("127.0.0.1", p)
			conn, err := DialURL(context.Background(), u, transport)

			// Make sure dialing doesn't mutate the transport's TLSConfig
			if !reflect.DeepEqual(tc.TLSConfig, tlsConfigCopy) {
				t.Errorf("%s: transport's copy of TLSConfig was mutated\n%s", k, cmp.Diff(tc.TLSConfig, tlsConfigCopy))
			}

			if err != nil {
				if tc.ExpectError == "" {
					t.Errorf("%s: expected no error, got %q", k, err.Error())
				}
				if tc.ExpectError != "" && !regexp.MustCompile(tc.ExpectError).MatchString(err.Error()) {
					t.Errorf("%s: expected error containing %q, got %q", k, tc.ExpectError, err.Error())
				}
				return
			}

			tlsConn := conn.(*tls.Conn)
			if tc.ExpectProto != "" {
				if tlsConn.ConnectionState().NegotiatedProtocol != tc.ExpectProto {
					t.Errorf("%s: expected proto %s, got %s", k, tc.ExpectProto, tlsConn.ConnectionState().NegotiatedProtocol)
				}
			}

			conn.Close()
			if tc.ExpectError != "" {
				t.Errorf("%s: expected error %q, got none", k, tc.ExpectError)
			}
		}()
	}

}

// localhostCert was generated from crypto/tls/generate_cert.go with the following command:
//
//	go run generate_cert.go  --rsa-bits 2048 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var localhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDGTCCAgGgAwIBAgIRAKfNl1LEAt7nFPYvHBnpv2swDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBAKww39FwmV5lDIbAUIAuSYYVtZke6bca1oyq19ZrRL0uavwPXSJm
+Qxt4RKUQhzYhZ/alJp8iRfu/Z+Yv9Beez89dQB9V8YnHj/AX4Jph9lJ2aawWMI6
AqPLdIzKLQVVvPw+UVKH9x8yy08H/23AIFGyK4Dbht+KZJeUbJQFiGlRFJim8atx
KA3C9NzCHw6hyhP46jguLl65rcxLMSzcTz97ToG0MP66YEUbsA/YzFTKDwht7ESH
nRMBnQ4wZfWpvAiXMr3XJGOa3NYJy1A+WkWyrfZO7guwsZ4L6dGqnlPpzA5QkKYx
H9Z5K1bUaYEi0Yi2ug7Jkvd1HE179nkF7t0CAwEAAaNoMGYwDgYDVR0PAQH/BAQD
AgKkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA8GA1UdEwEB/wQFMAMBAf8wLgYDVR0R
BCcwJYILZXhhbXBsZS5jb22HBH8AAAGHEAAAAAAAAAAAAAAAAAAAAAEwDQYJKoZI
hvcNAQELBQADggEBAAKSQToD1iLujFhQwaLnPVRV6r4nEFVXCxXYtQNEX1DVSKSj
JYbBGJnL50oc0N4Ar+Spqofm+THkiTQJUzptPtnYIzNpKYdE6+bPwqURWzFEI2OF
ks3fYZ4ZdbMbmJRo1qPJO34emm4KrOl9aoV0qwp2QyTvHgLroU3icKoe4e7+p4KK
02Rt3qczHvCKoUnw6m07Ql0n9e7Ncpujcs2A8PaQ1iPX+BVOmvjTVT8y5NSRDzwL
a2wur8BSZ5E8SVzzvNZJlLSi6BbObQUjALHkjVYm11dWv/BY8jHdt+iFhbNBRASx
ENuih3pX1Poki1qRYOtB/vAS99E1ORj9zJlUlzo=
-----END CERTIFICATE-----`)

// localhostKey is the private key for localhostCert.
var localhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEArDDf0XCZXmUMhsBQgC5JhhW1mR7ptxrWjKrX1mtEvS5q/A9d
Weq1pt5KZarMvWDXr0OVhEipL+48v8C/3K57XT40ZM5dSJyrq6RVPuAt4nidzAar
sKpIA8C9chOExBG6/M9dtep4DnWKZDZ/ViypvHjdGKtI96WAN9chjkMYLMaSIewY
q8Lronj17WoOgdIIl/gBDr2FnIRJvyuWMBHQi4gEkLDF/jUKdOvtr2QapghrjS1F
FwBHQAZVuQOX3uN3MhnkZkzIM1MH5IBRfW9aEzKm0m0ht9OCDRyF4vVEF+SfPQiQ
KnJf6DIwBEVDGybstHN6BwwH16DXCvU8nppO3mQgcNporxBgyMoYve+0LQO6rwR7
ruEefaZNWBhy56gSnDeRZGtooSZ0RlPYaEgt1qsZB+YTqfSa3ogI/k5ng5/qzowc
Fn1PxNZBbpllmkSXm/VCSI5NVJdMWq3LWPIu5zRz4aL2bf+SK3gQ+lpNGOsmgf4O
ZMo08DGaPH+QJ/v/DDiNTnC044+FQOXf6oo87eu395W20LorIXMDYWyGa/PlqPSS
fAE5XQWC7+1lEOsv4idBLq9VoiP133ntyA32miN+2vTH8aTwiFJtM/LAg2zYzpKz
2JOD9bno+Pto5mzIPPIRYD+jsKMnzZEQgDESfnpGo4iW8RvyRkJO5Uh3V3PWs/qf
P6p+MKodFNvVFSoWQYXRfi3NiFF+RRk+hh+8FCoE/ei2rA0sxiqt1g9TJkMCLHmj
DGXn9YeSuSNNbCgN5Q7UOvgsTkVMdMSyNXfTvLzwZ7Uf/v9QqOhXFGOyPTjA5gD7
uni5XSXahGXYRGzABZt5viq8p5ZCP4L7AvmnKpRvzSj+5gQqDiR+lg6XNZRE6Vin
vtNZlCvrpFIqm52niwbNfYaLtDHJyOCuUgx9Hb4fuLJ7JmqrBOAJgiO/phCdJ3Pb
XIFjkcYI9hlkPy+dDf402dCIxjfi5l5w6cp2kXmCmd5gR9bhprmof7ea6W2RMYv4
3bemHKN3ll8dDyeVQabzdztiRvRTCm6N5LdF0DkADcJq1WAiSQhYFHVCoC+nnW5j
ZSzdSxFGkGTllb1L+r1MELDtvllHR5aq+Rq8pO7CIF+EucJcdBnZ+mk3bTMbfg61
cRqM1qAoBjEoraARAQQ3UWxxe5u6WiS1g8igZlAyGQWOsTT7dyu31Ioctm4yjzJV
b+VIfg/TeedzHHKj3TbkjtNKzU7R85wfAlBjny026BkvbjP1w9Et03DjQUIt3P2x
IdTfaFKy2XiBWSmeQkEsAtaj95d2MzDxSKklcTg+Rjmtw9BGqk0KqJTk0fKm5M+g
ex9UNhu/70e5IqcVGcK0nrYxe8/EUae21o1UfYMbLJJPspNdO5a1G34LJQfGR7Q4
cs0NXSLLdk2LAcKc9mX74IiVQGQ6mUi/1oHQADyrBkmDz5gjsd1IEOYXkYyUkLFE
yl9XN1bqtsZXt3JThYgEbqN6gwEKEvU9gErcH05R3bYMAm5QmDW32fCekzMpLcs9
KsBrDx/1CzA9BjWd7jZh3v1iN6IXYsc6cg7wUXg4h5RmOG7HOOn=
-----END RSA PRIVATE KEY-----
`)
