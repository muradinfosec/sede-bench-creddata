package cookie

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeAndDecodeAccessToken(t *testing.T) {
	const secret = "1810151434lplmqpqkjeoflqysfgynlu"
	const token = "my access token"
	c, err := NewCipher([]byte(secret))
	assert.Equal(t, nil, err)

	encoded, err := c.Encrypt(token)
	assert.Equal(t, nil, err)

	decoded, err := c.Decrypt(encoded)
	assert.Equal(t, nil, err)

	assert.NotEqual(t, token, encoded)
	assert.Equal(t, token, decoded)
}

func TestEncodeAndDecodeAccessTokenB64(t *testing.T) {
	const secret_b64 = "O5Obl2tx5Ex5AjypG4azrg-yArvltsZPJ-DyMhy2PKs="
	const token = "my access token"

	secret, err := base64.URLEncoding.DecodeString(secret_b64)
	c, err := NewCipher([]byte(secret))
	assert.Equal(t, nil, err)

	encoded, err := c.Encrypt(token)
	assert.Equal(t, nil, err)

	decoded, err := c.Decrypt(encoded)
	assert.Equal(t, nil, err)

	assert.NotEqual(t, token, encoded)
	assert.Equal(t, token, decoded)
}
