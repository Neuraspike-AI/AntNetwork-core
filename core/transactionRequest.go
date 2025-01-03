package main

type transactionRequest struct {
	txnID                string
	sender               string
	blockNonceSignature  []byte
	blockNonce           int
	reciever             string
	amount               float32
	txnHash              string
	proofLastTransaction txn
}
