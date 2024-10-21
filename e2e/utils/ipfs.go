package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/iotexproject/w3bstream/util/ipfs"
)

type ipfsContainer struct {
	testcontainers.Container
}

func SetupIPFS() (*ipfsContainer, string, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ipfs/go-ipfs:latest",
		ExposedPorts: []string{"5001/tcp", "8080/tcp", "4001/tcp"},
		WaitingFor:   wait.ForHTTP("/version").WithPort("5001/tcp").WithStartupTimeout(2 * time.Minute),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	if err != nil {
		return nil, "", err
	}

	mapPort, err := container.MappedPort(ctx, "5001")
	if err != nil {
		return nil, "", err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, "", err
	}

	apiURL := fmt.Sprintf("http://%s:%s", ip, mapPort.Port())

	return &ipfsContainer{Container: container}, apiURL, nil
}

func TestIPfs2(apiURL string) {
	ipfs := ipfs.NewIPFS(apiURL)

	// Add file to IPFS
	cid, err := ipfs.AddContent([]byte("Hello, IPFS!"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File added with CID:", cid)

	// Get file from IPFS
	content, err := ipfs.Cat(cid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File content:", string(content))
}
