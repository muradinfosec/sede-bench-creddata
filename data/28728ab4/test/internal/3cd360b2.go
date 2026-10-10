package internal

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICWgIBAAKBgGFfgMY+DuO8l0RYrMLhcl6U/NigNIiOVhoo/xnYyoQALpWxBaBR
+umgwTnVZQVAnz65AOKj5sSuOZmN7GG2MSmBITzNI39OzUFlNUXPMb+3C2zC0K/u
TN8nioaz4XDR55FtdbtQo093tB3ZQiQRQPK+OwcJNg9hPFod+qasECCDMzZpryVl
BkxeaR2ctuIhczOR1mSaksscGCAt1baBZ/nCj+cqC92Oy0x9axwMYn0dWmff4ZuF
72dxUh3ExYAhKVskn5d9v5YuyWM1i1ET/Obetx9QBvVfCg/M3u2QCT46TKeE9uHv
zuMbxs4dCqSd48XxGogiUN1qpnoZIWoym2D1lrawdYf5pyJKNDOcTtHMbBK+oE1r
kWWUQihxGqSvxrscVq8tVVkw7CvDHRHA+nFa5JSNMJ4mc2ahWnRVTnWThUzEFa4z
lM/SLBnaReLEm2ysbwGD+WuBlyBsTlyqpC1vYcbVuaFr5nOQL8x7fhtlUwKITcb8
GRHPytaQwTCdVHkV/axcKNvW61BEhMBabMPL4Cs7dt3BHHa0kB/lPiWVvLsMM8FK
cWmsnE41IAJaZM8XZ/1fATWUpz9cGcvEjSInAVm7TOMGP0pbCegTZeEUqFxMggAd
/Iec4eObG7lvw6YSaAOvUTw3gi5MlDbG7wJmv390NRypB/znw2JHTpKhLBJR1wRn
Y6t+tC8V2Xb4P2NYLWDtasjFLIMDsdAMvJWFNrFAYukIr4DDS0s+H5E8o+Xj48Rt
CCHK729YuzkPu7nIkVqBCcCpwYELEGxQqsSb1Jz8
-----END RSA PRIVATE KEY-----`

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	KeyID     string `json:"kid"`
}

func TestPayload_buildToken(t *testing.T) {
	signer, err := getRSASigner(privateKey, "sampleKeyId")
	require.NoError(t, err)

	payload := Payload{IssuedAt: 1234, Expiry: 4321, Audience: "api.url", Issuer: "issuer", Subject: "subject"}

	token, err := payload.buildToken(&signer)
	require.NoError(t, err)

	segments := strings.Split(token, ".")
	require.Len(t, segments, 3)

	headerString, err := base64.RawStdEncoding.DecodeString(segments[0])
	require.NoError(t, err)

	var headerStruct Header
	err = json.Unmarshal(headerString, &headerStruct)
	require.NoError(t, err)

	payloadString, err := base64.RawStdEncoding.DecodeString(segments[1])
	require.NoError(t, err)

	var payloadStruct Payload
	err = json.Unmarshal(payloadString, &payloadStruct)
	require.NoError(t, err)

	expectedHeader := Header{Algorithm: "RS256", Type: "JWT", KeyID: "sampleKeyId"}

	assert.Equal(t, expectedHeader, headerStruct)
	assert.Equal(t, payload, payloadStruct)
}
