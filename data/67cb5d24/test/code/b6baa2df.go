package codec

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/fs"
)

const (
	priKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC4TJk3onpqb2RYE3wwt23J9SHLFstHGSkUYFLe+nl1dEKHbD+/
Nc39D717f6OaJYIXMw9jXhtkufI+HSD3n29xCWH/ZmL5In7Yvx/d4tN83ezkc9YV
aXmOiAdg4dtc69BYWGbs+jONXY6G9BzYYOELj42xQ3p0BGwc217vuYmzviZpoZWd
xrUPm5NCmIXnEq5u1JyE81/3JNrXpglhUmp6Wzakq8FjjCBB9hiJl4Y3/Ayq930q
YnQPIePFrEuX5luwZCGpRUHZr1Xq1lrkwd+2zUiK4RyxY4Q137iu7FiJsGrt/P5u
aQ1Q3T/IuUr7iMp10kjtIZOmoQvqDY/Vpgm+qJyDJkJyAfWyKxA3PS2qt/6tcq9O
s2B3rK8jnA16E6cxLwox4SW/RkczQ2hqzuGBZ4YM//Oz331285oDEoXR7/t72miF
RdCmZio3mOwMgQs0hES/kaZIY8aGBIXaqFfAf/FFM4zhYBFZl2WN8fSrJ9Hpqqb8
WudRfHoKwAxZgftwWwZ9L+vpsHBxej3cwojQrrn8aOGI3O8VcLDjFHUQ37i3CYxS
qT87HvIYSQx1aHG7tO6XdEFxg3yfIGgfFHJ8QaOK8xN5mLUL/oaiR1rcsKADPVSs
nFzMiqWVQaMXsseiNugrMeRomJRXnjoRCrAv+XzrrQOq7qPpGumnJ8Prn8yZU8BE
BWdmnqdQRkw37hX+nnJZGriEo76wKOEyIslD7dupXtvBkh958o8d93w7RFrOUdHa
IymG2UqarxrkjkRs1eVXm8ZZj90vz8LgIeF2Pj8m4p74
-----END RSA PRIVATE KEY-----`
	pubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC4TJk3onpqb2RYE3wwt23J9SHL
FstHGSkUYFLe+nl1dEKHbD+/Zt95L757J3xGTrwoTc7KCTxbrgn+stn0w52BNjj/
kIE2ko4lbh/v8Fl14AyVR9msfKtKOnhe5FCT72mdtApr+qvzcC3q9hfXwkyQU32p
v7q5UimZ205iKSBmgQIDAQAB
-----END PUBLIC KEY-----`
	testBody = `this is the content`
)

func TestCryption(t *testing.T) {
	enc, err := NewRsaEncrypter([]byte(pubKey))
	assert.Nil(t, err)
	ret, err := enc.Encrypt([]byte(testBody))
	assert.Nil(t, err)

	file, err := fs.TempFilenameWithText(priKey)
	assert.Nil(t, err)
	dec, err := NewRsaDecrypter(file)
	assert.Nil(t, err)
	actual, err := dec.Decrypt(ret)
	assert.Nil(t, err)
	assert.Equal(t, testBody, string(actual))

	actual, err = dec.DecryptBase64(base64.StdEncoding.EncodeToString(ret))
	assert.Nil(t, err)
	assert.Equal(t, testBody, string(actual))
}

func TestBadPubKey(t *testing.T) {
	_, err := NewRsaEncrypter([]byte("foo"))
	assert.Equal(t, ErrPublicKey, err)
}
