package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/apinode/api"
)

func main() {
	prv, err := crypto.HexToECDSA("your private key")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse private key"))
	}
	req := &api.HandleMessageReq{
		ProjectID:      912,
		ProjectVersion: "v1.0.0",
		Data:           "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}",
	}
	reqJson, _ := json.Marshal(req)
	fmt.Println(string(reqJson))

	h := crypto.Keccak256Hash(reqJson)
	sig, err := crypto.Sign(h.Bytes(), prv)
	if err != nil {
		panic(err)
	}
	fmt.Println(hexutil.Encode(sig))
}
