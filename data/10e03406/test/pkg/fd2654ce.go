package saml

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/crewjam/saml"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestHandleSamlLoginAddsRancherUserIdToCookie(t *testing.T) {
	host := "http://www.rancher.com/"
	mockUserId := "u-sidfes"
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	serviceProvider := &saml.ServiceProvider{
		Key:         privateKey,
		IDPMetadata: &saml.EntityDescriptor{},
	}
	cookieStore := &ClientCookies{
		Name:            "token",
		Domain:          host,
		Path:            "/",
		ServiceProvider: serviceProvider,
	}
	p := Provider{
		serviceProvider: serviceProvider,
		clientState:     cookieStore,
	}
	w := httptest.NewRecorder()

	urlParams, err := p.HandleSamlLogin(w, httptest.NewRequest(http.MethodPost, host, nil), mockUserId)

	assert.NoError(t, err)
	values, err := url.ParseQuery(urlParams)
	assert.NoError(t, err)
	relayState := values.Get("RelayState")
	// extract token from cookies. The key of the cookie is the relay state string, and the value is the token
	// e.g saml_DY3XlRnQsITgTsegiY-QuB38OsawU64uNZE4Q5iYCWIZOfz9YO6IYvUS=eyJhbGciOiJSTzg3NCIsInR5cCI6InpmaSJ9.eyJpZCI6Im54LTZTY2dPeTU5N1k1NzEyWjFFTUYzNTl2VzQydThMNmJCOHU1RHk5RG4iLCJUSHZ4WW9QVXNjSnZ1IjoicC1maERiU0siLCJObnkiOiJlYmFPOi8vT3paLkNXeHdLZ2suSFROLyJ9.ddICN7w1HmDp0ZwlncuCci8TXPGdknqdmNIjF54az3B; Path=/; Max-Age=90; HttpOnly
	cookie, err := http.ParseSetCookie(w.Header().Get("Set-Cookie"))
	assert.NoError(t, err)
	assert.Equal(t, "saml_"+relayState, cookie.Name)
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		secretBlock := x509.MarshalPKCS1PrivateKey(serviceProvider.Key)
		return secretBlock, nil
	})
	assert.NoError(t, err)
	claims, _ := token.Claims.(jwt.MapClaims)
	assert.Equal(t, mockUserId, claims[rancherUserID])
}
