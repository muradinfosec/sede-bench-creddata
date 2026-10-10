package handler

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/codec"
	"github.com/tal-tech/go-zero/rest/httpx"
)

const timeDiff = time.Hour * 2 * 24

var (
	fingerprint = "12345"
	pubKey      = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD7bq4FLG0ctccbEFEsUBuRxkjE
eJ5U+0CAEjJk20V9/u2Fu76i1oKoShCs7GXtAFbDb5A/ImIXkPY62nAaxTGK4KVH
miYbRgh5Fy6336KepLCtCmV/r0PKZeCyJH9uYLs7EuE1z9Hgm5UUjmpHDhJtkAwR
my47YlhspwszKdRP+wIDAQAB
-----END PUBLIC KEY-----`)
	priKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQD7bq4FLG0ctccbEFEsUBuRxkjEeJ5U+0CAEjJk20V9/u2Fu76i
9QkaNVae8PsPqKFEI0G/WydNTjM63ecXRjYo8sCUadYxVXq8TY5246yrcDStOgF/
x8gEPsVThe5sWuR8eTm6D7xgy2CAsuPauCgEIUOknp60OiKzIsGlFFMy+qTsWDLM
IAgRTwt6dR9eFbbXH2kKLB7X7L1R48zKLAzrPz+fTE9x6qhREhDZmh8hcyDuHvrE
aU8ft0To2OPR56sjP/0hFB0NbFoZzTv06ildfFO5ssD/24gl6zPUqEwH3Ws2VAHs
S2fcIqezgah/RHZb4H6xzts31t08O7usxhysEVjfAcwXiEyIXOYo/mDLew+73kb3
os1ibf6Rmuy9YhBNWUDvXVE1oZ8MvEZx2YmtPtZBEsABwgnGRU7CnlPhPKNaXpzT
ijuiRAr6ESMMgXz9EVLAPoIVymOBVsIBbtNCOzTu7CivKp0EVcFk1e9sCotojziI
uVGm0KonYkbelD9hzByV6HXHPMlaMl8VJLrUExtlhoTtOz3KCRFeybXL5rXIZndh
2SugbyiBWa5bQSplGzTil7Ly+UZS5PB2SRUT7psO9pgW2ly3bew8h8qf23bddElY
XpnB0NIhlLInsJkTjpUTkI8CPbUxkmS+11sV3LfYAL2CvK1is5mewIxfmlNWac5a
R/YvPH2X/nvvqv6FxtWrvW670RWJzXrB2+R3Vko4Vqvfd9sDjJGuUgAoLHNPu1nP
KzIgxFfFE2INfJifqigqm59Oj4tskMrHl4JpWbs0Kno=
-----END RSA PRIVATE KEY-----`)
	key = []byte("z2a5l!g%C*X-WnEwTaEdIl0s9r6f/H?W")
)

