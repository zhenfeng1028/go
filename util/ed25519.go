package main

import (
	"crypto/ed25519"
)

func main() {
	_, privateKey, _ := ed25519.GenerateKey(nil)
	publicKey := privateKey.Public().(ed25519.PublicKey)
	message := []byte("The quick brown fox jumps over the lazy dog")
	sig := ed25519.Sign(privateKey, message)
	if !ed25519.Verify(publicKey, message, sig) {
		panic("signature not valid")
	}
}
