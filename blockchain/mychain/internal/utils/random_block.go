package utils

import (
	"crypto/rand"
	"io"
	"time"

	mrand "math/rand"

	"parthw.in/blockchain/mychain/proto"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(rand.Reader, hash)
	return hash
}

func RandomBlock() *proto.Block {
	headers := &proto.Header{
		Version:      1,
		Height:       int32(mrand.Intn(1000)),
		Timestamp:    time.Now().UnixNano(),
		PreviousHash: RandomHash(),
		RootHash:     RandomHash(),
	}
	return &proto.Block{
		Header: headers,
	}
}
