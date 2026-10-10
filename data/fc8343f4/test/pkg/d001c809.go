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

package serviceaccount_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/google/go-cmp/cmp"
	jose "gopkg.in/go-jose/go-jose.v2"

	"k8s.io/kubernetes/pkg/routes"
	"k8s.io/kubernetes/pkg/serviceaccount"
)

const (
	exampleIssuer = "https://issuer.example.com"
)

func setupServer(t *testing.T, iss string, keys serviceaccount.PublicKeysGetter) (*httptest.Server, string) {
	t.Helper()

	c := restful.NewContainer()
	s := httptest.NewServer(c)

	// JWKS needs to be https, so swap that for the test
	jwksURI, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	jwksURI.Scheme = "https"
	jwksURI.Path = serviceaccount.JWKSPath

	md, err := serviceaccount.NewOpenIDMetadataProvider(
		iss, jwksURI.String(), "", keys)
	if err != nil {
		t.Fatal(err)
	}

	srv := routes.NewOpenIDMetadataServer(md)
	srv.Install(c)

	return s, jwksURI.String()
}

var defaultKeys = []interface{}{getPublicKey(rsaPublicKey), getPublicKey(ecdsaPublicKey)}

// Configuration is an OIDC configuration, including most but not all required fields.
// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
type Configuration struct {
	Issuer        string   `json:"issuer"`
	JWKSURI       string   `json:"jwks_uri"`
	ResponseTypes []string `json:"response_types_supported"`
	SigningAlgs   []string `json:"id_token_signing_alg_values_supported"`
	SubjectTypes  []string `json:"subject_types_supported"`
}

type proxyKeyGetter struct {
	serviceaccount.PublicKeysGetter
	listeners []serviceaccount.Listener
}

func (p *proxyKeyGetter) AddListener(listener serviceaccount.Listener) {
	p.listeners = append(p.listeners, listener)
	p.PublicKeysGetter.AddListener(listener)
}

func TestServeConfiguration(t *testing.T) {
	ecKeysGetter, err := serviceaccount.StaticPublicKeysGetter([]interface{}{getPublicKey(ecdsaPublicKey)})
	if err != nil {
		t.Fatal(err)
	}
	rsaKeysGetter, err := serviceaccount.StaticPublicKeysGetter([]interface{}{getPublicKey(rsaPublicKey)})
	if err != nil {
		t.Fatal(err)
	}
	keysGetter := &proxyKeyGetter{PublicKeysGetter: ecKeysGetter}
	s, jwksURI := setupServer(t, exampleIssuer, keysGetter)
	defer s.Close()

	wantEC := Configuration{
		Issuer:        exampleIssuer,
		JWKSURI:       jwksURI,
		ResponseTypes: []string{"id_token"},
		SubjectTypes:  []string{"public"},
		SigningAlgs:   []string{"ES256"},
	}
	wantRSA := Configuration{
		Issuer:        exampleIssuer,
		JWKSURI:       jwksURI,
		ResponseTypes: []string{"id_token"},
		SubjectTypes:  []string{"public"},
		SigningAlgs:   []string{"RS256"},
	}
	reqURL := s.URL + "/.well-known/openid-configuration"

	expectConfiguration(t, reqURL, wantEC)

	// modify the underlying keys, expect the same response
	keysGetter.PublicKeysGetter = rsaKeysGetter
	expectConfiguration(t, reqURL, wantEC)

	// notify the metadata the keys changed, expected a modified response
	for _, listener := range keysGetter.listeners {
		listener.Enqueue()
	}
	expectConfiguration(t, reqURL, wantRSA)
}

