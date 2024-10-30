package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type chainContainer struct {
	testcontainers.Container
}

func SetupLocalChain() (*chainContainer, string, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:latest",
		ExposedPorts: []string{"8545/tcp"},
		// Entrypoint:   []string{"anvil", "--block-time", "5", "--host", "0.0.0.0"},
		Entrypoint: []string{"anvil", "--host", "0.0.0.0"},
		WaitingFor: wait.ForListeningPort("8545"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	if err != nil {
		return nil, "", err
	}

	mapPort, err := container.MappedPort(ctx, "8545")
	if err != nil {
		return nil, "", err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, "", err
	}

	endpoint := fmt.Sprintf("http://%s:%s", ip, mapPort.Port())

	if err := TestChain(endpoint); err != nil {
		return nil, "", err
	}

	return &chainContainer{Container: container}, endpoint, nil
}

func TestChain(endpoint string) error {
	ethClient, err := ethclient.Dial(endpoint)
	if err != nil {
		return err
	}
	_, err = ethClient.NetworkID(context.Background())
	if err != nil {
		return err

	}
	return nil
}

var (
	txTimeout = 10 * time.Second
)

// WaitForTransactionReceipt waits for the transaction receipt or returns an error if it times out
func WaitForTransactionReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), txTimeout)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("transaction receipt timeout")
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(context.Background(), txHash)
			if err == nil && receipt != nil {
				if receipt.Status != types.ReceiptStatusSuccessful {
					return nil, errors.New("transaction failed")
				}
				return receipt, nil
			}
		}
	}
}
