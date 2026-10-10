package security

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/codec"
	"github.com/tal-tech/go-zero/core/fs"
	"github.com/tal-tech/go-zero/rest/httpx"
)

const (
	pubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCyeDYV2ieOtNDi6tuNtAbmUjN9
pTHluAU5yiKEz8826QohcxqUKP3hybZBcm60p+rUxMAJFBJ8Dt+UJ6sEMzrf1rOF
YOImVvORkXjpFU7sCJkhnLMs/kxtRzcZJG6ADUlG4GDCNcZpY/qELEvwgm2kCcHi
tGC2mO8opFFFHTR0aQIDAQAB
-----END PUBLIC KEY-----`
	priKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCyeDYV2ieOtNDi6tuNtAbmUjN9pTHluAU5yiKEz8826QohcxqU
zM0OulotnN40w+qqAbqbLYc7Ve+EQ1vsQhMJ8IZjYIxZZZwoNYnxXV8WlmLJNZAl
/FicpItOTu4ImCoQ1pyAETLSM/RATxAqhd5tXqYFQOr8Ph8yRHIRVWp4HuZTDfHf
IxojRIkk+fa2SIIOE5zHfPBL35goRQbNE0+TaHmB3Jm49XHjQG4o14GTMTx58ast
XBVHNmRGUYJQL1LWkfmszw8D3Q19rbW/6ULmSW6nfk3O1S8+4z1gNnI36PFS5z+h
dqDYHYDl/4wgA5MPzq14fSCx/+LmKZxAynOyM7UyLdh5fuGlXuqq2vy7KjKYGOgg
p+ngD/B3fBjKGgslxFJTGPPd7f9OIxuQzNxBxo5bd9HyZHN6EdweZMn2r4/2SYC9
8ILlwXd0oHRyrVhzHOeYiKnIzzPAPQMEp7/1aLyYe+xLtBXxfalMZwp7jfRRC6B+
WgKHK7fQgAODaV9ZeA2PHDyLp38IGxE98QuSfwTn0/uGuALFoTQ9oKr0djU6kV18
MLoaLRyBym4GPQ+6RKewLKbTK5m9eIABsXf7zpW/NUq9OJsfjo0QWkUlclrzaAlG
WTCBdtZMPBOIHv8uVP3ShCFoDnTLx3XDqDXRWrv9zQcx87/cL+sVp5W9AseKdV24
3BVp+V9iYtvS8VPGzO9Bkpqe1DdfikguJKS+HXMZqLJWTJZUpCHa3xsPjGuhCEfb
WKLH5J+NyPOPbG5CIyTNhxPUowCNthM2vA+YPdd7hqTX
-----END RSA PRIVATE KEY-----`
	body = "hello world!"
)

var key = []byte("r7i2i!s%C*U-TjKuGeStAv4j8i5g/C?P")

func TestContentSecurity(t *testing.T) {
	tests := []struct {
		name        string
		mode        string
		extraKey    string
		extraSecret string
		extraTime   string
		err         error
		code        int
	}{
		{
			name: "encrypted",
			mode: "1",
		},
		{
			name: "unencrypted",
			mode: "0",
		},
		{
			name: "bad content type",
			mode: "a",
			err:  ErrInvalidContentType,
		},
		{
			name:        "bad secret",
			mode:        "1",
			extraSecret: "any",
			err:         ErrInvalidSecret,
		},
		{
			name:     "bad key",
			mode:     "1",
			extraKey: "any",
			err:      ErrInvalidKey,
		},
		{
			name:      "bad time",
			mode:      "1",
			extraTime: "any",
			code:      httpx.CodeSignatureInvalidHeader,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			r, err := http.NewRequest(http.MethodPost, "http://localhost:3333/a/b?c=first&d=second",
				strings.NewReader(body))
			assert.Nil(t, err)

			timestamp := time.Now().Unix()
			sha := sha256.New()
			sha.Write([]byte(body))
			bodySign := fmt.Sprintf("%x", sha.Sum(nil))
			contentOfSign := strings.Join([]string{
				strconv.FormatInt(timestamp, 10),
				http.MethodPost,
				r.URL.Path,
				r.URL.RawQuery,
				bodySign,
			}, "\n")
			sign := hs256(key, contentOfSign)
			content := strings.Join([]string{
				"version=v1",
				"type=" + test.mode,
				fmt.Sprintf("key=%s", base64.StdEncoding.EncodeToString(key)) + test.extraKey,
				"time=" + strconv.FormatInt(timestamp, 10) + test.extraTime,
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
			r.Header.Set("X-Content-Security", strings.Join([]string{
				fmt.Sprintf("key=%s", fingerprint(pubKey)),
				"secret=" + encryptedContent + test.extraSecret,
				"signature=" + sign,
			}, "; "))

			file, err := fs.TempFilenameWithText(priKey)
			assert.Nil(t, err)
			defer os.Remove(file)

			dec, err := codec.NewRsaDecrypter(file)
			assert.Nil(t, err)

			header, err := ParseContentSecurity(map[string]codec.RsaDecrypter{
				fingerprint(pubKey): dec,
			}, r)
			assert.Equal(t, test.err, err)
			if err != nil {
				return
			}

			assert.Equal(t, test.code, VerifySignature(r, header, time.Minute))
		})
	}
}

func fingerprint(key string) string {
	h := md5.New()
	io.WriteString(h, key)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func hs256(key []byte, body string) string {
	h := hmac.New(sha256.New, key)
	io.WriteString(h, body)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
