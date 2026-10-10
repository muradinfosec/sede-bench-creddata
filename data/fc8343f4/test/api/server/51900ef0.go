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

package podlogs

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	authorizationv1 "k8s.io/api/authorization/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/cert"
	"k8s.io/component-base/cli/flag"
	"k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	"k8s.io/kubernetes/test/integration/framework"
	"k8s.io/kubernetes/test/utils/ktesting"
)

type caWithClient struct {
	CACert     []byte
	ClientCert []byte
	ClientKey  []byte
}

func newTestCAWithClient(caSubject pkix.Name, caSerial *big.Int, clientSubject pkix.Name, subjectSerial *big.Int) (*caWithClient, error) {
	ca := &x509.Certificate{
		SerialNumber:          caSerial,
		Subject:               caSubject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return nil, err
	}

	caPEM := new(bytes.Buffer)
	err = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	if err != nil {
		return nil, err
	}

	clientCert := &x509.Certificate{
		SerialNumber: subjectSerial,
		Subject:      clientSubject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(24 * time.Hour),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	clientCertPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	clientCertPrivateKeyPEM := new(bytes.Buffer)
	err = pem.Encode(clientCertPrivateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(clientCertPrivateKey),
	})
	if err != nil {
		return nil, err
	}

	clientCertBytes, err := x509.CreateCertificate(rand.Reader, clientCert, ca, &clientCertPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return nil, err
	}

	clientCertPEM := new(bytes.Buffer)
	err = pem.Encode(clientCertPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCertBytes,
	})
	if err != nil {
		return nil, err
	}

	return &caWithClient{
		CACert:     caPEM.Bytes(),
		ClientCert: clientCertPEM.Bytes(),
		ClientKey:  clientCertPrivateKeyPEM.Bytes(),
	}, nil
}

func TestClientCAUpdate(t *testing.T) {
	testClientCA(t, false)
}

func TestClientCARecreate(t *testing.T) {
	testClientCA(t, true)
}

