package clients_test

import (
	"testing"

	"github.com/machinefi/sprout/clients"
)

func TestClientByClientID(t *testing.T) {
	contractAddr := "0xB63e6034361283dc8516480a46EcB9C122c983Bb"
	chainEndpoint := "https://babel-api.testnet.iotex.io"
	ioRegistryEndpoint := "did.iotex.me"
	clientDID := "did:io:0x1c89860d3eed129fe1996bb72044cc22cc46a756"

	mgr, err := clients.NewManager(contractAddr, chainEndpoint, ioRegistryEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	client := mgr.ClientByIoID(clientDID)
	if client == nil {
		t.Log("client is not fetched")
		return
	}

	t.Log("client DID:", clientDID)
	t.Log("client DID:", client.DID())
	t.Log("client KID:", client.KeyAgreementKID())

	client = mgr.ClientByIoID("did:io:0x5b7902df415485c7e21334ca95ca94667278030e")

	if client == nil {
		t.Log("client is not fetched")
		return
	}

	t.Log("client DID:", clientDID)
	t.Log("client DID:", client.DID())
	t.Log("client KID:", client.KeyAgreementKID())
}