func expectConfiguration(t *testing.T, reqURL string, want Configuration) {
	resp, err := http.Get(reqURL)
	if err != nil {
		t.Fatalf("Get(%s) = %v, %v want: <response>, <nil>", reqURL, resp, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Get(%s) = %v, _ want: %v, _", reqURL, resp.StatusCode, http.StatusOK)
	}

	if got, want := resp.Header.Get("Content-Type"), "application/json"; got != want {
		t.Errorf("Get(%s) Content-Type = %q, _ want: %q, _", reqURL, got, want)
	}
	if got, want := resp.Header.Get("Cache-Control"), "public, max-age=3600"; got != want {
		t.Errorf("Get(%s) Cache-Control = %q, _ want: %q, _", reqURL, got, want)
	}

	var got Configuration
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("Decode(_) = %v, want: <nil>", err)
	}

	if !cmp.Equal(want, got) {
		t.Errorf("unexpected diff in received configuration (-want, +got):%s",
			cmp.Diff(want, got))
	}
}

func TestServeKeys(t *testing.T) {
	wantPubRSA := getPublicKey(rsaPublicKey).(*rsa.PublicKey)
	wantPubECDSA := getPublicKey(ecdsaPublicKey).(*ecdsa.PublicKey)

	alternateGetter, err := serviceaccount.StaticPublicKeysGetter([]interface{}{wantPubRSA})
	if err != nil {
		t.Fatal(err)
	}

	var serveKeysTests = []struct {
		Name               string
		Keys               []interface{}
		WantKeys           []jose.JSONWebKey
		updatedKeysGetter  serviceaccount.PublicKeysGetter
		WantKeysPostUpdate []jose.JSONWebKey
	}{
		{
			Name: "configured public keys",
			Keys: []interface{}{
				getPublicKey(rsaPublicKey),
				getPublicKey(ecdsaPublicKey),
			},
			WantKeys: []jose.JSONWebKey{
				{
					Algorithm:                   "RS256",
					Key:                         wantPubRSA,
					KeyID:                       rsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
				{
					Algorithm:                   "ES256",
					Key:                         wantPubECDSA,
					KeyID:                       ecdsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
			},
		},
		{
			Name: "only publishes public keys",
			Keys: []interface{}{
				getPrivateKey(rsaPrivateKey),
				getPrivateKey(ecdsaPrivateKey),
			},
			WantKeys: []jose.JSONWebKey{
				{
					Algorithm:                   "RS256",
					Key:                         wantPubRSA,
					KeyID:                       rsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
				{
					Algorithm:                   "ES256",
					Key:                         wantPubECDSA,
					KeyID:                       ecdsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
			},
		},
		{
			Name: "configured public keys reacting to update",
			Keys: []interface{}{
				getPublicKey(rsaPublicKey),
				getPublicKey(ecdsaPublicKey),
			},
			WantKeys: []jose.JSONWebKey{
				{
					Algorithm:                   "RS256",
					Key:                         wantPubRSA,
					KeyID:                       rsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
				{
					Algorithm:                   "ES256",
					Key:                         wantPubECDSA,
					KeyID:                       ecdsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
			},
			updatedKeysGetter: alternateGetter,
			WantKeysPostUpdate: []jose.JSONWebKey{
				{
					Algorithm:                   "RS256",
					Key:                         wantPubRSA,
					KeyID:                       rsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
			},
		},
		{
			Name: "configured public keys reacting to update while excluding keys",
			Keys: []interface{}{
				getPublicKey(rsaPublicKey),
				getPublicKey(ecdsaPublicKey),
			},
			WantKeys: []jose.JSONWebKey{
				{
					Algorithm:                   "RS256",
					Key:                         wantPubRSA,
					KeyID:                       rsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
				{
					Algorithm:                   "ES256",
					Key:                         wantPubECDSA,
					KeyID:                       ecdsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
			},
			updatedKeysGetter: dummyPublicKeyGetter{
				keys: []serviceaccount.PublicKey{
					{
						KeyID:                    rsaKeyID,
						PublicKey:                wantPubRSA,
						ExcludeFromOIDCDiscovery: true,
					},
					{
						KeyID:                    ecdsaKeyID,
						PublicKey:                wantPubECDSA,
						ExcludeFromOIDCDiscovery: false,
					},
				},
			},
			WantKeysPostUpdate: []jose.JSONWebKey{
				{
					Algorithm:                   "ES256",
					Key:                         wantPubECDSA,
					KeyID:                       ecdsaKeyID,
					Use:                         "sig",
					Certificates:                []*x509.Certificate{},
					CertificateThumbprintSHA1:   []uint8{},
					CertificateThumbprintSHA256: []uint8{},
				},
			},
		},
	}

	for _, tt := range serveKeysTests {
		t.Run(tt.Name, func(t *testing.T) {
			initialKeysGetter, err := serviceaccount.StaticPublicKeysGetter(tt.Keys)
			if err != nil {
				t.Fatal(err)
			}
			keysGetter := &proxyKeyGetter{PublicKeysGetter: initialKeysGetter}
			s, _ := setupServer(t, exampleIssuer, keysGetter)
			defer s.Close()

			reqURL := s.URL + "/openid/v1/jwks"
			expectKeys(t, reqURL, tt.WantKeys)

			if tt.updatedKeysGetter != nil {
				// modify the underlying keys, expect the same response
				keysGetter.PublicKeysGetter = tt.updatedKeysGetter
				expectKeys(t, reqURL, tt.WantKeys)

				// notify the metadata the keys changed, expected a modified response
				for _, listener := range keysGetter.listeners {
					listener.Enqueue()
				}
				expectKeys(t, reqURL, tt.WantKeysPostUpdate)
			}
		})
	}
}
func expectKeys(t *testing.T, reqURL string, wantKeys []jose.JSONWebKey) {
	resp, err := http.Get(reqURL)
	if err != nil {
		t.Fatalf("Get(%s) = %v, %v want: <response>, <nil>", reqURL, resp, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Get(%s) = %v, _ want: %v, _", reqURL, resp.StatusCode, http.StatusOK)
	}

	if got, want := resp.Header.Get("Content-Type"), "application/jwk-set+json"; got != want {
		t.Errorf("Get(%s) Content-Type = %q, _ want: %q, _", reqURL, got, want)
	}
	if got, want := resp.Header.Get("Cache-Control"), "public, max-age=3600"; got != want {
		t.Errorf("Get(%s) Cache-Control = %q, _ want: %q, _", reqURL, got, want)
	}

	ks := &jose.JSONWebKeySet{}
	if err := json.NewDecoder(resp.Body).Decode(ks); err != nil {
		t.Fatalf("Decode(_) = %v, want: <nil>", err)
	}

	bigIntComparer := cmp.Comparer(
		func(x, y *big.Int) bool {
			return x.Cmp(y) == 0
		})
	if !cmp.Equal(wantKeys, ks.Keys, bigIntComparer) {
		t.Errorf("unexpected diff in JWKS keys (-want, +got): %v",
			cmp.Diff(wantKeys, ks.Keys, bigIntComparer))
	}
}

func TestURLBoundaries(t *testing.T) {
	keysGetter, err := serviceaccount.StaticPublicKeysGetter(defaultKeys)
	if err != nil {
		t.Fatal(err)
	}
	s, _ := setupServer(t, exampleIssuer, keysGetter)
	defer s.Close()

	for _, tt := range []struct {
		Name   string
		Path   string
		WantOK bool
	}{
		{"OIDC config path", "/.well-known/openid-configuration", true},
		{"JWKS path", "/openid/v1/jwks", true},
		{"well-known", "/.well-known", false},
		{"subpath", "/openid/v1/jwks/foo", false},
		{"query", "/openid/v1/jwks?format=yaml", true},
		{"fragment", "/openid/v1/jwks#issuer", true},
	} {
		t.Run(tt.Name, func(t *testing.T) {
			resp, err := http.Get(s.URL + tt.Path)
			if err != nil {
				t.Fatal(err)
			}

			if tt.WantOK && (resp.StatusCode != http.StatusOK) {
				t.Errorf("Get(%v)= %v, want %v", tt.Path, resp.StatusCode, http.StatusOK)
			}
			if !tt.WantOK && (resp.StatusCode != http.StatusNotFound) {
				t.Errorf("Get(%v)= %v, want %v", tt.Path, resp.StatusCode, http.StatusNotFound)
			}
		})
	}
}

func TestNewOpenIDMetadata(t *testing.T) {
	cases := []struct {
		name            string
		issuerURL       string
		jwksURI         string
		externalAddress string
		keys            []interface{}
		wantConfig      string
		wantKeyset      string
		err             bool
	}{
		{
			name:       "valid inputs",
			issuerURL:  exampleIssuer,
			jwksURI:    exampleIssuer + serviceaccount.JWKSPath,
			keys:       defaultKeys,
			wantConfig: `{"issuer":"https://issuer.example.com","jwks_uri":"https://issuer.example.com/openid/v1/jwks","response_types_supported":["id_token"],"subject_types_supported":["public"],"id_token_signing_alg_values_supported":["ES256","RS256"]}`,
			wantKeyset: `{"bzct":[{"kuq":"goe","yee":"YIA","wyu":"ORJvnSNGXueenDVY-RpMbD8Qj5VDnZBh0umF4J_SFRA","cqv":"FC842","d":"908RcNw5p6jP9zUhZ5zcKgxhT-WtQt684ryZ3Qn1SHr1ga58e4ejhnrI8fU1IrpIiNTTXmAoDitWZcVq-qvMqspZvo8yWZxaxAJQLmC7hM_3VTCeNgQyuJ_mcuHAjzs01cYDBVD-aIG_xbyoWpBP20YyqIaiXrcTqNgKyFEfMOJV3Z2Yewn8f6dOhQRAwYBFS9uA7D9J9f2BlBLAbuNLi3BIh_HjjSyLpovpNL3hsNbOE14SnDtK8Wai0CAZbKxDdwsa-OyP83j0arw0AyCxW9Rdr4TcubFplj9ceGzOT_Adm5Zf2gQwRTp6-ZrgzbWvM4JcgSQKK5y02E7TwjNIQe","h":"UXUC"},{"wqf":"zxe","lov":"OB","ytk":"BcPRyemSsDj2AcwEwHCNrvY6NiwjcNqqJf6gmDJHsIs","bsm":"J-703","jqm":"OY194","v":"X8bqvY1YmZ9peq7si1Y9lLfcCHWDT5l3kEk_W_WYprS","c":"EuUjunNU8KzTu4kr4p1A_9Xpyp9vx23LFgEIr5DBa7V"}]}`,
		},
		{
			name:      "valid inputs, default JWKSURI to external address",
			issuerURL: exampleIssuer,
			jwksURI:   "",
			// We expect host + port, no scheme, when API server calculates ExternalAddress.
			externalAddress: "192.0.2.1:80",
			keys:            defaultKeys,
			wantConfig:      `{"issuer":"https://issuer.example.com","jwks_uri":"https://192.0.2.1:80/openid/v1/jwks","response_types_supported":["id_token"],"subject_types_supported":["public"],"id_token_signing_alg_values_supported":["ES256","RS256"]}`,
			wantKeyset:      `{"zisp":[{"loy":"pnp","var":"SIM","job":"ZCBqeSQXIghjnOVF-OxOaE3Fk0HUtBMt4wrD2G_AEBD","zdb":"AF329","k":"694HzEk8j4xA0nUvN4hiSeymU-QwBc377qtO5Yl3DHr8sq80f1kjomaT6dW9YgiGpLXGDtFqEfvSEjNn-rsGfhsHkl9uXJchgPIFIbV7zB_6MVRfYnKsdO_quzPYvym50uAUTTX-vCK_rzxgSiYH36MilHryBigHqNsNeGWrNXBR6K3Prtx2a6tPwTVWmRVFT0vH7Q2M7i1GoXNRswLDx2OUa_RkjJgSpfypEK0hgBcCD08BeViA4Tpk7MXHdFwToqfp-DuI47r0djg6NlHqL3Adk0IfkzUwnn2ygIqIC_Pna3Io1mTcYOk2-YaxceXzY7MgkRWRQ8r68N2ItfOVEn","i":"WBGZ"},{"xrl":"npr","zel":"HD","har":"UwDGfpkHsHc9HkhNaITShsV9AjzbmSnpTr4snHJCxNg","uot":"J-610","fad":"NV040","v":"T5wvzF6HfL3rax4wp0X1bQdvBQVLI4j4oOo_B_HCifE","m":"DcVngxCO1XmEh9th1a7G_5Ayzq1fd26IXvXSs0SAw0U"}]}`,
		},
		{
			name:       "valid inputs, IP addresses instead of domains",
			issuerURL:  "https://192.0.2.1:80",
			jwksURI:    "https://192.0.2.1:80" + serviceaccount.JWKSPath,
			keys:       defaultKeys,
			wantConfig: `{"issuer":"https://192.0.2.1:80","jwks_uri":"https://192.0.2.1:80/openid/v1/jwks","response_types_supported":["id_token"],"subject_types_supported":["public"],"id_token_signing_alg_values_supported":["ES256","RS256"]}`,
			wantKeyset: `{"nhrr":[{"mpn":"plx","gkr":"UCT","bgu":"SOWrfYBWMfeezMPB-PeCsM7Ki1AQyRFk4ikO3M_AXNM","cto":"VM594","k":"328SvRt9x1nQ8hMhG2qvYnhiF-MoDq767hiS3Lj8OMp5jm37m4kfsosC7vK2EpjBaJWCSlEgTuyYCaXq-okYhlgGen4cQPvzeMMGTeA3pW_5LFOgXmQccW_zssXCqvv41iKMIVO-uEN_ihfzIqBG19FkbSlrKdyCcHoWiSRsTMSH1D4Nypj6a1nArZNCjTEDJ1mU1S5K6v5PeOYJboPGr9UCf_PgtAlNbrfvXX2rgYrXN91RpQgM1Maq0SIWkLiSdcez-JwQ40c6uzn0KdYoN5Zxs9CkufLlmc9qqQlHE_Qax3Rf8sHnSRm8-ShkfwItR7SpaDVGH9c08H3KrpWWVu","l":"PETL"},{"wtx":"fbi","nre":"OY","owj":"MaNDwupVqHb3LoyPrFBUdtQ6GzgghWnfRw0mgOECaWa","ojs":"F-220","qml":"CH846","d":"T9tgmN8PfD7vpi7cx0R1fBgcHROID6f6mTp_M_XLbrU","p":"AiOpcbZS6GaWo1ty5b7S_2Vibk2pt75BZwIGo2GQp2P"}]}`,
		},
		{
			name:       "response only contains public keys, even when private keys are provided",
			issuerURL:  exampleIssuer,
			jwksURI:    exampleIssuer + serviceaccount.JWKSPath,
			keys:       []interface{}{getPrivateKey(rsaPrivateKey), getPrivateKey(ecdsaPrivateKey)},
			wantConfig: `{"issuer":"https://issuer.example.com","jwks_uri":"https://issuer.example.com/openid/v1/jwks","response_types_supported":["id_token"],"subject_types_supported":["public"],"id_token_signing_alg_values_supported":["ES256","RS256"]}`,
			wantKeyset: `{"xeyj":[{"gpx":"lip","wmm":"UOR","jna":"LRDzdUFOYujfnDNQ-BnKrB7Lj3FEvPIi7cmC4N_FHDP","hel":"MS596","i":"995YhTl3v7jW2aHeZ3lwSngeY-RkAx727qfC5Ir8YPw2xz48q3gxhwhS4eH2RljTxFXKIhJwNglHFyZj-uhUkpeUhz7xFXdkhSJORuY9nD_0RDShIyVhiB_ktxNOohy66uDYNPY-iVC_kvrdPtVR44VvwWiqOgyGfMhAmDHwQARY6H9Ryuv6h0wGxTDOxUSWE4sF1V3D9q8TlXUXpyPQu2ZSn_FlrUyBlhawOU7csOsKE11JyAlA3Ggo5BJDhZnFnnmm-YfI57p0ujr6EdXoN7Fya3CwpfAvgw4mmNsCE_Pul1Zw7iJyNJq3-DphykBqW2LkgWEMR0p79C0DdmXEPy","c":"MBME"},{"wfl":"ean","jwo":"GX","blj":"IeIQyiwSkEi5TvpKhPTOcuM4FndpkEquVf2cwNSKxYh","djt":"N-057","teh":"HH699","l":"V9iydZ6CfG9cvk3bn0D4qMqeVXVHM5l0eXu_A_XXfiL","l":"SpYokzNW6BoLk1ug9w9L_5Pqpq9st68HWgJAi2LTm5W"}]}`,
		},
		{
			name:      "issuer missing https",
			issuerURL: "http://issuer.example.com",
			jwksURI:   exampleIssuer + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "issuer missing scheme",
			issuerURL: "issuer.example.com",
			jwksURI:   exampleIssuer + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "issuer includes query",
			issuerURL: "https://issuer.example.com?foo=bar",
			jwksURI:   exampleIssuer + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "issuer includes fragment",
			issuerURL: "https://issuer.example.com#baz",
			jwksURI:   exampleIssuer + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "issuer includes query and fragment",
			issuerURL: "https://issuer.example.com?foo=bar#baz",
			jwksURI:   exampleIssuer + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "issuer is not a valid URL",
			issuerURL: "issuer",
			jwksURI:   exampleIssuer + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "jwks missing https",
			issuerURL: exampleIssuer,
			jwksURI:   "http://issuer.example.com" + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "jwks missing scheme",
			issuerURL: exampleIssuer,
			jwksURI:   "issuer.example.com" + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:      "jwks is not a valid URL",
			issuerURL: exampleIssuer,
			jwksURI:   "issuer" + serviceaccount.JWKSPath,
			keys:      defaultKeys,
			err:       true,
		},
		{
			name:            "external address also has a scheme",
			issuerURL:       exampleIssuer,
			externalAddress: "https://192.0.2.1:80",
			keys:            defaultKeys,
			err:             true,
		},
		{
			name:      "missing external address and jwks",
			issuerURL: exampleIssuer,
			keys:      defaultKeys,
			err:       true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			keysGetter, err := serviceaccount.StaticPublicKeysGetter(tc.keys)
			if err != nil {
				t.Fatal(err)
			}
			md, err := serviceaccount.NewOpenIDMetadataProvider(tc.issuerURL, tc.jwksURI, tc.externalAddress, keysGetter)
			if tc.err {
				if err == nil {
					t.Fatalf("got <nil>, want error")
				}
				return
			} else if !tc.err && err != nil {
				t.Fatalf("got error %v, want <nil>", err)
			}

			config, _ := md.GetConfigJSON()
			keyset, _ := md.GetKeysetJSON()
			if string(config) != tc.wantConfig {
				t.Errorf("got metadata %s, want %s", string(config), tc.wantConfig)
			}
			if string(keyset) != tc.wantKeyset {
				t.Errorf("got keyset %s, want %s", string(keyset), tc.wantKeyset)
			}
		})
	}
}

type dummyPublicKeyGetter struct {
	keys []serviceaccount.PublicKey
}

func (d dummyPublicKeyGetter) AddListener(listener serviceaccount.Listener) {
	// no-op
}

func (d dummyPublicKeyGetter) GetCacheAgeMaxSeconds() int {
	return 3600
}

func (d dummyPublicKeyGetter) GetPublicKeys(ctx context.Context, keyIDHint string) []serviceaccount.PublicKey {
	return d.keys
}
