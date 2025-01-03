package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	passphrase, err := GeneratePrivatePassphrase(15)
	if err != nil {
		log.Fatalf("Error generating passphrase: %v", err)
	}

	// Print the passphrase
	fmt.Printf("Generated Passphrase: %s\n", passphrase)

	TestVerifySignature()

	// Create a 32-byte slice (256 bits)
	randomBytes := make([]byte, 32)

	// Fill the slice with random bytes
	_, err = rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("Failed to generate random bytes: %v", err)
	}

	// Convert to hexadecimal format
	randomHex := hex.EncodeToString(randomBytes)

	// Print the 256-bit random number
	fmt.Printf("256-bit random number (hex): %s\n", randomHex)

	fmt.Println(CreateNewWallet().address)

}
