package provider

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	corecontrollers "github.com/rancher/wrangler/v3/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/v3/pkg/generic/fake"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

const (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAqpXFceskscHq4hxKlJtbAvfh0YF3Wcnjy+k1U2ZxbiHaByLr
Yefh8+qMlZpVPFJ56tb4dF2oRN8iSXJxA6p42i6Ecbhs8iPxN9YfGFxnBzzKwPjY
pnWp96lNg79Ol5E7C/CqiK1HK3xidsdMMe1KzLt4HQAMCkico+68f5q5WFe3xSJ0
YwdeqrRfZO4qJ6q/KDaFJLuvLB2uz4UCJj16CE4qJoPqexlPtuJk+6Vr7M4MdDR4
n3eQ4+pl6O4cZZ9jWDZMjstCt76mcIlFi+TGujG6M+6+jktT3itUzpzg0D0U28Fr
V5u65cHnIrx6r7CKMoN3UE8vBLxuhDWzjduCOcSGgCBUTmfvtiKDvlqOPnS/Iiry
32sqz7CsjXUCi7ZaHUHRZRzGoOcyET+q15OY49hKBzdGgKiCWZ2p8rzAH/enrvhZ
bo56Nxanh9ayj0wKpXSfQhmoiaQDLOOS+yOTLB8U6FGgktlUmUaW2XGUtob6WLGG
mfGYPMAmDrL1ncCWpB9eoHE21wpPFPeLThlv8LcKoXu9aKsvqC2Cu/VCv/4WIygY
6FW4+mjlzKxYtPN+a26Y4eKpOkIXOneo6cQGN1i/Qhz3yLroxyMDegYrVo5GBlri
KzUTkh53LVaWjldSvb4NgtxaoWpg2FV2YfVU4eBVNtE8rf9oH9XatUHethDDF7VV
5aW4UK2UBtNd2tvvN0s9H0532sxKIvyup8+LBl54zIcB+CPyK1Ic0JwK0lR+agv2
DmrwRe9iwYNG+HTWvqPD6sMoZ8N3oLiU87vG1FQIfy4dYBJEjSRHOALGRffZT32t
vPEerWFbXWBAYmQ8AQP2FlFLBOC0JRuBTH9eHZMOhdN/Y/Eki4VHvcdzKvngdHbE
EW64P8sG9W2KaKzKwWASzcejY9SYNtsFiYb4N5c+b7ijx7xxK+5Gq8QwP7fgHcjY
ItM2hKSqXf+Hjik+mp873orKJcNGZk1rYSJbtcsXn5P+jrCUXE2B+tFpHLewlZgM
ZQR4TXlDj9NalK5e8d7Js9GYDWfJAMHAgxSCLT4OAZxbPVAsezF0BWbitHzqt0ua
2lSTtApjsiaJ4a9K2Zei8cqR30WeEkNpMDwj7twA1c2/KJXIm1tOYoJT+B5XavNw
k+/pUvaTca2eC9W0i3fXlGpH+0nYtpL0plc4JUXwpJa0Fl/e4YY+l5DXnh9G1kp3
LPnQxl0NssMWVqOflfkBuCHPOjYsjFRTBiOjvM+R1f9exF8MepFR2JS5MA9OJAr9
ZQtRvnajVgBl/ACFETgpss6x6kNu6Z2NCw9BfD31Q49Yur4Dd8kGWc9cMKegvp6Y
H/WqT7wbP7gcNbYvh4XNuOzfzLoEmksBxvqs1xw5c+/oEMA+INfw+m/jmCfrqKuu
sBJRRnmkKNWuWgCcgT3AD2zgwc84XWwFFauMN+1+te7c5pNTdJBF9mhQLjt3KquB
qJPYaYnt5xa/i9BxRGp9NCsCoZ+lbIFXkuGiCOaIAfKgaIm0jNLZmURc/7wlDtWJ
S2xTOVw2tXvN92EFigug21i60tv+mf5R3arfUT5wsALWnYjYLTe=
-----END RSA PRIVATE KEY-----`
	publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqpXFceskscHq4hxKlJtb
Avfh0YF3Wcnjy+k1U2ZxbiHaByLrUSuP7+TgmLaonsh63mW0xa0ReC7MgFWBf4z0
3S5FWZUs4IpFG6BwrQYYCsANwJPDlUxX42OeB28iZ2J6e/Laai3dv0YkzORlkl8m
kIt9LDDbcdnCR+78I3a6PHE5keO7NRuyNNVZcQ6RQ9F/sQfxzpnGkG0uP1eRwk81
Ii1ZrkVRYnNkuYwH+1FF8R5QYea5T4EN7+co6G3phO6irKAHWkNgX23PUYMSj+qy
Lcf7v+1+UumE8jELoNNY7F1M63XbX0i14qfcodj4H7WQQIj0LU5NkZJUAMmkxOJk
WQIDAQAB
-----END PUBLIC KEY-----`
	privateKeySmall = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAqpXFceskscHq4hxKlJtbAvfh0YF3Wcnjy+k1U2ZxbiHaByLr
pEcz0+eYgXZOuvp37Ob9Gi3hjg3JzNgah6i41I7FYGJZ0wZdH7PAkyHVpOoFyLoY
nJUk27BKS88pM4D0n/Niqv2Ql5HIhmOMgH0rvuc0VtlNrpiIz+52t2M1zaS4Xqc7
GiZHhiKDBA4Ro1B/BqwdmMiIde8ou9QtPY26IO3eUfrJKUtTJJrh+3ix3O5JeqZ4
p2Zo8+hL3W6Gql3IeMKxdGKhX74zeeWTS+VVmgp6L+3+vZhU4AUEUlml7Y1t90bk
C1A09ohgKBn6q4hzwgE1Uo6davOhWTCdHSiynVuJtUoqVoTlHCWTFfIPMqO/ECVe
55zEX5ZpEhWeI7DHzqVkqLUproMtqA+w56vD45GQWxMLyuLhXA6I4EIgo/aPObSw
yA60VyqSD5btQ6wViAEbEpAeKPPRdfsA+zeGhR6T4MVHyfVfSrgZ5IJJVcm9GQDv
kSxQeaJFGjE1WoxAPZ2fTpc44OWmGQXIUIkI3porPbh3RmFwWS0yO/jjW/9qgMEQ
5aT2+vjqFxbGMGX+r89s5ViNHyOJeUUm6IyjW9x/VzM7FhdCYZlSCLWYQR9uNanG
LQdAgs62sKLpNuUWUe7errjdGSai4bb5PBat1dPSgfh4WQ4tJ7SYTEfFqkWfC9uW
5Hw2dN2iQjlr9sJjE7X1Y4282KBptDWnv0+qQl52tCEP+vzyU7OJ3TDv9UR+pJh2
QZUFcF8MTmzA+asBmWAQ8msQM6D1JUoq69Ym5cLCsK8DABNApREIUxSsHvCxz04o
BgRMhgIjtywaqLs9Uns3EwJVGrk6kShSqC4ALuVsJUT/x/DTt3HTtYaHolRxFgZf
BE66R5qX5a3dAuxEphyXvOKfQ1VIezmwYRu6C3K+Y5Nnd1gqx+2gc8wdm6doprSl
blG2LckYnD+jbER+xE671zoEsKoaxg1ODAQxPjhAa5Y+uzxOMM1K+opUdHGyCMEC
bFf6KeOGG9KQdl4W7Y7lt1HscjFqJxxSxEmniN8fFDvDIhkRMuj4fyzgkvPHd4hj
8pfObdNAzcJc4Z1H8Ait8UiQ67DdDiepdohW8kqZ2F9/jwQAO3wBbFEM+n8bXZhc
p+/pSYuwFr0DX3C8P1WdfeDW+9kWHLE5WWO3DOGKkUp3uD/j0SV+P0lUTH8K9Xn2
BUxgQL3ZhnngaJGhcdPqohsZJdzhAEaKlwfzRN+B7g0ICs5JeMdz3qy7SK8LeKB6
EnXqlKvzhToQ/GzdjnRLWj0f9joT7f1Wrh5zsg65A89TJK0uJ6eHFS4rEiAKvz9F
y/SrC8mgZ6RuSkEeY6moJLTbbEbXXsfqvfXA1yZ3P+/OGha+vlBw+R/LtshQrsBS
TGnWZnvKCeBHtQKOsp6Ai2yufg70uBJGYIEID+8+YZ6V9MmiFOQo3LsdoVR7iIiQ
AqtgdDvG9sN/P9FGYRt6jgEkXO+rdTxvAjjhCPXpJtOaobL6bhuFHgOB/3fAUFmZ
L4LtTPn2wINh83sWBiGc39l98cz+PY6v1hkqrx6VIlgNYaqYgrQ=
-----END RSA PRIVATE KEY-----`
	publicKeySmall = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDa2fB2m18KR8whoq1DY51UYMcD
N1rcfAIT6pG7RiXsUoEd6YBA+LbOWe4b7cNZnvvz9S17b+r3bZu1NcE2AEfWbDoI
a5saf1YWiw9zRIRqG4BOBDeySbxbRhP15nWaCBy5fCQ3k1T3mloB0vWcg6IzdDMs
sZe92xd428OAH9tLlQIDAQAB
-----END PUBLIC KEY-----`
)

