package main

import (
	"fmt"
	"io"
	"math/rand"
	"runtime"
	"time"

	"golang.org/x/crypto/curve25519"
)

const (
	KeySize = 32
)

var (
	version   = "v0.0.0"
	gitCommit string
	buildDate = "1970-01-01T00:00:00Z"
	goVersion = runtime.Version()
)

func main() {
	fmt.Printf("Version: %s\nBuild date: %s\nGit commit: %s\nGo version: %s\n", version, buildDate, gitCommit, goVersion)
	rand.Seed(time.Now().UnixNano())
	for {
		// publicKey, privateKey, err := GenerateKeys(rand.Reader)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Printf("%x %x\n", publicKey, privateKey)
		// time.Sleep(1 * time.Second)
		fmt.Printf(
			"W %s [BlablaProcessor.onBlablaRequest-C-0]> NetworkClient: [Consumer clientId=consumer-%d, groupId=lol-service] Connection to node 0 could not be established. Broker may not be available. [lol-service,,,,,,,,,]\n",
			// "%s WARN Client session timed out, have not heard from server in %dms for sessionid 0x%x (org.apache.zookeeper.ClientCnxn)\n",
			time.Now().Format("2006-01-02 15:04:05.000"),
			rand.Intn(1000),
			// rand.Uint64(),
		)
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