type requestSettings struct {
	method      string
	url         string
	body        io.Reader
	strict      bool
	crypt       bool
	requestUri  string
	timestamp   int64
	fingerprint string
	missHeader  bool
	signature   string
}

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestContentSecurityHandler(t *testing.T) {
	tests := []struct {
		method      string
		url         string
		body        string
		strict      bool
		crypt       bool
		requestUri  string
		timestamp   int64
		fingerprint string
		missHeader  bool
		signature   string
		statusCode  int
	}{
		{
			method: http.MethodGet,
			url:    "http://localhost/a/b?c=d&e=f",
			strict: true,
			crypt:  false,
		},
		{
			method: http.MethodPost,
			url:    "http://localhost/a/b?c=d&e=f",
			body:   "hello",
			strict: true,
			crypt:  false,
		},
		{
			method: http.MethodGet,
			url:    "http://localhost/a/b?c=d&e=f",
			strict: true,
			crypt:  true,
		},
		{
			method: http.MethodPost,
			url:    "http://localhost/a/b?c=d&e=f",
			body:   "hello",
			strict: true,
			crypt:  true,
		},
		{
			method:     http.MethodGet,
			url:        "http://localhost/a/b?c=d&e=f",
			strict:     true,
			crypt:      true,
			timestamp:  time.Now().Add(timeDiff).Unix(),
			statusCode: http.StatusForbidden,
		},
		{
			method:     http.MethodPost,
			url:        "http://localhost/a/b?c=d&e=f",
			body:       "hello",
			strict:     true,
			crypt:      true,
			timestamp:  time.Now().Add(-timeDiff).Unix(),
			statusCode: http.StatusForbidden,
		},
		{
			method:     http.MethodPost,
			url:        "http://remotehost/",
			body:       "hello",
			strict:     true,
			crypt:      true,
			requestUri: "http://localhost/a/b?c=d&e=f",
		},
		{
			method:      http.MethodPost,
			url:         "http://localhost/a/b?c=d&e=f",
			body:        "hello",
			strict:      false,
			crypt:       true,
			fingerprint: "badone",
		},
		{
			method:      http.MethodPost,
			url:         "http://localhost/a/b?c=d&e=f",
			body:        "hello",
			strict:      true,
			crypt:       true,
			timestamp:   time.Now().Add(-timeDiff).Unix(),
			fingerprint: "badone",
			statusCode:  http.StatusForbidden,
		},
		{
			method:     http.MethodPost,
			url:        "http://localhost/a/b?c=d&e=f",
			body:       "hello",
			strict:     true,
			crypt:      true,
			missHeader: true,
			statusCode: http.StatusForbidden,
		},
		{
			method: http.MethodHead,
			url:    "http://localhost/a/b?c=d&e=f",
			strict: true,
			crypt:  false,
		},
		{
			method:     http.MethodGet,
			url:        "http://localhost/a/b?c=d&e=f",
			strict:     true,
			crypt:      false,
			signature:  "badone",
			statusCode: http.StatusForbidden,
		},
	}

	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			if test.statusCode == 0 {
				test.statusCode = http.StatusOK
			}
			if len(test.fingerprint) == 0 {
				test.fingerprint = fingerprint
			}
			if test.timestamp == 0 {
				test.timestamp = time.Now().Unix()
			}

			func() {
				keyFile, err := createTempFile(priKey)
				defer os.Remove(keyFile)

				assert.Nil(t, err)
				decrypter, err := codec.NewRsaDecrypter(keyFile)
				assert.Nil(t, err)
				contentSecurityHandler := ContentSecurityHandler(map[string]codec.RsaDecrypter{
					fingerprint: decrypter,
				}, time.Hour, test.strict)
				handler := contentSecurityHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				}))

				var reader io.Reader
				if len(test.body) > 0 {
					reader = strings.NewReader(test.body)
				}
				setting := requestSettings{
					method:      test.method,
					url:         test.url,
					body:        reader,
					strict:      test.strict,
					crypt:       test.crypt,
					requestUri:  test.requestUri,
					timestamp:   test.timestamp,
					fingerprint: test.fingerprint,
					missHeader:  test.missHeader,
					signature:   test.signature,
				}
				req, err := buildRequest(setting)
				assert.Nil(t, err)
				resp := httptest.NewRecorder()
				handler.ServeHTTP(resp, req)
				assert.Equal(t, test.statusCode, resp.Code)
			}()
		})
	}
}

