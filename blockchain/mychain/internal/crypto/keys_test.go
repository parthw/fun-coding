package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey, err := GeneratePrivateKey()
	assert.NoError(t, err)
	assert.NotNil(t, privKey)
	assert.Equal(t, privateKeySize, len(privKey.Bytes()))

	pubKey := privKey.PublicKey()
	assert.NotNil(t, pubKey)
	assert.Equal(t, publicKeySize, len(pubKey.Bytes()))
}

func TestSignAndVerify(t *testing.T) {
	privKey, err := GeneratePrivateKey()
	assert.NoError(t, err)
	assert.NotNil(t, privKey)

	pubKey := privKey.PublicKey()
	assert.NotNil(t, pubKey)

	msg := []byte("hello world")
	sig := privKey.Sign(msg)
	assert.NotNil(t, sig)
	assert.True(t, pubKey.Verify(msg, sig))
}
