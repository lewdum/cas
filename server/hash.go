package main

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/blake2b"
)

// TODO: support other hash functions
// TODO: support other hash sizes

func hashFile(data []byte) string {
	hasher, err := blake2b.New256([]byte(*flagSecret))
	if err != nil {
		panic(fmt.Errorf("failed to create hasher: %w", err))
	}
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