func TestContentSecurityHandler_UnsignedCallback(t *testing.T) {
	keyFile, err := createTempFile(priKey)
	defer os.Remove(keyFile)

	assert.Nil(t, err)
	decrypter, err := codec.NewRsaDecrypter(keyFile)
	assert.Nil(t, err)
	contentSecurityHandler := ContentSecurityHandler(
		map[string]codec.RsaDecrypter{
			fingerprint: decrypter,
		},
		time.Hour,
		true,
		func(w http.ResponseWriter, r *http.Request, next http.Handler, strict bool, code int) {
			w.WriteHeader(http.StatusOK)
		})
	handler := contentSecurityHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	setting := requestSettings{
		method:    http.MethodGet,
		url:       "http://localhost/a/b?c=d&e=f",
		signature: "badone",
	}
	req, err := buildRequest(setting)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestContentSecurityHandler_UnsignedCallback_WrongTime(t *testing.T) {
	keyFile, err := createTempFile(priKey)
	defer os.Remove(keyFile)

	assert.Nil(t, err)
	decrypter, err := codec.NewRsaDecrypter(keyFile)
	assert.Nil(t, err)
	contentSecurityHandler := ContentSecurityHandler(
		map[string]codec.RsaDecrypter{
			fingerprint: decrypter,
		},
		time.Hour,
		true,
		func(w http.ResponseWriter, r *http.Request, next http.Handler, strict bool, code int) {
			assert.Equal(t, httpx.CodeSignatureWrongTime, code)
			w.WriteHeader(http.StatusOK)
		})
	handler := contentSecurityHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	var reader io.Reader
	reader = strings.NewReader("hello")
	setting := requestSettings{
		method:      http.MethodPost,
		url:         "http://localhost/a/b?c=d&e=f",
		body:        reader,
		strict:      true,
		crypt:       true,
		timestamp:   time.Now().Add(time.Hour * 24 * 365).Unix(),
		fingerprint: fingerprint,
	}
	req, err := buildRequest(setting)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func buildRequest(rs requestSettings) (*http.Request, error) {
	var bodyStr string
	var err error

	if rs.crypt && rs.body != nil {
		var buf bytes.Buffer
		io.Copy(&buf, rs.body)
		bodyBytes, err := codec.EcbEncrypt(key, buf.Bytes())
		if err != nil {
			return nil, err
		}
		bodyStr = base64.StdEncoding.EncodeToString(bodyBytes)
	}

	r := httptest.NewRequest(rs.method, rs.url, strings.NewReader(bodyStr))
	if len(rs.signature) == 0 {
		sha := sha256.New()
		sha.Write([]byte(bodyStr))
		bodySign := fmt.Sprintf("%x", sha.Sum(nil))
		var path string
		var query string
		if len(rs.requestUri) > 0 {
			u, err := url.Parse(rs.requestUri)
			if err != nil {
				return nil, err
			}

			path = u.Path
			query = u.RawQuery
		} else {
			path = r.URL.Path
			query = r.URL.RawQuery
		}
		contentOfSign := strings.Join([]string{
			strconv.FormatInt(rs.timestamp, 10),
			rs.method,
			path,
			query,
			bodySign,
		}, "\n")
		rs.signature = codec.HmacBase64([]byte(key), contentOfSign)
	}

	var mode string
	if rs.crypt {
		mode = "1"
	} else {
		mode = "0"
	}
	content := strings.Join([]string{
		"version=v1",
		"type=" + mode,
		fmt.Sprintf("key=%s", base64.StdEncoding.EncodeToString(key)),
		"time=" + strconv.FormatInt(rs.timestamp, 10),
	}, "; ")

	encrypter, err := codec.NewRsaEncrypter([]byte(pubKey))
	if err != nil {
		log.Fatal(err)
	}

	output, err := encrypter.Encrypt([]byte(content))
	if err != nil {
		log.Fatal(err)
	}

	encryptedContent := base64.StdEncoding.EncodeToString(output)
	if !rs.missHeader {
		r.Header.Set(httpx.ContentSecurity, strings.Join([]string{
			fmt.Sprintf("key=%s", rs.fingerprint),
			"secret=" + encryptedContent,
			"signature=" + rs.signature,
		}, "; "))
	}
	if len(rs.requestUri) > 0 {
		r.Header.Set("X-Request-Uri", rs.requestUri)
	}

	return r, nil
}

func createTempFile(body []byte) (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "go-unit-*.tmp")
	if err != nil {
		return "", err
	}

	tmpFile.Close()
	err = ioutil.WriteFile(tmpFile.Name(), body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}