func TestJWKSEndpoint(t *testing.T) {
	ctlr := gomock.NewController(t)

	tests := map[string]struct {
		secretCache  func() corecontrollers.SecretCache
		expectedCode int
		expectedBody string
	}{
		"jwks success response": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(&v1.Secret{
					Data: map[string][]byte{
						"key.pem": []byte(privateKey),
						"key.pub": []byte(publicKey),
					},
				}, nil)

				return mock
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"keys":[{"kty":"RSA","use":"sig","kid":"key","n":"qpXFceskscHq4hxKlJtbAvfh0YF3Wcnjy-k1U2ZxbiHaByLrUSuP7-TgmLaonsh63mW0xa0ReC7MgFWBf4z03S5FWZUs4IpFG6BwrQYYCsANwJPDlUxX42OeB28iZ2J6e_Laai3dv0YkzORlkl8mkIt9LDDbcdnCR-78I3a6PHE5keO7NRuyNNVZcQ6RQ9F_sQfxzpnGkG0uP1eRwk81Ii1ZrkVRYnNkuYwH-1FF8R5QYea5T4EN7-co6G3phO6irKAHWkNgX23PUYMSj-qyLcf7v-1-UumE8jELoNNY7F1M63XbX0i14qfcodj4H7WQQIj0LU5NkZJUAMmkxOJkWQ","e":"AQAB"}]}`, // contains modulus and exponent for the public key
		},
		"jwks can't get secret": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(nil, errors.New("unexpected error"))

				return mock
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"server_error","error_description":"failed to get secret with public keys"}`,
		},
		"jwks ignores 1024 key": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(&v1.Secret{
					Data: map[string][]byte{
						"key.pem": []byte(privateKeySmall),
						"key.pub": []byte(publicKeySmall),
					},
				}, nil)

				return mock
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"keys":null}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			h := jwksHandler{secretCache: test.secretCache()}

			h.jwksEndpoint(rec, &http.Request{})

			assert.Equal(t, test.expectedCode, rec.Code)
			assert.JSONEq(t, test.expectedBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestGetSigningKey(t *testing.T) {
	ctlr := gomock.NewController(t)
	block, _ := pem.Decode([]byte(privateKey))
	privKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	tests := map[string]struct {
		secretCache func() corecontrollers.SecretCache
		expectedKid string
		expectedKey *rsa.PrivateKey
		expectedErr string
	}{
		"get signing key": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(&v1.Secret{
					Data: map[string][]byte{
						"key.pem": []byte(privateKey),
						"key.pub": []byte(publicKey),
					},
				}, nil)

				return mock
			},
			expectedKid: "key",
			expectedKey: privKey,
		},
		"error retrieving secret": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(nil, errors.New("unexpected error"))

				return mock
			},
			expectedErr: "unexpected error",
		},
		"no signing key in secret": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(&v1.Secret{
					Data: map[string][]byte{
						"key.pub": []byte(publicKey),
					},
				}, nil)

				return mock
			},
			expectedErr: "signing key not found",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			h := jwksHandler{secretCache: test.secretCache()}

			key, kid, err := h.GetSigningKey()

			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedKid, kid)
			assert.Equal(t, test.expectedKey, key)
		})
	}
}

