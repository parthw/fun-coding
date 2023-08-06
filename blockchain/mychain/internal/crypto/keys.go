package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

func (p *PrivateKey) PublicKey() *PublicKey {
	b := make([]byte, publicKeySize)
	copy(b, p.key[32:])
	return &PublicKey{
		key: b,
	}
}

func GeneratePrivateKey() (*PrivateKey, error) {
	seed := make([]byte, seedSize)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}, nil
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func (p *PublicKey) Verify(msg, sig []byte) bool {
	return ed25519.Verify(p.key, msg, sig)
}
