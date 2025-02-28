package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"errors"
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
	copy(b, p.key[len(p.key)-publicKeySize:])
	return &PublicKey{
		key: b,
	}
}

func NewPrivateKeyFromSeed(seed string) (*PrivateKey, error) {
	b, err := hex.DecodeString(seed)
	if err != nil {
		return nil, err
	}

	if len(b) != seedSize {
		return nil, errors.New("invalid seed size")
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(b),
	}, nil
}

func NewSeed() ([]byte, error) {
	b := make([]byte, seedSize)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GeneratePrivateKey() (*PrivateKey, error) {
	seed, err := NewSeed()
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

func (p *PublicKey) Address() *Address {
	return &Address{
		value: p.key[len(p.key)-addressSize:],
	}
}

func (p *PublicKey) Verify(msg, sig []byte) bool {
	return ed25519.Verify(p.key, msg, sig)
}

type Address struct {
	value []byte
}

func (a *Address) Bytes() []byte {
	return a.value
}

func (a *Address) String() string {
	return hex.EncodeToString(a.value)
}
