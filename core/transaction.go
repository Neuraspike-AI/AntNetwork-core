package main

type TransactionType int

const (
	P2P = iota
	BLOCK_REWARD
	FEE
)

type TransactionID struct {
	TransactionID int `bson:"txnID"`
	BlockHeight   int `bson:"blHeight"`
}

type Transaction struct {
	Sender               string          `bson:"sndr"`
	SenderSignature      []byte          `bson:"sndrSig"`
	Reciever             string          `bson:"rcvr"`
	TransactionTimestamp int64           `bson:"txnTmstmp"`
	Amount               float32         `bson:"amnt"`
	TransactionHash      string          `bson:"txnHash"`
	Message              string          `bson:"msg"`
	Signature            string          `bson:"sig"`
	ProofLasttransaction TransactionID   `bson:"prfTxn"`
	TypeTransaction      TransactionType `bson:"typeTxn"`
	FinalBalance         float32         `bson:"fBlnc"`
}
