package e2e

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/e2e/services"
	"github.com/iotexproject/w3bstream/project"
)

const (
	// private keys in Anvil local chain
	payerHex = "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
)

func TestE2E(t *testing.T) {
	if os.Getenv("TEST_E2E") != "true" {
		t.Skip("Skipping E2E tests.")
	}
	var chainEndpoint string
	if runtime.GOARCH == "arm64" {
		chainEndpoint = "http://localhost:8545"
		log.Printf("Using local chain at %s", chainEndpoint)
		if err := services.TestChain(chainEndpoint); err != nil {
			t.Fatalf("failed to test chain %s: %v", chainEndpoint, err)
		}
	} else {
		// Setup local chain
		chainContainer, endpoint, err := services.SetupLocalChain()
		t.Cleanup(func() {
			if err := chainContainer.Terminate(context.Background()); err != nil {
				t.Logf("failed to terminate chain container: %v", err)
			}
		})
		require.NoError(t, err)
		chainEndpoint = endpoint
	}

	// Deploy contract to local chain
	contracts, err := services.DeployContract(chainEndpoint, payerHex)
	require.NoError(t, err)

	// Setup clickhouse
	dbName := "w3bstream"
	chContainer, chDSN, err := services.SetupClickhouse(dbName)
	t.Cleanup(func() {
		if err := chContainer.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate clickhouse container: %v", err)
		}
	})
	require.NoError(t, err)

	// Setup IPFS
	ipfsContainer, ipfsEndpoint, err := services.SetupIPFS()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := ipfsContainer.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate ipfs container: %v", err)
		}
	})

	// Setup VM
	risc0VMContainer, risc0VMEndpoint, err := services.SetupRisc0VM()
	require.NoError(t, err)
	gnarkVMContainer, gnarkVMEndpoint, err := services.SetupGnarkVM()
	require.NoError(t, err)

	// APINode init
	tempApiNodeDB, err := os.CreateTemp("", "apinode.db")
	require.NoError(t, err)
	defer os.Remove(tempApiNodeDB.Name())
	defer tempApiNodeDB.Close()
	apiNode, apiNodeUrl, err := apiNodeInit(chDSN, tempApiNodeDB.Name(), chainEndpoint, contracts)
	require.NoError(t, err)
	err = apiNode.Start()
	require.NoError(t, err)
	defer apiNode.Stop()

	// Sequencer init
	tempSequencerDB, err := os.CreateTemp("", "sequencer.db")
	require.NoError(t, err)
	defer os.Remove(tempSequencerDB.Name())
	defer tempSequencerDB.Close()
	sequencer, err := sequencerInit(chDSN, tempSequencerDB.Name(), chainEndpoint, contracts)
	require.NoError(t, err)
	err = sendETH(t, chainEndpoint, payerHex, sequencer.Address(), 200)
	require.NoError(t, err)
	err = sequencer.Start()
	require.NoError(t, err)
	defer sequencer.Stop()

	// Prover init
	tempProverDB, err := os.CreateTemp("", "prover.db")
	require.NoError(t, err)
	defer os.Remove(tempProverDB.Name())
	defer tempProverDB.Close()
	prover, proverKey, err := proverInit(chDSN, tempProverDB.Name(), chainEndpoint,
		map[int]string{1: risc0VMEndpoint, 5: gnarkVMEndpoint}, contracts)
	require.NoError(t, err)
	err = prover.Start()
	require.NoError(t, err)
	defer prover.Stop()

	// Register project
	projectOwnerKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	projectOwnerAddr := crypto.PubkeyToAddress(projectOwnerKey.PublicKey)
	err = sendETH(t, chainEndpoint, payerHex, projectOwnerAddr, 20)
	require.NoError(t, err)
	projectID, err := registerProject(t, chainEndpoint, contracts, projectOwnerKey)
	require.NoError(t, err)

	// Register prover
	proverAddr := crypto.PubkeyToAddress(proverKey.PublicKey)
	err = sendETH(t, chainEndpoint, payerHex, proverAddr, 20)
	require.NoError(t, err)
	err = registerProver(t, chainEndpoint, contracts, proverKey)
	require.NoError(t, err)

	// Register device
	deviceKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	deviceAddr := crypto.PubkeyToAddress(deviceKey.PublicKey)
	err = sendETH(t, chainEndpoint, payerHex, deviceAddr, 20)
	require.NoError(t, err)
	registerIoID(t, chainEndpoint, contracts, deviceKey, projectID)

	t.Run("RISC0", func(t *testing.T) {
		t.Cleanup(func() {
			if err := risc0VMContainer.Terminate(context.Background()); err != nil {
				t.Logf("failed to terminate vm container: %v", err)
			}
		})
		risc0CodePath := "./testdata/risc0.code"
		project := &project.Project{Configs: []*project.Config{{Version: "v1", VMTypeID: 1}}}

		// Upload project
		uploadProject(t, chainEndpoint, ipfsEndpoint, project, &risc0CodePath, nil, contracts, projectOwnerKey, projectID, false)
		require.NoError(t, err)

		// Wait a few seconds for the device info synced on api node
		time.Sleep(2 * time.Second)

		// Send message
		msgData := struct {
			PrivateInput string `json:"private_input"`
			PublicInput  string `json:"public_input"`
			ReceiptType  string `json:"receipt_type"`
			Timestamp    uint64 `json:"timestamp"`
		}{
			PrivateInput: "14",
			PublicInput:  "3,34",
			ReceiptType:  "Stark",
			Timestamp:    uint64(time.Now().Unix()),
		}
		dataJson, err := json.Marshal(msgData)
		require.NoError(t, err)

		sendMessage(t, dataJson, projectID, nil, deviceKey, apiNodeUrl)
	})

	t.Run("GNARK", func(t *testing.T) {
		t.Cleanup(func() {
			if err := gnarkVMContainer.Terminate(context.Background()); err != nil {
				t.Logf("failed to terminate vm container: %v", err)
			}
		})
		gnarkCodePath := "./testdata/gnark.code"
		gnarkMetadataPath := "./testdata/gnark.metadata"
		project := &project.Project{Configs: []*project.Config{{Version: "v1", VMTypeID: 5}}}

		// Upload project
		uploadProject(t, chainEndpoint, ipfsEndpoint, project, &gnarkCodePath, &gnarkMetadataPath, contracts, projectOwnerKey, projectID, true)
		require.NoError(t, err)

		// Wait a few seconds for the device info synced on api node
		time.Sleep(2 * time.Second)

		// Send message: prove 1+1=2
		data, err := hex.DecodeString("00000001000000010000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001")
		require.NoError(t, err)

		sendMessage(t, data, projectID, nil, deviceKey, apiNodeUrl)
	})
}

func sendMessage(t *testing.T, dataJson []byte, projectID *big.Int,
	projectConfig *project.Config, deviceKey *ecdsa.PrivateKey, apiNodeUrl string) {
	reqBody, err := signMesssage(dataJson, projectID.Uint64(), projectConfig, deviceKey)
	require.NoError(t, err)

	taskID, err := createTask(reqBody, apiNodeUrl)
	require.NoError(t, err)

	err = waitUntil(func() (bool, error) {
		states, err := queryTask(taskID, apiNodeUrl)
		if err != nil {
			return false, err
		}
		for _, state := range states.States {
			if state.State == "assigned" {
				return true, nil
			}
		}
		return false, nil
	}, 30*time.Second)
	require.NoError(t, err)

	err = waitUntil(func() (bool, error) {
		states, err := queryTask(taskID, apiNodeUrl)
		if err != nil {
			return false, err
		}
		for _, state := range states.States {
			if state.State == "settled" {
				return true, nil
			}
		}
		return false, nil
	}, 120*time.Second)
	require.NoError(t, err)
}
