/*
Copyright 2014 The Kubernetes Authors.

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

package serviceaccount_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	jose "gopkg.in/go-jose/go-jose.v2"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	typedv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	v1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/keyutil"
	serviceaccountcontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
	"k8s.io/kubernetes/pkg/serviceaccount"
)

const otherPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArXz0QkIG1B5Bj2/W69GH
rsm5e+RC3kE+VTgocge0atqlLBek35tRqLgUi3AcIrBZ/0YctMSWDVcRt5fkhWwe
Lqjj6qvAyNyOkrkBi1NFDpJBjYJtuKHgRhNxXbOzTSNpdSKXTfOkzqv56MwHOP25
yP/NNAODUtr92D5ySI5QX8RbXW+uDn+ixul286PBW/BCrE4tuS88dA0tYJPf8LCu
sqQOwlXYH/rNUg4Pyl9xxhR5DIJR0OzNNfChjw60zieRIt2LfM83fXhwk8IxRGkc
gPZm7ZsipmfbZK2Tkhnpsa4QxDg7zHJPMsB5kxRXW0cQipXcC3baDyN9KBApNXa0
PwIDAQAB
-----END PUBLIC KEY-----`

const rsaPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA249XwEo9k4tM8fMxV7zx
OhcrP+WvXn917koM5Qr2ZXs4vo26e4ytdlrV0bQ9SlcLpQVSYjIxNfhTZdDt+ecI
zshKuv1gKIxbbLQMOuK1eA/4HALyEkFgmS/tleLJrhc65tKPMGD+pKQ/xhmzRuCG
51RoiMgbQxaCyYxGfNLpLAZK9L0Tctv9a0mJmGIYnIOQM4kC1A1I1n3EsXMWmeJU
j7OTh/AjjCnMnkgvKT2tpKxYQ59PgDgU8Ssc7RDSmSkLxnrv+OrN80j6xrw0OjEi
B4Ycr0PqfzZcvy8efTtFQ/Jnc4Bp1zUtFXt7+QeevePtQ2EcyELXE0i63T1CujRM
WwIDAQAB
-----END PUBLIC KEY-----
`

// Obtained by:
//
//  1. Serializing rsaPublicKey as DER
//  2. Taking the SHA256 of the DER bytes
//  3. URLSafe Base64-encoding the sha bytes
const rsaKeyID = "JHJehTTTZlsspKHT-GaJxK7Kd1NQgZJu3fyK6K_QDYU"

// Fake value for testing.
const rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA249XwEo9k4tM8fMxV7zxOhcrP+WvXn917koM5Qr2ZXs4vo26
b9GgRmEF6mC9RKkteNHLlrpMGJZzLdtG+XReIRjViK4QUfzXQGKJsCG3KD/7UlzX
KcFBNE/wJLveFRb04RdkWrd+FJf/edcwRkNY13IXdXXPTfrdbRiathgrmFCp3p3R
UPS9k9rWTuLeOSpHX7Nz3o5h8u4iYfDBkqSBs1fkW/vEBIZrbkmSiK3kkJMwR35M
AYcQ3ATI0MwHWguqNjtF+oCW49I0Mah5kwtOJ4Zyd3lMpUhtxD8uHFMyZ/RxI0La
1VnVFZb4+gppZGEjv1KfBrTZm0Z26q3SWXEGZdyEdfyrumWntvWC9SUKJfQrUfc1
f0/Nl7YW5a22nfHt2Gx7cafcw71MCqAB+5PP0fo9ORsoP02MZwWDK5La5NSl7htg
fDp7QUatpu/a/R7X3+8nP+Xj4JEEU8A2yfCGSd7RKy360bbcV9il3hGZ4JJysCru
s5sgivf8RvG+OL+Zt3mgMDlaq2Hp536SCaWrS6arkMj8Va2zxioDppiIU0lG1pKm
++imBkqfOp63fwewL/SDRSC0ygZy/FyS0/xyQkRGrg7lLnsfHwRhnm60Ad5qNVh0
9FAHlYylE7+RYBipmrpfbFHr6Wx2lYmsaiBgpxJ86QF55f5SCbuS968pw4L9hDN5
UZYJpPvkkrsz28YQ7yxv9pEB74L2vFeWvk75wIkxCPSlYNA0EIMApeuakdZAUe4U
VGiMGD1bofzQsQugeKqBfDIvWyC2QDZy/kGGXq39Q0c6qZsXzpb0sUPWzaZvWPLV
BrpWbYQIeL1Wp7JeiGOhHyD4h/FXNzWZAUXWecU5D2rjA6YjbL2yZBVipZDu8F8l
w/ejV3tjM1WxEHd/PnGjHlfCEN6IisyBtu4p5EOAfx9nG/V1/2dfvx7dAivhX8Us
d2xolnpWecbAfzkpVWiETH44QKTYHmhUCTHe0V0qGNgMvu640Gv8o8A9h8P17pkG
caViL4TIESUS2NSGP3YFJr33z5JKL8nlOiJbIqreEjEamIttCRces5Cdkhry6O47
CyUxh2jtyhvp7b5YyFnvAX2isZX7da7nInVwAKm+iEQUz+NXhLvE9rTu/ajAKUh/
+Jgal+33shNPMn6HSyCywT84fiZU1eaKq4lolC/usFo1kVa9bPgtLnf48tng/Wq9
PJQAahnSE3+DbDWIFR/4Xf2PddKPP0w4nzaHW1EtduB7hF9aFPoAS+D1F9ht4RAE
fWl3kpdHXT+fp66jcC166pZFALcv2w4staOzC9n0Lb24JqAHEDLUuuKW/RByeT+I
J3djgxPAGgjPAAAgOgQmlwI2d8WoPLE4RmBfFhpzi/GD7fagEIggk+UbF6abbwLP
UFbR8qTTaKkrzqqgfGqBodZCFpQ9MHn2r9L0dWlFHZJDF76T51HjhYSq/g4Y3yXD
ikn3WawMgVUOLazYaB2IYjRgVsQwe45mkdYuduoac14pUgIKi9G40Bw/QIU3/SrJ
C285naSjf1L79VprGdPtuMUPeBW6edD9OlipXB/k5fA26/E47pkG
-----END RSA PRIVATE KEY-----
`

// openssl ecparam -name prime256v1 -genkey -noout -out ecdsa256.pem
// Fake value for testing.
const ecdsaPrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIEZmTmUhuanLjPA2CLquXivuwBDHTt5XYwgIr/kA1LtRoAoGCCqGSM49
CSKgmJwZemzoR0oUul5bDN6mjU9UD3X5TELRswfRv9N2HKA/z/NggfTtTErhCoT5
/GB3PCGqSXz/yRrwMLl5m6ngmAFPf3okVs==
-----END EC PRIVATE KEY-----`

// openssl ec -in ecdsa256.pem -pubout -out ecdsa256pub.pem
const ecdsaPublicKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEH6cuzP8XuD5wal6wf9M6xDljTOPL
X2i8uIp/C/ASqiIGUeeKQtX0/IR3qCXyThP/dbCiHrF3v1cuhBOHY8CLVg==
-----END PUBLIC KEY-----`

// Obtained by:
//
//  1. Serializing ecdsaPublicKey as DER
//  2. Taking the SHA256 of the DER bytes
//  3. URLSafe Base64-encoding the sha bytes
const ecdsaKeyID = "SoABiieYuNx4UdqYvZRVeuC6SihxgLrhLy9peHMHpTc"

func getPrivateKey(data string) interface{} {
	key, err := keyutil.ParsePrivateKeyPEM([]byte(data))
	if err != nil {
		panic(fmt.Errorf("unexpected error parsing private key: %v", err))
	}
	return key
}

func getPublicKey(data string) interface{} {
	keys, err := keyutil.ParsePublicKeysPEM([]byte(data))
	if err != nil {
		panic(fmt.Errorf("unexpected error parsing public key: %v", err))
	}
	return keys[0]
}

func TestTokenGenerateAndValidate(t *testing.T) {
	expectedUserName := "system:serviceaccount:test:my-service-account"
	expectedUserUID := "12345"

	// Related API objects
	serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-service-account",
			UID:       "12345",
			Namespace: "test",
		},
	}
	rsaSecret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-rsa-secret",
			Namespace: "test",
		},
	}
	invalidAutoSecret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-rsa-secret",
			Namespace: "test",
			Labels: map[string]string{
				"kubernetes.io/legacy-token-invalid-since": "2022-12-20",
			},
		},
	}
	ecdsaSecret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-ecdsa-secret",
			Namespace: "test",
		},
	}

	// Generate the RSA token
	rsaGenerator, err := serviceaccount.JWTTokenGenerator(serviceaccount.LegacyIssuer, getPrivateKey(rsaPrivateKey))
	if err != nil {
		t.Fatalf("error making generator: %v", err)
	}
	c, pc := serviceaccount.LegacyClaims(*serviceAccount, *rsaSecret)
	rsaToken, err := rsaGenerator.GenerateToken(context.TODO(), c, pc)
	if err != nil {
		t.Fatalf("error generating token: %v", err)
	}
	if len(rsaToken) == 0 {
		t.Fatalf("no token generated")
	}
	rsaSecret.Data = map[string][]byte{
		"token": []byte(rsaToken),
	}

	checkJSONWebSignatureHasKeyID(t, rsaToken, rsaKeyID)

	// Generate RSA token with invalidAutoSecret
	c, pc = serviceaccount.LegacyClaims(*serviceAccount, *invalidAutoSecret)
	invalidAutoSecretToken, err := rsaGenerator.GenerateToken(context.TODO(), c, pc)
	if err != nil {
		t.Fatalf("error generating token: %v", err)
	}
	if len(invalidAutoSecretToken) == 0 {
		t.Fatalf("no token generated")
	}
	invalidAutoSecret.Data = map[string][]byte{
		"token": []byte(invalidAutoSecretToken),
	}

	checkJSONWebSignatureHasKeyID(t, invalidAutoSecretToken, rsaKeyID)

	// Generate the ECDSA token
	ecdsaToken := generateECDSAToken(t, serviceaccount.LegacyIssuer, serviceAccount, ecdsaSecret)

	ecdsaSecret.Data = map[string][]byte{
		"token": []byte(ecdsaToken),
	}

	checkJSONWebSignatureHasKeyID(t, ecdsaToken, ecdsaKeyID)

	ecdsaTokenMalformedIss := generateECDSATokenWithMalformedIss(t, serviceAccount, ecdsaSecret)

	// Generate signer with same keys as RSA signer but different unrecognized issuer
	badIssuerGenerator, err := serviceaccount.JWTTokenGenerator("foo", getPrivateKey(rsaPrivateKey))
	if err != nil {
		t.Fatalf("error making generator: %v", err)
	}
	c, pc = serviceaccount.LegacyClaims(*serviceAccount, *rsaSecret)
	badIssuerToken, err := badIssuerGenerator.GenerateToken(context.TODO(), c, pc)
	if err != nil {
		t.Fatalf("error generating token: %v", err)
	}

	// Generate signer with same keys as RSA signer but different recognized issuer
	differentIssuerGenerator, err := serviceaccount.JWTTokenGenerator("bar", getPrivateKey(rsaPrivateKey))
	if err != nil {
		t.Fatalf("error making generator: %v", err)
	}
	c, pc = serviceaccount.LegacyClaims(*serviceAccount, *rsaSecret)
	differentIssuerToken, err := differentIssuerGenerator.GenerateToken(context.TODO(), c, pc)
	if err != nil {
		t.Fatalf("error generating token: %v", err)
	}

	testCases := map[string]struct {
		Client clientset.Interface
		Keys   []interface{}
		Token  string

		ExpectedErr      bool
		ExpectedOK       bool
		ExpectedUserName string
		ExpectedUserUID  string
		ExpectedGroups   []string
	}{
		"no keys": {
			Token:       rsaToken,
			Client:      nil,
			Keys:        []interface{}{},
			ExpectedErr: true,
			ExpectedOK:  false,
		},
		"invalid keys (rsa)": {
			Token:       rsaToken,
			Client:      nil,
			Keys:        []interface{}{getPublicKey(otherPublicKey), getPublicKey(ecdsaPublicKey)},
			ExpectedErr: true,
			ExpectedOK:  false,
		},
		"invalid keys (ecdsa)": {
			Token:       ecdsaToken,
			Client:      nil,
			Keys:        []interface{}{getPublicKey(otherPublicKey), getPublicKey(rsaPublicKey)},
			ExpectedErr: true,
			ExpectedOK:  false,
		},
		"valid key (rsa)": {
			Token:            rsaToken,
			Client:           nil,
			Keys:             []interface{}{getPublicKey(rsaPublicKey)},
			ExpectedErr:      false,
			ExpectedOK:       true,
			ExpectedUserName: expectedUserName,
			ExpectedUserUID:  expectedUserUID,
			ExpectedGroups:   []string{"system:serviceaccounts", "system:serviceaccounts:test"},
		},
		"valid key, invalid issuer (rsa)": {
			Token:       badIssuerToken,
			Client:      nil,
			Keys:        []interface{}{getPublicKey(rsaPublicKey)},
			ExpectedErr: false,
			ExpectedOK:  false,
		},
		"valid key, different issuer (rsa)": {
			Token:            differentIssuerToken,
			Client:           nil,
			Keys:             []interface{}{getPublicKey(rsaPublicKey)},
			ExpectedErr:      false,
			ExpectedOK:       true,
			ExpectedUserName: expectedUserName,
			ExpectedUserUID:  expectedUserUID,
			ExpectedGroups:   []string{"system:serviceaccounts", "system:serviceaccounts:test"},
		},
		"valid key (ecdsa)": {
			Token:            ecdsaToken,
			Client:           nil,
			Keys:             []interface{}{getPublicKey(ecdsaPublicKey)},
			ExpectedErr:      false,
			ExpectedOK:       true,
			ExpectedUserName: expectedUserName,
			ExpectedUserUID:  expectedUserUID,
			ExpectedGroups:   []string{"system:serviceaccounts", "system:serviceaccounts:test"},
		},
		"rotated keys (rsa)": {
			Token:            rsaToken,
			Client:           nil,
			Keys:             []interface{}{getPublicKey(otherPublicKey), getPublicKey(ecdsaPublicKey), getPublicKey(rsaPublicKey)},
			ExpectedErr:      false,
			ExpectedOK:       true,
			ExpectedUserName: expectedUserName,
			ExpectedUserUID:  expectedUserUID,
			ExpectedGroups:   []string{"system:serviceaccounts", "system:serviceaccounts:test"},
		},
		"rotated keys (ecdsa)": {
			Token:            ecdsaToken,
			Client:           nil,
			Keys:             []interface{}{getPublicKey(otherPublicKey), getPublicKey(rsaPublicKey), getPublicKey(ecdsaPublicKey)},
			ExpectedErr:      false,
			ExpectedOK:       true,
			ExpectedUserName: expectedUserName,
			ExpectedUserUID:  expectedUserUID,
			ExpectedGroups:   []string{"system:serviceaccounts", "system:serviceaccounts:test"},
		},
		"valid lookup": {
			Token:            rsaToken,
			Client:           fake.NewSimpleClientset(serviceAccount, rsaSecret, ecdsaSecret),
			Keys:             []interface{}{getPublicKey(rsaPublicKey)},
			ExpectedErr:      false,
			ExpectedOK:       true,
			ExpectedUserName: expectedUserName,
			ExpectedUserUID:  expectedUserUID,
			ExpectedGroups:   []string{"system:serviceaccounts", "system:serviceaccounts:test"},
		},
		"invalid secret lookup": {
			Token:       rsaToken,
			Client:      fake.NewSimpleClientset(serviceAccount),
			Keys:        []interface{}{getPublicKey(rsaPublicKey)},
			ExpectedErr: true,
			ExpectedOK:  false,
		},
		"invalid serviceaccount lookup": {
			Token:       rsaToken,
			Client:      fake.NewSimpleClientset(rsaSecret, ecdsaSecret),
			Keys:        []interface{}{getPublicKey(rsaPublicKey)},
			ExpectedErr: true,
			ExpectedOK:  false,
		},
		"secret is marked as invalid": {
			Token:       invalidAutoSecretToken,
			Client:      fake.NewSimpleClientset(serviceAccount, invalidAutoSecret),
			Keys:        []interface{}{getPublicKey(rsaPublicKey)},
			ExpectedErr: true,
		},
		"malformed iss": {
			Token:       ecdsaTokenMalformedIss,
			Client:      nil,
			Keys:        []interface{}{getPublicKey(ecdsaPublicKey)},
			ExpectedErr: false,
			ExpectedOK:  false,
		},
	}

	for k, tc := range testCases {
		auds := authenticator.Audiences{"api"}
		getter := serviceaccountcontroller.NewGetterFromClient(
			tc.Client,
			v1listers.NewSecretLister(newIndexer(func(namespace, name string) (interface{}, error) {
				return tc.Client.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
			})),
			v1listers.NewServiceAccountLister(newIndexer(func(namespace, name string) (interface{}, error) {
				return tc.Client.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), name, metav1.GetOptions{})
			})),
			v1listers.NewPodLister(newIndexer(func(namespace, name string) (interface{}, error) {
				return tc.Client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
			})),
			v1listers.NewNodeLister(newIndexer(func(_, name string) (interface{}, error) {
				return tc.Client.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
			})),
		)
		var secretsWriter typedv1core.SecretsGetter
		if tc.Client != nil {
			secretsWriter = tc.Client.CoreV1()
		}
		validator, err := serviceaccount.NewLegacyValidator(tc.Client != nil, getter, secretsWriter)
		if err != nil {
			t.Fatalf("While creating legacy validator, err: %v", err)
		}
		staticKeysGetter, err := serviceaccount.StaticPublicKeysGetter(tc.Keys)
		if err != nil {
			t.Fatal(err)
		}
		keysGetter := &keyIDPrefixer{PublicKeysGetter: staticKeysGetter}

		authn := serviceaccount.JWTTokenAuthenticator([]string{serviceaccount.LegacyIssuer, "bar"}, keysGetter, auds, validator)

		// An invalid, non-JWT token should always fail
		ctx := authenticator.WithAudiences(context.TODO(), auds)
		if _, ok, err := authn.AuthenticateToken(ctx, "invalid token"); err != nil || ok {
			t.Errorf("%s: Expected err=nil, ok=false for non-JWT token", k)
			continue
		}

		if tc.ExpectedOK {
			// if authentication is otherwise expected to succeed, demonstrate changing key ids makes it fail
			keysGetter.keyIDPrefix = "bogus"
			if _, ok, err := authn.AuthenticateToken(ctx, tc.Token); err == nil || !strings.Contains(err.Error(), "no keys found") || ok {
				t.Errorf("%s: Expected err containing 'no keys found', ok=false when key lookup by ID fails", k)
				continue
			}
			keysGetter.keyIDPrefix = ""
		}

		resp, ok, err := authn.AuthenticateToken(ctx, tc.Token)
		if (err != nil) != tc.ExpectedErr {
			t.Errorf("%s: Expected error=%v, got %v", k, tc.ExpectedErr, err)
			continue
		}

		if ok != tc.ExpectedOK {
			t.Errorf("%s: Expected ok=%v, got %v", k, tc.ExpectedOK, ok)
			continue
		}

		if err != nil || !ok {
			continue
		}

		if resp.User.GetName() != tc.ExpectedUserName {
			t.Errorf("%s: Expected username=%v, got %v", k, tc.ExpectedUserName, resp.User.GetName())
			continue
		}
		if resp.User.GetUID() != tc.ExpectedUserUID {
			t.Errorf("%s: Expected userUID=%v, got %v", k, tc.ExpectedUserUID, resp.User.GetUID())
			continue
		}
		if !reflect.DeepEqual(resp.User.GetGroups(), tc.ExpectedGroups) {
			t.Errorf("%s: Expected groups=%v, got %v", k, tc.ExpectedGroups, resp.User.GetGroups())
			continue
		}
	}
}

type keyIDPrefixer struct {
	serviceaccount.PublicKeysGetter
	keyIDPrefix string
}

func (k *keyIDPrefixer) GetPublicKeys(ctx context.Context, keyIDHint string) []serviceaccount.PublicKey {
	if k.keyIDPrefix == "" {
		return k.PublicKeysGetter.GetPublicKeys(context.TODO(), keyIDHint)
	}
	if keyIDHint != "" {
		keyIDHint = k.keyIDPrefix + keyIDHint
	}
	var retval []serviceaccount.PublicKey
	for _, key := range k.PublicKeysGetter.GetPublicKeys(context.TODO(), keyIDHint) {
		key.KeyID = k.keyIDPrefix + key.KeyID
		retval = append(retval, key)
	}
	return retval
}

func checkJSONWebSignatureHasKeyID(t *testing.T, jwsString string, expectedKeyID string) {
	jws, err := jose.ParseSigned(jwsString)
	if err != nil {
		t.Fatalf("Error checking for key ID: couldn't parse token: %v", err)
	}

	if jws.Signatures[0].Header.KeyID != expectedKeyID {
		t.Errorf("Token %q has the wrong KeyID (got %q, want %q)", jwsString, jws.Signatures[0].Header.KeyID, expectedKeyID)
	}
}

func newIndexer(get func(namespace, name string) (interface{}, error)) cache.Indexer {
	return &fakeIndexer{get: get}
}

type fakeIndexer struct {
	cache.Indexer
	get func(namespace, name string) (interface{}, error)
}

func (f *fakeIndexer) GetByKey(key string) (interface{}, bool, error) {
	parts := strings.SplitN(key, "/", 2)
	namespace := parts[0]
	name := ""
	// implies the key does not contain a / (this is a cluster-scoped object)
	if len(parts) == 1 {
		name = parts[0]
		namespace = ""
	}
	if len(parts) == 2 {
		name = parts[1]
	}
	obj, err := f.get(namespace, name)
	return obj, err == nil, err
}

func generateECDSAToken(t *testing.T, iss string, serviceAccount *v1.ServiceAccount, ecdsaSecret *v1.Secret) string {
	t.Helper()

	ecdsaGenerator, err := serviceaccount.JWTTokenGenerator(iss, getPrivateKey(ecdsaPrivateKey))
	if err != nil {
		t.Fatalf("error making generator: %v", err)
	}
	c, pc := serviceaccount.LegacyClaims(*serviceAccount, *ecdsaSecret)
	ecdsaToken, err := ecdsaGenerator.GenerateToken(context.TODO(), c, pc)
	if err != nil {
		t.Fatalf("error generating token: %v", err)
	}
	if len(ecdsaToken) == 0 {
		t.Fatalf("no token generated")
	}

	return ecdsaToken
}

func generateECDSATokenWithMalformedIss(t *testing.T, serviceAccount *v1.ServiceAccount, ecdsaSecret *v1.Secret) string {
	t.Helper()

	ecdsaToken := generateECDSAToken(t, "panda", serviceAccount, ecdsaSecret)

	ecdsaTokenJWS, err := jose.ParseSigned(ecdsaToken)
	if err != nil {
		t.Fatal(err)
	}

	dataFullSerialize := map[string]any{}
	if err := json.Unmarshal([]byte(ecdsaTokenJWS.FullSerialize()), &dataFullSerialize); err != nil {
		t.Fatal(err)
	}

	dataFullSerialize["malformed_iss"] = "." + base64.RawURLEncoding.EncodeToString([]byte(`{"iss":"bar"}`)) + "."

	out, err := json.Marshal(dataFullSerialize)
	if err != nil {
		t.Fatal(err)
	}

	return string(out)
}

func TestStaticPublicKeysGetter(t *testing.T) {
	ecPrivate := getPrivateKey(ecdsaPrivateKey)
	ecPublic := getPublicKey(ecdsaPublicKey)
	rsaPublic := getPublicKey(rsaPublicKey)

	testcases := []struct {
		Name       string
		Keys       []interface{}
		ExpectErr  bool
		ExpectKeys []serviceaccount.PublicKey
	}{
		{
			Name:       "empty",
			Keys:       nil,
			ExpectKeys: []serviceaccount.PublicKey{},
		},
		{
			Name: "simple",
			Keys: []interface{}{ecPublic, rsaPublic},
			ExpectKeys: []serviceaccount.PublicKey{
				{KeyID: "SoABiieYuNx4UdqYvZRVeuC6SihxgLrhLy9peHMHpTc", PublicKey: ecPublic},
				{KeyID: "JHJehTTTZlsspKHT-GaJxK7Kd1NQgZJu3fyK6K_QDYU", PublicKey: rsaPublic},
			},
		},
		{
			Name: "private --> public",
			Keys: []interface{}{ecPrivate},
			ExpectKeys: []serviceaccount.PublicKey{
				{KeyID: "SoABiieYuNx4UdqYvZRVeuC6SihxgLrhLy9peHMHpTc", PublicKey: ecPublic},
			},
		},
		{
			Name:      "invalid",
			Keys:      []interface{}{"bogus"},
			ExpectErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			getter, err := serviceaccount.StaticPublicKeysGetter(tc.Keys)
			if tc.ExpectErr {
				if err == nil {
					t.Fatal("expected construction error, got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected construction error: %v", err)
			}

			bogusKeys := getter.GetPublicKeys(context.TODO(), "bogus")
			if len(bogusKeys) != 0 {
				t.Fatalf("unexpected bogus keys: %#v", bogusKeys)
			}

			allKeys := getter.GetPublicKeys(context.TODO(), "")
			if !reflect.DeepEqual(tc.ExpectKeys, allKeys) {
				t.Fatalf("unexpected keys: %#v", allKeys)
			}
			for _, key := range allKeys {
				keysByID := getter.GetPublicKeys(context.TODO(), key.KeyID)
				if len(keysByID) != 1 {
					t.Fatalf("expected 1 key for id %s, got %d", key.KeyID, len(keysByID))
				}
				if !reflect.DeepEqual(key, keysByID[0]) {
					t.Fatalf("unexpected key for id %s", key.KeyID)
				}
			}
		})
	}
}
