package main

import (
	"encoding/json"
	"fmt"
)

type Block struct {
	Height    int              `bson:"height"`
	Hash      string           `bson:"hash"`
	Txns      [256]Transaction `bson:"txns"`
	Timestamp int64            `bson:"tmstmp"`
	IsGenesis bool             `bson:"isGns"`
	HashPrev  string           `bson:"hashPrev"`
}

func (block *Block) toJSON() []byte {
	jsonData, err := json.MarshalIndent(block, "", "  ")
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return []byte{0}
	}
	return jsonData
}

func (block *Block) getHash() string {
	str := block.toJSON()
	return Sha256Hash(string(str))
}