func testClientCA(t *testing.T, recreate bool) {
	tCtx := ktesting.Init(t)

	frontProxyCA, err := newTestCAWithClient(
		pkix.Name{
			CommonName: "test-front-proxy-ca",
		},
		big.NewInt(43),
		pkix.Name{
			CommonName:   "test-aggregated-apiserver",
			Organization: []string{"system:masters"},
		},
		big.NewInt(86),
	)
	if err != nil {
		t.Error(err)
		return
	}

	clientCA, err := newTestCAWithClient(
		pkix.Name{
			CommonName: "test-client-ca",
		},
		big.NewInt(42),
		pkix.Name{
			CommonName:   "system:admin",
			Organization: []string{"system:masters"},
		},
		big.NewInt(84),
	)
	if err != nil {
		t.Error(err)
		return
	}

	clientCAFilename := ""
	frontProxyCAFilename := ""

	kubeClient, kubeconfig, tearDownFn := framework.StartTestServer(tCtx, t, framework.TestServerSetup{
		ModifyServerRunOptions: func(opts *options.ServerRunOptions) {
			opts.GenericServerRunOptions.MaxRequestBodyBytes = 1024 * 1024
			clientCAFilename = opts.Authentication.ClientCert.ClientCA
			frontProxyCAFilename = opts.Authentication.RequestHeader.ClientCAFile
			opts.Authentication.RequestHeader.AllowedNames = append(opts.Authentication.RequestHeader.AllowedNames, "test-aggregated-apiserver")
		},
	})
	defer tearDownFn()

	// wait for request header info
	err = wait.PollImmediate(100*time.Millisecond, 30*time.Second, waitForConfigMapCAContent(t, kubeClient, "requestheader-client-ca-file", "-----BEGIN CERTIFICATE-----", 1))
	if err != nil {
		t.Fatal(err)
	}
	// wait for client cert info
	err = wait.PollImmediate(100*time.Millisecond, 30*time.Second, waitForConfigMapCAContent(t, kubeClient, "client-ca-file", "-----BEGIN CERTIFICATE-----", 1))
	if err != nil {
		t.Fatal(err)
	}

	if recreate {
		if err := os.Remove(path.Join(clientCAFilename)); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(path.Join(frontProxyCAFilename)); err != nil {
			t.Fatal(err)
		}
	}

	// when we run this the second time, we know which one we are expecting
	if err := os.WriteFile(clientCAFilename, clientCA.CACert, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(frontProxyCAFilename, frontProxyCA.CACert, 0644); err != nil {
		t.Fatal(err)
	}

	time.Sleep(4 * time.Second)

	acceptableCAs, err := cert.GetClientCANamesForURL(kubeconfig.Host)
	if err != nil {
		t.Fatal(err)
	}

	expectedCAs := []string{"test-client-ca", "test-front-proxy-ca"}
	if len(expectedCAs) != len(acceptableCAs) {
		t.Fatal(strings.Join(acceptableCAs, ":"))
	}
	for i := range expectedCAs {
		if !strings.Contains(acceptableCAs[i], expectedCAs[i]) {
			t.Errorf("expected %q, got %q", expectedCAs[i], acceptableCAs[i])
		}
	}

	// wait for updated request header info that contains both
	err = wait.PollImmediate(100*time.Millisecond, 30*time.Second, waitForConfigMapCAContent(t, kubeClient, "requestheader-client-ca-file", "-----BEGIN CERTIFICATE-----", 2))
	if err != nil {
		t.Error(err)
	}
	err = wait.PollImmediate(100*time.Millisecond, 30*time.Second, waitForConfigMapCAContent(t, kubeClient, "requestheader-client-ca-file", string(frontProxyCA.CACert), 1))
	if err != nil {
		t.Error(err)
	}
	// wait for updated client cert info that contains both
	err = wait.PollImmediate(100*time.Millisecond, 30*time.Second, waitForConfigMapCAContent(t, kubeClient, "client-ca-file", "-----BEGIN CERTIFICATE-----", 2))
	if err != nil {
		t.Error(err)
	}
	err = wait.PollImmediate(100*time.Millisecond, 30*time.Second, waitForConfigMapCAContent(t, kubeClient, "client-ca-file", string(clientCA.CACert), 1))
	if err != nil {
		t.Error(err)
	}

	// Test an aggregated apiserver client (signed by the new front proxy CA) is authorized
	extensionApiserverClient, err := kubernetes.NewForConfig(&rest.Config{
		Host: kubeconfig.Host,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:     kubeconfig.TLSClientConfig.CAData,
			CAFile:     kubeconfig.TLSClientConfig.CAFile,
			ServerName: kubeconfig.TLSClientConfig.ServerName,
			KeyData:    frontProxyCA.ClientKey,
			CertData:   frontProxyCA.ClientCert,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	// Call an endpoint to make sure we are authenticated
	err = extensionApiserverClient.AuthorizationV1().RESTClient().
		Post().
		Resource("subjectaccessreviews").
		VersionedParams(&metav1.CreateOptions{}, scheme.ParameterCodec).
		Body(&authorizationv1.SubjectAccessReview{
			Spec: authorizationv1.SubjectAccessReviewSpec{
				ResourceAttributes: &authorizationv1.ResourceAttributes{
					Verb:      "create",
					Resource:  "pods",
					Namespace: "default",
				},
				User: "deads2k",
			},
		}).
		SetHeader("X-Remote-User", "test-aggregated-apiserver").
		SetHeader("X-Remote-Group", "system:masters").
		Do(context.Background()).
		Into(&authorizationv1.SubjectAccessReview{})
	if err != nil {
		t.Error(err)
	}

	// Test a client signed by the new ClientCA is authorized
	testClient, err := kubernetes.NewForConfig(&rest.Config{
		Host: kubeconfig.Host,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:     kubeconfig.TLSClientConfig.CAData,
			CAFile:     kubeconfig.TLSClientConfig.CAFile,
			ServerName: kubeconfig.TLSClientConfig.ServerName,
			KeyData:    clientCA.ClientKey,
			CertData:   clientCA.ClientCert,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	// Call an endpoint to make sure we are authenticated
	_, err = testClient.CoreV1().Nodes().List(tCtx, metav1.ListOptions{})
	if err != nil {
		t.Error(err)
	}
}

func waitForConfigMapCAContent(t *testing.T, kubeClient kubernetes.Interface, key, content string, count int) func() (bool, error) {
	return func() (bool, error) {
		clusterAuthInfo, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.TODO(), "extension-apiserver-authentication", metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return false, nil
		}
		if err != nil {
			return false, err
		}

		ca := clusterAuthInfo.Data[key]
		if strings.Count(ca, content) == count {
			return true, nil
		}
		t.Log(ca)
		return false, nil
	}
}

var serverKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA13f50PPWuR/InxLIoJjHdNSG+jVUd25CY7ZL2J023X2BAY+1
X5LoRT7a6KqZOdV18dUVM47/X9c/iK6OMB496LL522b380a+QVJXXq2KkHNy7c8a
i80IHf0EoX/+9Cr3voQHUep1I6rv5tgsESfz95ga7XhWql9u6Z10FQnRWKrhFBml
NTMAQYLkC6V0ZccDVScZXEjrrbvYxqxe6BUZGDDqOOi3Ld26NfYQ07w53zN5d3Y7
WAty/JJATA0BzLbeJoIoV070CbHEdcTGP1NSiyf1h42y1hgY6dwVVhJUGlVUn1hU
84Y9D2YpFn/iA7P7ywJsbAJ1nhBcvCrJXeDWxmUHYZXLGQlYRVTJCybeyl4g8H+U
BrVXg1wo52skHoprri4d++YIPZNzMsHTGNLPOkSst7IbsldNuI3Lga3OF/HjHWw8
y3Z+9vCbxbU4z/mi7aQLf1PhwYxdMiHPGgS15AbJHji8kL8TY8HT+0cGXxT4qfla
XPK29dAy6+xzjWavWmgNBP53E8GcZ3gTfQ+ZphQWpNQE9oBaqVky7HhjADZT6ddK
pNym0QnBI5Im/tYY4tdteeqeReZmnPe+bkHwokOogvy4vUFIiS+eeVbQ67o/ImGG
YWdHntv2PDFO9+3NGiDVZ+g6dqTz8tvA080DQcSTa15BcESwVHGzeJMvQZe12V3y
vA9h38AZsOda3wdw0UF6Ji1H1CZX1Yr5eJ06ANQ3j4CDNj1JEZaGAbSiI5DpOUAg
sN/3vCYSWMuylHiAw/XbdHqyKND1N+jp/tGVxkx85yW1b3P08GyoMMM2PPHDVnny
Zntqf2/zvWaztTEuEklpl/KqrJf2dtwh9jo+eqB7MnjQ8TaXT6P7oz3SGBJC9H0z
wAbCQG4AS7q3W6B5VQq1aYCrnXQoMLNPL6Fm7TojFhnpH7MOSk8AtCCqqLg/7PJl
CUmmE3HZBwxXK30eYM0AnDh1Lq/kRdE7gsqTepclw2U6lbnX50c14vHjFa2mp9J3
JRUkiZjboTke0b+0hCZIsgQJDLeEcqwsV+nbiMJKrzZMsfvonnAdvHg6APxx+del
2ZpbNXdbAribIK51vHiTPVm90a9YWifVtXhrUyIvlMoJCtRsczvUYdEs5+2u5FXj
rQzXJ5mv3abFI4idPGUYUcsQawzYTWJLKb0zQuWLD5n9tDDXQpwP+et5fGmJfNLs
mBGH1RGCFY3e0LCZF6MbolkqtZjrp5w1loNQm+8orlahvA+zToQipD4jhVTX4ZB6
bbCF9AT6OGlQlcBgbhfey4TE/VgtP5DVhYTlJNW11okFfiU2lHOFfc43Hp5PmLUp
aS2ebgbGLwUgrOiASqmmSzC0ykP5n2qSEGTgY0xCDBj8OxCN0ofPRShuL+TvnIms
qmrNIEjrNxMdMghHzH62UV/MnYOD6IuUOJUOEog0ITKK6mL52HuizYwXaGh0ci7p
fGLGImQMoV/ERBpUjPww+XIizXZ/ZDFaRaTnCsmyBbHbYyRavcGDPad4YNj/Zk6X
Fl1rqorLYbSgmkZ4dFFUURUwubRfZXDO91dDoHiEvZ4DzGqJkQL5
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

var anotherServerKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

var anotherServerCert = []byte(`-----BEGIN CERTIFICATE-----
MIIFJjCCAw6gAwIBAgIJAOcEAbv8NslfMA0GCSqGSIb3DQEBCwUAMEAxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTETMBEGA1UECgwKQWNtZSwgSW5jLjEPMA0GA1UE
AwwGc29tZUNBMCAXDTE4MDYwODEzMzkyNFoYDzIyMTgwNDIxMTMzOTI0WjBDMQsw
CQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExEzARBgNVBAoMCkFjbWUsIEluYy4xEjAQ
BgNVBAMMCWxvY2FsaG9zdDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
AJWSTkcwo28xdUmglxR8orbnZRPqt2eWxoUtk6Vxunei2qKrcDaOFAt4vk5zJ3ba
z9qTZFf2WIr7KdXx8/dhahQE/l6WHBvKS5uIi9JNFLnzlnSz01/OxwiYymeK+35s
7Lco6iuLL/SN8ntYyotFRNQS6xdSlbxBP9T1wroOvBAT/2hdOvrr3K+qPIWzGM4H
Sks4UWJvQt5BnGsDaH6gA/ypP85/qTJOAZaMuFILVxV4I9mxJ5lRPrUVsi79vkzj
ZZi512gVt/zL2bSgt5iLNihvumgTbmPD3Vs1rzLad9K/4m4Jx87jCrb1yeV0YU2k
k+EGFzWxsG9WyFdcpZ1/M3VzxEJLmE0A3LWmEAntEV7dSPRad4FUEcckzPtCHrW7
FNJPVbwcENO5oFPEYpNw8VZ327iWP8ERJxEU5IL+qI7UcIrM6xiGOcd+qXkbN9Aq
kTL97ibagvxVOpwkGx8cLLEMgf7pjei+23KbEBQ8nXiveYb0UMurW9r/7zHMOzDK
OZEhVB46C+WXOAzm8IqrPqexe0tFiCTAcyTnI5oAjM4GeyZnO2zJq5uNxi8zmlSy
I7j1VaTqe2ICf5akX+yt4WV6zHWl5AptNWi/C+fXft1X3Y6dxx75mKQofG3Pio7P
50fUXPtiSivx9dyx91uQq/BZ3PxT+k4G77zcLEFdNYBvAgMBAAGjHjAcMBoGA1Ud
EQQTMBGCCWxvY2FsaG9zdIcEfwAAATANBgkqhkiG9w0BAQsFAAOCAgEABL8kffi7
48qSD+/l/UwCYdmqta1vAbOkvLnPtfXe1XlDpJipNuPxUBc8nNTemtrbg0erNJnC
jQHodqmdKBJJOdaEKTwAGp5pYvvjlU3WasmhfJy+QwOWgeqjJcTUo3+DEaHRls16
AZXlsp3hB6z0gzR/qzUuZwpMbL477JpuZtAcwLYeVvLG8bQRyWyEy8JgGDoYSn8s
Z16s+r6AX+cnL/2GHkZ+oc3iuXJbnac4xfWTKDiYnyzK6RWRnoyro7X0jiPz6XX3
wyoWzB1uMSCXscrW6ZcKyKqz75lySLuwGxOMhX4nGOoYHY0ZtrYn5WK2ZAJxsQnn
8QcjPB0nq37U7ifk1uebmuXe99iqyKnWaLvlcpe+HnO5pVxFkSQEf7Zh+hEnRDkN
IBzLFnqwDS1ug/oQ1aSvc8oBh2ylKDJuGtPNqGKibNJyb2diXO/aEUOKRUKPAxKa
dbKsc4Y1bhZNN3/MICMoyghwAOiuwUQMR5uhxTkQmZUwNrPFa+eW6GvyoYLFUsZs
hZfWLNGD5mLADElxs0HF7F9Zk6pSocTDXba4d4lfxsq88SyZZ7PbjJYFRfLQPzd1
CfvpRPqolEmZo1Y5Q644PELYiJRKpBxmX5GtC5j5eaUD9XdGKvXsGhb0m0gW75rq
iUnnLkZt2ya1cDJDiCnJjo7r5KxMo0XXFDc=
-----END CERTIFICATE-----
`)

func TestServingCertUpdate(t *testing.T) {
	testServingCert(t, false)
}

func TestServingCertRecreate(t *testing.T) {
	testServingCert(t, true)
}

func testServingCert(t *testing.T, recreate bool) {
	tCtx := ktesting.Init(t)

	var servingCertPath string

	_, kubeconfig, tearDownFn := framework.StartTestServer(tCtx, t, framework.TestServerSetup{
		ModifyServerRunOptions: func(opts *options.ServerRunOptions) {
			opts.GenericServerRunOptions.MaxRequestBodyBytes = 1024 * 1024
			servingCertPath = opts.SecureServing.ServerCert.CertDirectory
		},
	})
	defer tearDownFn()

	if recreate {
		if err := os.Remove(path.Join(servingCertPath, "apiserver.key")); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(path.Join(servingCertPath, "apiserver.crt")); err != nil {
			t.Fatal(err)
		}
	}

	if err := os.WriteFile(path.Join(servingCertPath, "apiserver.key"), serverKey, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path.Join(servingCertPath, "apiserver.crt"), serverCert, 0644); err != nil {
		t.Fatal(err)
	}

	time.Sleep(4 * time.Second)

	// get the certs we're actually serving with
	_, actualCerts, err := cert.GetServingCertificatesForURL(kubeconfig.Host, "")
	if err != nil {
		t.Fatal(err)
	}
	if err := checkServingCerts(serverCert, actualCerts); err != nil {
		t.Fatal(err)
	}
}

func TestSNICert(t *testing.T) {
	var servingCertPath string

	tCtx := ktesting.Init(t)

	_, kubeconfig, tearDownFn := framework.StartTestServer(tCtx, t, framework.TestServerSetup{
		ModifyServerRunOptions: func(opts *options.ServerRunOptions) {
			opts.GenericServerRunOptions.MaxRequestBodyBytes = 1024 * 1024
			servingCertPath = opts.SecureServing.ServerCert.CertDirectory

			if err := os.WriteFile(path.Join(servingCertPath, "foo.key"), anotherServerKey, 0644); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path.Join(servingCertPath, "foo.crt"), anotherServerCert, 0644); err != nil {
				t.Fatal(err)
			}

			opts.SecureServing.SNICertKeys = []flag.NamedCertKey{{
				Names:    []string{"foo"},
				CertFile: path.Join(servingCertPath, "foo.crt"),
				KeyFile:  path.Join(servingCertPath, "foo.key"),
			}}
		},
	})
	defer tearDownFn()

	// When we run this the second time, we know which one we are expecting.
	_, actualCerts, err := cert.GetServingCertificatesForURL(kubeconfig.Host, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if err := checkServingCerts(anotherServerCert, actualCerts); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(path.Join(servingCertPath, "foo.key"), serverKey, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path.Join(servingCertPath, "foo.crt"), serverCert, 0644); err != nil {
		t.Fatal(err)
	}

	time.Sleep(4 * time.Second)

	_, actualCerts, err = cert.GetServingCertificatesForURL(kubeconfig.Host, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if err := checkServingCerts(serverCert, actualCerts); err != nil {
		t.Fatal(err)
	}
}

func checkServingCerts(expectedBytes []byte, actual [][]byte) error {
	expectedCerts, err := cert.ParseCertsPEM(expectedBytes)
	if err != nil {
		return err
	}
	expected := [][]byte{}
	for _, curr := range expectedCerts {
		currBytes, err := cert.EncodeCertificates(curr)
		if err != nil {
			return err
		}
		expected = append(expected, []byte(strings.TrimSpace(string(currBytes))))
	}

	if len(expected) != len(actual) {
		var certs []string
		for _, a := range actual {
			certs = append(certs, string(a))
		}
		return fmt.Errorf("unexpected number of certs %d vs %d: %v", len(expected), len(actual), strings.Join(certs, "\n"))
	}
	for i := range expected {
		if !bytes.Equal(actual[i], expected[i]) {
			return fmt.Errorf("expected %q, got %q", string(expected[i]), string(actual[i]))
		}
	}
	return nil
}
