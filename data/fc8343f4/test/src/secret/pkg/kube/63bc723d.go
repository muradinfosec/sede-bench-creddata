/*
Copyright 2021 The Kubernetes Authors.

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

package create

import (
	"os"
	"path/filepath"
	"testing"

	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var rsaCertPEM = `-----BEGIN CERTIFICATE-----
MIIB0zCCAX2gAwIBAgIJAI/M7BYjwB+uMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTIwOTEyMjE1MjAyWhcNMTUwOTEyMjE1MjAyWjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANLJ
hPHhITqQbPklG3ibCVxwGMRfp/v4XqhfdQHdcVfHap6NQ5Wok/4xIA+ui35/MmNa
rtNuC+BdZ1tMuVCPFZcCAwEAAaNQME4wHQYDVR0OBBYEFJvKs8RfJaXTH08W+SGv
zQyKn0H8MB8GA1UdIwQYMBaAFJvKs8RfJaXTH08W+SGvzQyKn0H8MAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQEFBQADQQBJlffJHybjDGxRMqaRmDhX0+6v02TUKZsW
r5QuVbpQhH6u+0UgcW0jp9QwpxoPTLTWGXEWBBBurxFwiCBhkQ+V
-----END CERTIFICATE-----
`

var rsaKeyPEM = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBANLJhPHhITqQbPklG3ibCVxwGMRfp/v4XqhfdQHdcVfHap6NQ5Wo
g/9qJl+pb22/UtpHENQja+AYQ1vgvVfUZoOSurqdqUaZGy9s+tmy8EA2/k4mOC2k
3Sl1y5oq+D+DBzl/7bUgtiy1b6TAIcypvbrXLzejeHzLh21aIpuDCW5e9LrG/K9R
ElnREBB+XmRl9vF1kYX6S74cWiLbN9IKh4uWhJ72w/Nicw1wgCdJ51CAtUXqxfZB
CVYZeQc9wtKMyTYsP8vOY5qrQWopGOuMOQY4kIASj6vB9rsMD0Ipxkv5YEP50H6n
crbLONDD/9GkGsBlj0nsSKYG68hLddrPKvRN1gwrX2zUZiL/goWAT/iqYyqDQvJu
R0KJFDDU7/swgeeQSGPiXvS7MNbQbrJ/GjIu4b6T2R==
-----END RSA PRIVATE KEY-----
`

const mismatchRSAKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC/665h55hWD4V2
cvy+S/c5wyGFo62ZVhTKTY9hWyppzK77KX4B3eqjdv66MNuQJQ131Gl1KAnPITCK
POoScfiOCS0LJNRCxr284j05GUcqkNEaDLvOwhjFJk09FQC1LBWMX+KM4eprK+0/
yfF89sFVLfezpEMaD9w6W3aHSONA1awI4Onz82BT2JUyCoFf0LFVFCXHo6d+Jl0y
UC4ySLOSXMVUiaZEjtU444mtK6+keP59o+ATnMGIeDZvtfVa2OBtgN3wvu9TVJ7h
5ujp8rLny2Ax5kVCOnlZ/19lMksVwDfqrXA1Th/Osr/dQc3KyK/WQXBMulj9qcRo
kVDtG0LeoVCKSCkdxShpq+JaAIlslj9Fly4E8r7Eo7VX1lMEVScpATwCnKlZ6Am9
t3h+bp7ycMcdGKkj1tndTnJP5Nvk5IRPe98dv1qngsk8tnybOh9uykLN1/4RZVyG
Sq/R3HWbQ9kw2vcJBbbN6o2bUTHZewwvWA+aOcRNBXunVkDlj/mR9NgD7cktcF4e
0z2UuVgT/bBGkHOtBGNu+DrXDpI6ykWL6LQEFOtYHWUnBnNDpwPv+cKvgI9G/iEh
FwumiwX+bbirzyTVov1ZC6ef9noUbIB/oXQwkt1H3gN8NzqMPhq/pXmRi/Tttfox
nE4axQZLdfUqHsPMiEqHvqovurzm+ac5Z360dZDhYlUVZgu0/yr0XL/Gho8f4FsV
ND2yd2M4GRv2aafuBdLgN09VuDnUkfsUFKZ4EpmB7qFVva6SwwbrBd/uipnjJStL
6yT37jBCCQo9o7wz/FUm5F3Nsq0X3fh6cLPYYEUIJ0uo9JZZICO9MrnFBjIcu5KZ
cqhYWrx6aWWs9ePvkr/+NC2d5AzvnyVc8QYzLtrqw951wvpHMOdoyQh7ezXxW9Zx
hkrGIQqcMm51d7WPXpX/UCUT3xoEyxGqLeJTqTDZwuyDVH2eQhyLNRCiUoCqUxaM
jklLF9oZgSiJXg9SrF2uqNy6ipveI+RwpZVQDs3LmVVxNP+Wj9TyRnm6vc6W2ASN
TDg9R8FLwibvVQZKtzGk16uP/qOjzNrggGEPdg7sOBzJioBRrvAH83cKiI52973m
owjUwgJR/3pw44iUitSZorxEZUC+GuIwwYgHmZKCuzYEr3FnqAgNyjLdzfUtuW9I
7EZYFWQ4wMtSKJwqodjBhYK5dsYWaFmgS0v8uQNtdJzQcYQ8uKq4qwFLsGrJsq9d
xegmgFrpH6kB7cw0fdvTleATu2G3WjL/h882sG1BC0waNcqmM77fJpCVQHuxx7bj
k5ysDHcesCdSjOrCKPt6yMPip2pNVXDZe1bHdtcMOi5ZxLukOGwd2ryf7U6za0OC
MXEeUxboSd3kKzOk8mqBIReK2WjqZ+3yBwmjF3onqgEai7k+H+mwMDCcMzlQWCqe
noNidUj+vagC/ctYsYLjnkFsLXnh453nbi0AwMyTPVv57pS7d2MVQLSDoPthk3h1
FrmN1IwnIqTZc+OZz6dldtm3DpmyU1oTo67VRuPA3yYEVhhDGcb383MsTvjUaQEr
onWi6DUPqDLbSOQK2JhlzL==
-----END PRIVATE KEY-----`

func TestCreateSecretTLS(t *testing.T) {

	validCertTmpDir := t.TempDir()
	validKeyPath, validCertPath := writeKeyPair(validCertTmpDir, rsaKeyPEM, rsaCertPEM, t)

	invalidCertTmpDir := t.TempDir()
	invalidKeyPath, invalidCertPath := writeKeyPair(invalidCertTmpDir, "test", "test", t)

	mismatchCertTmpDir := t.TempDir()
	mismatchKeyPath, mismatchCertPath := writeKeyPair(mismatchCertTmpDir, rsaKeyPEM, mismatchRSAKeyPEM, t)

	tests := map[string]struct {
		tlsSecretName string
		tlsKey        string
		tlsCert       string
		appendHash    bool
		expected      *corev1.Secret
		expectErr     bool
	}{
		"create_secret_tls": {
			tlsSecretName: "foo",
			tlsKey:        validKeyPath,
			tlsCert:       validCertPath,
			expected: &corev1.Secret{
				TypeMeta: metav1.TypeMeta{
					APIVersion: corev1.SchemeGroupVersion.String(),
					Kind:       "Secret",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
				},
				Type: corev1.SecretTypeTLS,
				Data: map[string][]byte{
					corev1.TLSPrivateKeyKey: []byte(rsaKeyPEM),
					corev1.TLSCertKey:       []byte(rsaCertPEM),
				},
			},
			expectErr: false,
		},
		"create_secret_tls_hash": {
			tlsSecretName: "foo",
			tlsKey:        validKeyPath,
			tlsCert:       validCertPath,
			appendHash:    true,
			expected: &corev1.Secret{
				TypeMeta: metav1.TypeMeta{
					APIVersion: corev1.SchemeGroupVersion.String(),
					Kind:       "Secret",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo-272h6tt825",
				},
				Type: corev1.SecretTypeTLS,
				Data: map[string][]byte{
					corev1.TLSPrivateKeyKey: []byte(rsaKeyPEM),
					corev1.TLSCertKey:       []byte(rsaCertPEM),
				},
			},
			expectErr: false,
		},
		"create_secret_invalid_tls": {
			tlsSecretName: "foo",
			tlsKey:        invalidKeyPath,
			tlsCert:       invalidCertPath,
			expectErr:     true,
		},
		"create_secret_mismatch_tls": {
			tlsSecretName: "foo",
			tlsKey:        mismatchKeyPath,
			tlsCert:       mismatchCertPath,
			expectErr:     true,
		},
		"create_invalid_filepath_and_certpath_secret_tls": {
			tlsSecretName: "foo",
			tlsKey:        "testKeyPath",
			tlsCert:       "testCertPath",
			expectErr:     true,
		},
	}

	// Run all the tests
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			secretTLSOptions := CreateSecretTLSOptions{
				Name:       test.tlsSecretName,
				Key:        test.tlsKey,
				Cert:       test.tlsCert,
				AppendHash: test.appendHash,
			}
			secretTLS, err := secretTLSOptions.createSecretTLS()

			if !test.expectErr && err != nil {
				t.Errorf("test %s, unexpected error: %v", name, err)
			}
			if test.expectErr && err == nil {
				t.Errorf("test %s was expecting an error but no error occurred", name)
			}
			if !apiequality.Semantic.DeepEqual(secretTLS, test.expected) {
				t.Errorf("test %s\n expected:\n%#v\ngot:\n%#v", name, test.expected, secretTLS)
			}
		})
	}
}

func write(path, contents string, t *testing.T) {
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to create %v.", path)
	}
	defer f.Close()
	_, err = f.WriteString(contents)
	if err != nil {
		t.Fatalf("Failed to write to %v.", path)
	}
}

func writeKeyPair(tmpDirPath, key, cert string, t *testing.T) (keyPath, certPath string) {
	keyPath = filepath.Join(tmpDirPath, "tls.key")
	certPath = filepath.Join(tmpDirPath, "tls.cert")
	write(keyPath, key, t)
	write(certPath, cert, t)
	return
}
