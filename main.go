package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
	"golang.org/x/crypto/curve25519"
)

const (
	KeySize = 32
)

func main() {
	for {
		publicKey, privateKey, err := GenerateKeys(rand.Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%x %x\n", publicKey, privateKey)
		time.Sleep(1 * time.Second)
	}
}

// GenerateKeys creates a public/private key pair
func GenerateKeys(rand io.Reader) ([]byte, []byte, error) {
	publicKey := new([KeySize]byte)
	privateKey := new([KeySize]byte)
	if _, err := io.ReadFull(rand, privateKey[:]); err != nil {
		return nil, nil, err
	}

	curve25519.ScalarBaseMult(publicKey, privateKey)
	return publicKey[:], privateKey[:], nil
}
