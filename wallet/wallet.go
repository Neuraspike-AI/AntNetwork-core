package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	base58 "github.com/btcsuite/btcutil/base58"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    string
}

const Base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func GeneratePrivatePassphrase(words int) (string, error) {
	filePath := "/home/oli/Projects/AntNetwork-core/wallet/5000-more-common.txt"
	var passphrase []string
	wordList, err := loadWords(filePath)
	if err != nil {
		log.Fatalf("Error loading word list: %v", err)
	}
	for i := 0; i < words; i++ {
		// Select a random index from the word list
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordList))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %v", err)
		}

		// Add the word to the passphrase
		passphrase = append(passphrase, wordList[index.Int64()])
	}

	// Combine the words into a single string separated by spaces
	return strings.Join(passphrase, " "), nil
}

func loadWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return words, nil
}

func generateRandomKey() string {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("Failed to generate random bytes: %v", err)
	}

	randomHex := hex.EncodeToString(randomBytes)

	return randomHex
}

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey
}

func signData(privateKey *ecdsa.PrivateKey, data string) []byte {
	hash := sha256.Sum256([]byte(data))
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Fatal(err)
	}
	return signature
}

func verifySignature(publicKey *ecdsa.PublicKey, data string, signature []byte) bool {
	hash := sha256.Sum256([]byte(data))
	return ecdsa.VerifyASN1(publicKey, hash[:], signature)
}

func TestVerifySignature() {
	privateKey, publicKey := generateKeyPair()

	signature := signData(privateKey, "hello world")

	fmt.Println(verifySignature(publicKey, "hello world", signature))
}

func EncodeBase58(input string) []byte {
	return base58.Encode([]byte(input))
}

func Sha256Hash(input string) string {
	h := sha256.New()

	h.Write([]byte(input))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x\n", bs)
}

func CreateNewWallet() Wallet {
	privateKey, publicKey := generateKeyPair()

	primaryPublicKeyHash := Sha256Hash(publicKey.X.Text(16))
	secondaryPublicKeyHash := Sha256Hash(primaryPublicKeyHash)

	address, err := HexToBase58(secondaryPublicKeyHash)

	if err != nil {
		fmt.Println(err)
		panic("AHHH")
	}

	return Wallet{privateKey, publicKey, address}
}
