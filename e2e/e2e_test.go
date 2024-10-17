package e2e

import (
	"context"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/iotexproject/w3bstream/e2e/utils"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	if os.Getenv("TEST_E2E") != "true" {
		t.Skip("Skipping TestE2E as requested.")
	}
	if runtime.GOARCH == "arm64" {
		log.Println("Skipping tests: Unsupported architecture (arm64)")
		t.Skip()
	}

	chainContainer, endpoint, err := utils.SetupLocalChain(context.Background())
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := chainContainer.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate chain container: %v", err)
		}
	})

	log.Printf("Chain endpoint: %s\n", endpoint)

	err = utils.DeployContract(endpoint)
	require.NoError(t, err)
}
