package utils

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type chainContainer struct {
	testcontainers.Container
}

func SetupLocalChain(ctx context.Context) (*chainContainer, string, error) {
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:latest",
		ExposedPorts: []string{"8545/tcp"},
		Entrypoint:   []string{"anvil", "--block-time", "5", "--host", "0.0.0.0"},
		WaitingFor:   wait.ForListeningPort("8545"),
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

	ethClient, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, "", err
	}
	_, err = ethClient.NetworkID(ctx)
	if err != nil {
		return nil, "", err

	}

	return &chainContainer{Container: container}, endpoint, nil
}
