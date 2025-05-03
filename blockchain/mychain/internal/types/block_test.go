package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"parthw.in/blockchain/mychain/internal/crypto"
	"parthw.in/blockchain/mychain/internal/utils"
)

func TestHashBlock(t *testing.T) {
	block := utils.RandomBlock()
	hash, err := HashBlock(block)
	assert.NoError(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, 32, len(hash))
}

func TestSignBlock(t *testing.T) {
	block := utils.RandomBlock()
	hash, err := HashBlock(block)
	assert.NoError(t, err)
	assert.NotNil(t, hash)

	priv, err := crypto.GeneratePrivateKey()
	assert.NoError(t, err)
	assert.NotNil(t, priv)

	signature, err := SignBlock(priv, block)
	assert.NoError(t, err)
	assert.NotNil(t, signature)

	assert.Equal(t, 64, len(signature))
	pub := priv.PublicKey()
	assert.True(t, pub.Verify(hash, signature))
}
