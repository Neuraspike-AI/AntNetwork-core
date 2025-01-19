package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"time"
)

type transactionRequest struct {
	TransactionID        string          `bson:"txnID"`
	Sender               string          `bson:"sndr"`
	Reciever             string          `bson:"rcvr"`
	Amount               float32         `bson:"amnt"`
	TransactionHash      string          `bson:"txnHash"`
	ProofLasttransaction TransactionID   `bson:"prfTxn"`
	TransactionTimestamp int64           `bson:"txnTmstmp"`
	Message              string          `bson:"msg"`
	Signature            []byte          `bson:"sig"`
	PublicKey            ecdsa.PublicKey `bson:"pbKey"`
}

// Verifies transaction, returns true if accepted
func (transactionRequest *transactionRequest) VerifytransactionRequest(wrapper DBWrapper) bool {
	//Verify timestamp, future timestamps are illegal
	if transactionRequest.TransactionTimestamp >= time.Now().UnixMilli() {
		return false
	}
	//Check transaction hash
	//
	jsonStr, err := json.Marshal(transactionRequest)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	jsonString := string(jsonStr)
	hash := Sha256Hash(jsonString)
	if hash != transactionRequest.TransactionHash {
		return false
	}
	//Verify signatures
	if !verifySignature(&transactionRequest.PublicKey, transactionRequest.TransactionHash, transactionRequest.Signature) {
		return false
	}
	//Verify origin of funds TODO: Add posiblity of unverified UTXO
	//verify ownership of proof transaction
	transactionProof, err := getTransactionFromID(transactionRequest.ProofLasttransaction, wrapper)

	if err != nil {
		return false
	}

	if transactionProof.Sender != transactionRequest.Sender {
		return false
	}

	//Verify latest transaction
	

	return true
}