func TestGetPublicKey(t *testing.T) {
	ctlr := gomock.NewController(t)
	block, _ := pem.Decode([]byte(privateKey))
	privKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	tests := map[string]struct {
		secretCache func() corecontrollers.SecretCache
		kid         string
		expectedKey *rsa.PublicKey
		expectedErr string
	}{
		"get signing key": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(&v1.Secret{
					Data: map[string][]byte{
						"key.pem": []byte(privateKey),
						"key.pub": []byte(publicKey),
					},
				}, nil)

				return mock
			},
			kid:         "key",
			expectedKey: &privKey.PublicKey,
		},
		"error retrieving secret": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(nil, errors.New("unexpected error"))

				return mock
			},
			expectedErr: "unexpected error",
		},
		"no signing key in secret": {
			secretCache: func() corecontrollers.SecretCache {
				mock := fake.NewMockCacheInterface[*v1.Secret](ctlr)
				mock.EXPECT().Get(keySecretNamespace, keySecretName).Return(&v1.Secret{
					Data: map[string][]byte{
						"key.pem": []byte(privateKey),
					},
				}, nil)

				return mock
			},
			expectedErr: "public key not found",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			h := jwksHandler{secretCache: test.secretCache()}

			key, err := h.GetPublicKey(test.kid)

			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedKey, key)
		})
	}
}
