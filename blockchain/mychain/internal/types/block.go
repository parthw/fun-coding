package types

import (
	"crypto/sha256"

	"parthw.in/blockchain/mychain/internal/crypto"
	"parthw.in/blockchain/mychain/proto"

	pb "google.golang.org/protobuf/proto"
)

func HashBlock(b *proto.Block) ([]byte, error) {
	bBytes, err := pb.Marshal(b)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(bBytes)
	return hash[:], nil
}

func SignBlock(p *crypto.PrivateKey, b *proto.Block) ([]byte, error) {
	hash, err := HashBlock(b)
	if err != nil {
		return nil, err
	}

	return p.Sign(hash), nil
}
