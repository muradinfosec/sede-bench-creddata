package handshake

import (
	"crypto"
	"crypto/tls"

	"golang.org/x/crypto/hkdf"

	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/qtls"
)

var (
	quicSaltOld     = []byte{0xff, 0xff, 0xbf, 0x91, 0x15, 0x08, 0xf4, 0x4d, 0x8e, 0x88, 0x14, 0xb6, 0x3e, 0x91, 0x56, 0xf9, 0x46, 0x78, 0xc3, 0x00}
	quicSaltDraft34 = []byte{0x34, 0x48, 0x7c, 0xf8, 0xf4, 0x52, 0x14, 0xc5, 0x1c, 0x22, 0x4f, 0xf7, 0xd8, 0xe7, 0x8c, 0xcc, 0xfe, 0xfd, 0x5e, 0x4c}
)

func getSalt(v protocol.VersionNumber) []byte {
	if v == protocol.VersionDraft34 || v == protocol.Version1 {
		return quicSaltDraft34
	}
	return quicSaltOld
}

var initialSuite = &qtls.CipherSuiteTLS13{
	ID:     tls.TLS_AES_128_GCM_SHA256,
	KeyLen: 16,
	AEAD:   qtls.AEADAESGCMTLS13,
	Hash:   crypto.SHA256,
}

// NewInitialAEAD creates a new AEAD for Initial encryption / decryption.
func NewInitialAEAD(connID protocol.ConnectionID, pers protocol.Perspective, v protocol.VersionNumber) (LongHeaderSealer, LongHeaderOpener) {
	clientSecret, serverSecret := computeSecrets(connID, v)
	var mySecret, otherSecret []byte
	if pers == protocol.PerspectiveClient {
		mySecret = clientSecret
		otherSecret = serverSecret
	} else {
		mySecret = serverSecret
		otherSecret = clientSecret
	}
	myKey, myIV := computeInitialKeyAndIV(mySecret)
	otherKey, otherIV := computeInitialKeyAndIV(otherSecret)

	encrypter := qtls.AEADAESGCMTLS13(myKey, myIV)
	decrypter := qtls.AEADAESGCMTLS13(otherKey, otherIV)

	return newLongHeaderSealer(encrypter, newHeaderProtector(initialSuite, mySecret, true)),
		newLongHeaderOpener(decrypter, newAESHeaderProtector(initialSuite, otherSecret, true))
}

func computeSecrets(connID protocol.ConnectionID, v protocol.VersionNumber) (clientSecret, serverSecret []byte) {
	initialSecret := hkdf.Extract(crypto.SHA256.New, connID, getSalt(v))
	clientSecret = hkdfExpandLabel(crypto.SHA256, initialSecret, []byte{}, "client in", crypto.SHA256.Size())
	serverSecret = hkdfExpandLabel(crypto.SHA256, initialSecret, []byte{}, "server in", crypto.SHA256.Size())
	return
}

func computeInitialKeyAndIV(secret []byte) (key, iv []byte) {
	key = hkdfExpandLabel(crypto.SHA256, secret, []byte{}, "quic key", 16)
	iv = hkdfExpandLabel(crypto.SHA256, secret, []byte{}, "quic iv", 12)
	return
}
