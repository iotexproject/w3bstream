package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	b, err := hexutil.Decode("0x" + "048877758949b314f9da52ebcb46c19a085c1ad12a7509228de04b24784e815e8d7b726e2564cb9f03175feb70fca96d2f5f117762304d8621afc424ebcb92f39c")
	if err != nil {
		panic(err)
	}
	pub, err := crypto.UnmarshalPubkey(b)
	if err != nil {
		panic(err)
	}
	addr := crypto.PubkeyToAddress(*pub)

	fmt.Println(addr.String())
}
