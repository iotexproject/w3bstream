//go:build e2e
// +build e2e

package e2e

import (
	"context"
	"log"
	"runtime"
	"testing"

	"github.com/iotexproject/w3bstream/e2e/utils"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
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
