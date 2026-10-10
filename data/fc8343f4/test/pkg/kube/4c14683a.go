/*
Copyright 2017 The Kubernetes Authors.

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

package certificate

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/wait"
	certificatesclient "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	"k8s.io/client-go/rest"
)

var (
	client1CertData = newCertificateData(`-----BEGIN CERTIFICATE-----
MIICBDCCAW2gAwIBAgIJAPgVBh+4xbGoMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNV
BAMMEHdlYmhvb2tfdGVzdHNfY2EwIBcNMTcwNzI4MjMxNTI4WhgPMjI5MTA1MTMy
MzE1MjhaMB8xHTAbBgNVBAMMFHdlYmhvb2tfdGVzdHNfY2xpZW50MIGfMA0GCSqG
SIb3DQEBAQUAA4GNADCBiQKBgQDkGXXSm6Yun5o3Jlmx45rItcQ2pmnoDk4eZfl0
rmPa674s2pfYo3KywkXQ1Fp3BC8GUgzPLSfJ8xXya9Lg1Wo8sHrDln0iRg5HXxGu
uFNhRBvj2S0sIff0ZG/IatB9I6WXVOUYuQj6+A0CdULNj1vBqH9+7uWbLZ6lrD4b
a44x/wIDAQABo0owSDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DAdBgNVHSUEFjAU
BggrBgEFBQcDAgYIKwYBBQUHAwEwDwYDVR0RBAgwBocEfwAAATANBgkqhkiG9w0B
AQsFAAOBgQCpN27uh/LjUVCaBK7Noko25iih/JSSoWzlvc8CaipvSPofNWyGx3Vu
OdcSwNGYX/pp4ZoAzFij/Y5u0vKTVLkWXATeTMVmlPvhmpYjj9gPkCSY6j/SiKlY
kGy0xr+0M5UQkMBcfIh9oAp9um1fZHVWAJAGP/ikZgkcUey0LmBn8w==
-----END CERTIFICATE-----`, `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDkGXXSm6Yun5o3Jlmx45rItcQ2pmnoDk4eZfl0rmPa674s2pfY
a1RCvoYh8ZP1lf6xRYLDAVOo1LLWn9fk8mM9WtvHhi3rEF6uIXrnZhEkoUsO8a2i
lqW7Zf/JkUe5Y9GaMhQTpYJ8+w5BLlUWC1pyks2+7FMZBK4cWu3sX80Q/GGelUpV
eobxsxcxYdsIw5BQj7vWuRg2l072wn8VRs4jaQ4OLrOBfdjbJwz8EZW+1mdufaAl
6pqjmTmRKjFBcUQFf8XNQ3JeOo+/c5LP9JObBCJliSQh2sYn2QpSpoMHZZ9QwI2G
zw/lUFk6Ed9dxRhh91F6SxhbYKI4I/ayDDvUv7RSqjdrtgMJSaK8nLf+Gm08s1rY
MC+nQnyJRetEIUZh/uu2ZpKVt62NvzlOxXZJ984BujPY86oQoOXfNb0ClFYf2t0Q
k+cMTx3PFOTc4oHuMsG1po4vw6lhK+mBZa/57aBKsSsjOQkFA4u9oYDMLtsdHE7s
Kfp0VIPqA38C0Mxj6PDUqnVYuWzwQ0ismqehx7O18ihPdAdKtZvCwYdzc7D8hd+v
I+ouPX0/z5VicoGAPxwYgL7Z4XrX9sgsDNmKlWgEVn9YMI7GrVH3/TK0pKRDWLRc
Of/JgB0dt+knchSsEPI952TUwyV96Zf/FywLxMWOdLBDAaI0ihvASHDxKpOWZrTx
cC7WiMw5bLkgfdYiYBWvlNzC+qHD8cx7S++vstu45R5nF77aip3dmOzWsj5PmrQs
+VkSfzShKhNvwRTVfNjIeFfCvBFCjvxmJ7r+NTSnjb==
-----END RSA PRIVATE KEY-----`)
	client2CertData = newCertificateData(`-----BEGIN CERTIFICATE-----
MIICBDCCAW2gAwIBAgIJAPgVBh+4xbGnMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNV
BAMMEHdlYmhvb2tfdGVzdHNfY2EwIBcNMTcwNzI4MjMxNTI4WhgPMjI5MTA1MTMy
MzE1MjhaMB8xHTAbBgNVBAMMFHdlYmhvb2tfdGVzdHNfY2xpZW50MIGfMA0GCSqG
SIb3DQEBAQUAA4GNADCBiQKBgQDQQLzbrmHbtlxE7wViaoXFp5tQx7zzM2Ed7O1E
gs3JUws5KkPbNrejLwixvLkzzU152M43UGsyKDn7HPyjXDogTZSW6C257XpYodk3
S/gZS9oZtPss4UJuJioQk/M8X1ZjYP8kCTArOvVRJeNQL8GM7h5QQ6J5LUq+IdZb
T0retQIDAQABo0owSDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DAdBgNVHSUEFjAU
BggrBgEFBQcDAgYIKwYBBQUHAwEwDwYDVR0RBAgwBocEfwAAATANBgkqhkiG9w0B
AQsFAAOBgQBdAxoU5YAmp0d+5b4qg/xOGC5rKcnksQEXYoGwFBWwaKvh9oUlGGxI
A5Ykf2TEl24br4tLmicpdxUX4H4PbkdPxOjM9ghIKlmgHo8vBRC0iVIwYgQsw1W8
ETY34Or+PJqaeslqx/t7kUKY5UIF9DLVolsIiAHveJNR2uBWiP0KiQ==
-----END CERTIFICATE-----`, `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDQQLzbrmHbtlxE7wViaoXFp5tQx7zzM2Ed7O1Egs3JUws5KkPb
UHkIxAZpaozwlj632C92IFMsUrr4KNbQJhrrPafv1i062hNzSPJ7a/BAr3ksYkfS
1uSfuXFQx/m8G8mOmr1smUgzQfsDRuPop9Dd4N4wP0O6dtX+hGvUZ2GZYRJVageN
TQYaEbXDKe1AFNx4d+noRG3KECiUdnklpB7DXfa8pu1BZN9Atj8K3+HW18ogCMEo
U42lNCpdLmMa9VSS3UXZzpG7wGPNPh67YbGo31gQvWUq1f5QMbpdF6qCIKBQ1VOV
GD50Raurl7rx3YJmZcz53N9BJGq0COaD2044eIFER+0Dz6xUEvIv+/RdGBLAPSFH
5Fu+G68K0cCGxRLnBLUp0JPUgy8kC2RmgZUDVBh9jhtrMqDs8ztseez6GAQ5jh6B
ojJrgbj2UyOTUnHtlUXVaI/9vLSOMDFiXdkKGEx+0Cbch01C8aSzhA5wuWDVHHhO
PanZ7QA+hOyFp2rpiqLSj5GFRTA+RDLVuH3VqPCJwneT9F2GUqO/Tv61zKd0IzP9
uojvLR+ujTA/Xu5phbTQUTKyAmTguos87jHHX7mLNtINfAJ1cKtbFWUH11xFmX2f
5KSRKlokp2amEMAkc2GKD0WRxL/UZTAwyueuB6K8X4Pfh7WaCD1IIBGG1oCOYw5F
SA91dRJoP6UCytyZVAcPevaRnHjjuvNVsyxh1xShEK45GISoQ33ypRaQfiBWV8tH
hb8Bn8oNMzw4AbbQgHPures0zAsu6LjspK8MbRZD6L7K
-----END RSA PRIVATE KEY-----`)
)

type certificateData struct {
	keyPEM         []byte
	certificatePEM []byte
	certificate    *tls.Certificate
}

func newCertificateData(certificatePEM string, keyPEM string) *certificateData {
	certificate, err := tls.X509KeyPair([]byte(certificatePEM), []byte(keyPEM))
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize certificate: %v", err))
	}
	certs, err := x509.ParseCertificates(certificate.Certificate[0])
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize certificate leaf: %v", err))
	}
	certificate.Leaf = certs[0]
	return &certificateData{
		keyPEM:         []byte(keyPEM),
		certificatePEM: []byte(certificatePEM),
		certificate:    &certificate,
	}
}

type fakeManager struct {
	cert    atomic.Value // Always a *tls.Certificate
	healthy bool
}

func (f *fakeManager) SetCertificateSigningRequestClient(certificatesclient.CertificateSigningRequestInterface) error {
	return nil
}

func (f *fakeManager) ServerHealthy() bool { return f.healthy }

func (f *fakeManager) Start()                     {}
func (f *fakeManager) Stop()                      {}
func (f *fakeManager) RotateCerts() (bool, error) { return false, nil }

func (f *fakeManager) Current() *tls.Certificate {
	if val := f.cert.Load(); val != nil {
		return val.(*tls.Certificate)
	}
	return nil
}

func (f *fakeManager) setCurrent(cert *tls.Certificate) {
	f.cert.Store(cert)
}

func TestRotateShutsDownConnections(t *testing.T) {

	// This test fails if you comment out the t.closeAllConns() call in
	// transport.go and don't close connections on a rotate.

	stop := make(chan struct{})
	defer close(stop)

	m := new(fakeManager)
	m.setCurrent(client1CertData.certificate)

	// The last certificate we've seen.
	lastSeenLeafCert := new(atomic.Value) // Always *x509.Certificate

	lastSerialNumber := func() *big.Int {
		if cert := lastSeenLeafCert.Load(); cert != nil {
			return cert.(*x509.Certificate).SerialNumber
		}
		return big.NewInt(0)
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		if r.TLS != nil && len(r.TLS.PeerCertificates) != 0 {
			// Record the last TLS certificate the client sent.
			lastSeenLeafCert.Store(r.TLS.PeerCertificates[0])
		}
		w.Write([]byte(`{}`))
	}

	s := httptest.NewUnstartedServer(http.HandlerFunc(h))
	s.TLS = &tls.Config{
		// Just request a cert, we don't need to verify it.
		ClientAuth: tls.RequestClientCert,
	}
	s.StartTLS()
	defer s.Close()

	c := &rest.Config{
		Host: s.URL,
		TLSClientConfig: rest.TLSClientConfig{
			// We don't care about the server's cert.
			Insecure: true,
		},
		ContentConfig: rest.ContentConfig{
			// This is a hack. We don't actually care about the serializer.
			NegotiatedSerializer: serializer.NegotiatedSerializerWrapper(runtime.SerializerInfo{}),
		},
	}

	// Check for a new cert every 10 milliseconds
	if _, err := updateTransport(stop, 10*time.Millisecond, c, m, 0); err != nil {
		t.Fatal(err)
	}

	client, err := rest.UnversionedRESTClientFor(c)
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Get().Do(context.TODO()).Error(); err != nil {
		t.Fatal(err)
	}
	firstCertSerial := lastSerialNumber()

	// Change the manager's certificate. This should cause the client to shut down
	// its connections to the server.
	m.setCurrent(client2CertData.certificate)

	err = wait.PollImmediate(time.Millisecond*50, wait.ForeverTestTimeout, func() (done bool, err error) {
		client.Get().Do(context.TODO())
		if firstCertSerial.Cmp(lastSerialNumber()) != 0 {
			// The certificate changed!
			return true, nil
		}
		t.Logf("Certificate not changed, will retry.")
		return false, nil
	})
	if err != nil {
		t.Fatal("certificate rotated but client never reconnected with new cert")
	}
}
