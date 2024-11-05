package utils

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type VMContainer struct {
	testcontainers.Container
}

func SetupRisc0VM() (*VMContainer, string, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/iotexproject/risc0server:v0.0.13",
		ExposedPorts: []string{"4001/tcp"},
		WaitingFor:   wait.ForListeningPort("4001"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	if err != nil {
		return nil, "", err
	}

	mapPort, err := container.MappedPort(ctx, "4001")
	if err != nil {
		return nil, "", err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, "", err
	}

	endpoint := fmt.Sprintf("%s:%s", ip, mapPort.Port())

	return &VMContainer{Container: container}, endpoint, nil
}
