package main

import (
	"fmt"
	"time"
)

func main() {
	block := Block{Height: 0, Hash: "0", Txns: [256]Transaction{}, Timestamp: time.Now().UnixMilli(), IsGenesis: true, HashPrev: "0"}
	fmt.Printf("%+v\n", block)
	sha := Sha256Hash(string(block.toJSON()))
	wrapper := initiateDB()
	addBlockDB(wrapper, block)
	err, block_new := getBlockByHeight(0, wrapper.collection, wrapper.context)
	if err != nil {
		panic("Err")
	}
	sha_new := Sha256Hash(string(block_new.toJSON()))
	if sha != sha_new {
		panic("Not equals")
	} else {
		fmt.Println("Succesful")
	}
	disconnectDB(wrapper.client, wrapper.context, wrapper.cancel)
}
