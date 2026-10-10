package token

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/timex"
)

func TestTokenParser(t *testing.T) {
	const (
		key     = "14A98298-DD2A-714B-1C61-9267530AAE90"
		prevKey = "C17B302E-AAC1-7E37-84F7-F4984E49194D"
	)
	keys := []struct {
		key     string
		prevKey string
	}{
		{
			key,
			prevKey,
		},
		{
			key,
			"",
		},
	}

	for _, pair := range keys {
		req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		token, err := buildToken(key, map[string]interface{}{
			"key": "value",
		}, 3600)
		assert.Nil(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		parser := NewTokenParser(WithResetDuration(time.Minute))
		tok, err := parser.ParseToken(req, pair.key, pair.prevKey)
		assert.Nil(t, err)
		assert.Equal(t, "value", tok.Claims.(jwt.MapClaims)["key"])
	}
}

func TestTokenParser_Expired(t *testing.T) {
	const (
		key     = "60F96246-CB0E-596A-2B10-7009585FCA41"
		prevKey = "C44C533C-EDB7-4B13-94D7-C4072F15016C"
	)
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	token, err := buildToken(key, map[string]interface{}{
		"key": "value",
	}, 3600)
	assert.Nil(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	parser := NewTokenParser(WithResetDuration(time.Second))
	tok, err := parser.ParseToken(req, key, prevKey)
	assert.Nil(t, err)
	assert.Equal(t, "value", tok.Claims.(jwt.MapClaims)["key"])
	tok, err = parser.ParseToken(req, key, prevKey)
	assert.Nil(t, err)
	assert.Equal(t, "value", tok.Claims.(jwt.MapClaims)["key"])
	parser.resetTime = timex.Now() - time.Hour
	tok, err = parser.ParseToken(req, key, prevKey)
	assert.Nil(t, err)
	assert.Equal(t, "value", tok.Claims.(jwt.MapClaims)["key"])
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
