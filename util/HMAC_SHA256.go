package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// create a random 64 bytes (512 bits) secret
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	if err != nil {
		fmt.Println("error generating a random secret:", err)
		return
	}

	data := []byte("Hello World")

	// create a new HMAC by defining the hash type and the key
	hmac := hmac.New(sha256.New, secret)

	// compute the HMAC
	hmac.Write([]byte(data))
	dataHmac := hmac.Sum(nil)

	hmacHex := hex.EncodeToString(dataHmac)
	secretHex := hex.EncodeToString(secret)

	fmt.Printf("HMAC_SHA256(key: %s, data: %s): %s", secretHex, string(data), hmacHex)
}
