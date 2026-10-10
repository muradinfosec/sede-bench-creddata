package handler

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandlerFailed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	handler := Authorize("C15A185F-EEA1-0E22-99C9-D3144C39380B", WithUnauthorizedCallback(
		func(w http.ResponseWriter, r *http.Request, err error) {
			w.Header().Set("X-Test", "test")
			w.WriteHeader(http.StatusUnauthorized)
			_, err = w.Write([]byte("content"))
			assert.Nil(t, err)
		}))(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestAuthHandler(t *testing.T) {
	const key = "A26D554F-BFE6-9F69-02B8-E6814F86922D"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	token, err := buildToken(key, map[string]interface{}{
		"key": "value",
	}, 3600)
	assert.Nil(t, err)
	req.Header.Set("Authorization", "Bearer "+token)
	handler := Authorize(key)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Test", "test")
			_, err := w.Write([]byte("content"))
			assert.Nil(t, err)

			flusher, ok := w.(http.Flusher)
			assert.True(t, ok)
			flusher.Flush()
		}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "test", resp.Header().Get("X-Test"))
	assert.Equal(t, "content", resp.Body.String())
}

func TestAuthHandlerWithPrevSecret(t *testing.T) {
	const (
		key     = "22C98870-CB9A-516A-8E63-5102564DCC52"
		prevKey = "D72C295A-DEC2-7D04-03B5-A3825F11678F"
	)
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	token, err := buildToken(key, map[string]interface{}{
		"key": "value",
	}, 3600)
	assert.Nil(t, err)
	req.Header.Set("Authorization", "Bearer "+token)
	handler := Authorize(key, WithPrevSecret(prevKey))(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Test", "test")
			_, err := w.Write([]byte("content"))
			assert.Nil(t, err)
		}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "test", resp.Header().Get("X-Test"))
	assert.Equal(t, "content", resp.Body.String())
}

func TestAuthHandler_NilError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	resp := httptest.NewRecorder()
	assert.NotPanics(t, func() {
		unauthorized(resp, req, nil, nil)
	})
}

func TestAuthHandler_Flush(t *testing.T) {
	resp := httptest.NewRecorder()
	handler := newGuardedResponseWriter(resp)
	handler.Flush()
	assert.True(t, resp.Flushed)
}

func TestAuthHandler_Hijack(t *testing.T) {
	resp := httptest.NewRecorder()
	writer := newGuardedResponseWriter(resp)
	assert.NotPanics(t, func() {
		writer.Hijack()
	})

	writer = newGuardedResponseWriter(mockedHijackable{resp})
	assert.NotPanics(t, func() {
		writer.Hijack()
	})
}

func buildToken(secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

type mockedHijackable struct {
	*httptest.ResponseRecorder
}

func (m mockedHijackable) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}
