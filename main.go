package main

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generateKeyPair() (ad string, pk string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.String(), hexutil.Encode(privateKeyBytes)
}

func main() {
	var prfix = "abc"
	var tail = "1"
	var thread_num = 10

	for i := 0; i < thread_num; i++ {
		go func() {
			for {
				address, privateKey := generateKeyPair()
				if address[2:2+len(prfix)] == prfix && address[len(address)-len(tail):] == tail {
					println("Address: ", address)
					println("Private key: ", privateKey)
				}
			}
		}()
	}
	for {
	}
}
