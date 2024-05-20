package clients_test

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/machinefi/sprout/clients"
)

func TestClientByClientID(t *testing.T) {
	projectClientContractAddr := "0xF4d6282C5dDD474663eF9e70c927c0d4926d1CEb"
	ioIDContractAddr := "0x06b3Fcda51e01EE96e8E8873F0302381c955Fddd"
	w3bstreamProjectContractAddr := "0x6AfCB0EB71B7246A68Bb9c0bFbe5cD7c11c4839f"
	chainEndpoint := "https://babel-api.testnet.iotex.io"
	ioRegistryEndpoint := "did.iotex.me"

	mgr, err := clients.NewManager(projectClientContractAddr, ioIDContractAddr, w3bstreamProjectContractAddr, ioRegistryEndpoint, chainEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ioID contract address:              ", ioIDContractAddr)
	t.Log("project device contract address:    ", projectClientContractAddr)
	t.Log("w3bstream project contract address: ", projectClientContractAddr)
	t.Log("chain endpoint:                     ", chainEndpoint)
	t.Log("io registry:                        ", ioRegistryEndpoint)

	clientDIDs := []string{
		"did:io:0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73",
		// "did:io:0x1c89860d3eed129fe1996bb72044cc22cc46a756",
		// "did:io:0x5b7902df415485c7e21334ca95ca94667278030e",
	}
	for _, id := range clientDIDs {
		client := mgr.ClientByIoID(id)
		if client == nil {
			t.Logf("client is not fetched: %s ", id)
			continue
		}

		address := strings.TrimPrefix(id, "did:io:")
		t.Log(common.HexToAddress(address).String())

		approved, err := mgr.HasProjectPermission(id, 21)
		if err != nil {
			t.Logf("permission is not validated: %v", err)
			continue
		}

		t.Log("client DID: ", id)
		t.Log("project permission: ", approved)
	}

	_, err = mgr.HasProjectPermission("did:io:0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73", 23)
	if err != nil {
		t.Fatal(err)
	}
}
