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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/e2e/services"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/smartcontracts/go/mockdappmovementbatch"
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
	gnarkVMContainer, gnarkVMEndpoint, err := services.SetupGnarkVM()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := gnarkVMContainer.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate vm container: %v", err)
		}
	})

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
	sendETH(t, chainEndpoint, payerHex, sequencer.Address(), 200)
	err = sequencer.Start()
	require.NoError(t, err)
	defer sequencer.Stop()

	// Prover init
	tempProverDB, err := os.CreateTemp("", "prover.db")
	require.NoError(t, err)
	defer os.Remove(tempProverDB.Name())
	defer tempProverDB.Close()
	prover, proverKey, err := proverInit(chDSN, tempProverDB.Name(), chainEndpoint,
		map[int]string{1: gnarkVMEndpoint}, contracts)
	require.NoError(t, err)
	err = prover.Start()
	require.NoError(t, err)
	defer prover.Stop()

	// Register prover
	proverAddr := crypto.PubkeyToAddress(proverKey.PublicKey)
	sendETH(t, chainEndpoint, payerHex, proverAddr, 20)
	err = registerProver(t, chainEndpoint, contracts, proverKey)
	require.NoError(t, err)

	// Register device
	deviceKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	deviceAddr := crypto.PubkeyToAddress(deviceKey.PublicKey)
	sendETH(t, chainEndpoint, payerHex, deviceAddr, 20)

	// Register device2
	deviceKey2, err := crypto.GenerateKey()
	require.NoError(t, err)
	deviceAddr2 := crypto.PubkeyToAddress(deviceKey2.PublicKey)
	sendETH(t, chainEndpoint, payerHex, deviceAddr2, 20)

	// Register project
	projectOwnerKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	projectOwnerAddr := crypto.PubkeyToAddress(projectOwnerKey.PublicKey)
	sendETH(t, chainEndpoint, payerHex, projectOwnerAddr, 20)
	projectID := big.NewInt(1)
	registerIoID(t, chainEndpoint, contracts, deviceKey, projectID)
	registerIoID(t, chainEndpoint, contracts, deviceKey2, projectID)
	registerProject(t, chainEndpoint, contracts, projectOwnerKey, projectID)

	t.Run("gnark", func(t *testing.T) {
		bindProjectDapp(t, chainEndpoint, contracts, projectOwnerKey, projectID, common.HexToAddress(contracts.MockDapp))
		gnarkCodePath := "./testdata/gnark.code"
		gnarkMetadataPath := "./testdata/gnark.metadata"
		project := &project.Project{DefaultVersion: "v1", Configs: []*project.Config{{Version: "v1", VMTypeID: 1}}}
		// Upload project
		uploadProject(t, chainEndpoint, ipfsEndpoint, project, &gnarkCodePath, &gnarkMetadataPath, contracts, projectOwnerKey, projectID)
		require.NoError(t, err)
		// Wait a few seconds for the device info synced on api node
		time.Sleep(2 * time.Second)
		// Send message: prove 1+1=2
		data, err := hex.DecodeString("00000001000000010000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001")
		require.NoError(t, err)
		taskid := sendMessage(t, data, projectID, project.Configs[0], deviceKey, apiNodeUrl)
		waitSettled(t, taskid, apiNodeUrl, 30*time.Second, 30*time.Second)
	})
	t.Run("gnark-liveness", func(t *testing.T) {
		t.Skip()
		bindProjectDapp(t, chainEndpoint, contracts, projectOwnerKey, projectID, common.HexToAddress(contracts.MockDappLiveness))
		gnarkCodePath := "./testdata/pebble.circuit"
		gnarkMetadataPath := "./testdata/pebble.pk"
		project := &project.Project{
			DefaultVersion: "v2",
			Configs: []*project.Config{{
				Version:    "v2",
				VMTypeID:   1,
				ProofType:  "liveness",
				SignedKeys: []project.SignedKey{{Name: "timestamp", Type: "uint64"}},
			}}}

		// Upload project
		uploadProject(t, chainEndpoint, ipfsEndpoint, project, &gnarkCodePath, &gnarkMetadataPath, contracts, projectOwnerKey, projectID)
		require.NoError(t, err)

		// Wait a few seconds for the device info synced on api node
		time.Sleep(2 * time.Second)

		data, err := json.Marshal(struct {
			Timestamp uint64 `json:"timestamp"`
		}{
			Timestamp: uint64(time.Now().Unix()),
		})
		require.NoError(t, err)
		taskid := sendMessage(t, data, projectID, project.Configs[0], deviceKey, apiNodeUrl)
		waitSettled(t, taskid, apiNodeUrl, 30*time.Second, 30*time.Second)
	})
	t.Run("gnark-movement", func(t *testing.T) {
		bindProjectDapp(t, chainEndpoint, contracts, projectOwnerKey, projectID, common.HexToAddress(contracts.MockDappMovementBatch))
		project := &project.Project{
			DefaultVersion: "v3",
			Configs: []*project.Config{{
				Version:             "v3",
				VMTypeID:            1,
				ProofType:           "movement",
				Code:                "ipfs://ipfs.mainnet.iotex.io/QmSg3NFgVbHZyUZsHAjrLV7o8CmLKYChNa3CVi4uKKrHeL",
				CodeHash:            "0xdc3392204a56698891090ea1b729b7a25deba50ccad26ef70ff4b3d51f662ce4",
				Metadata:            "ipfs://ipfs.mainnet.iotex.io/QmTXxKjQEQ8gAzgJjvoxoMB5ocYi16VimHjFwTKf6hEHSm",
				MetadataHash:        "0xc969f436f4cef28377e3b5ec3ca2457bcaab2795f65c4e1bf3656b57087cd957",
				TaskProcessingBatch: 10,
				SignedKeys: []project.SignedKey{
					{Name: "timestamp", Type: "uint64"},
					{Name: "latitude", Type: "uint64"},
					{Name: "longitude", Type: "uint64"}},
			}}}

		uploadProject(t, chainEndpoint, ipfsEndpoint, project, nil, nil, contracts, projectOwnerKey, projectID)
		// Wait a few seconds for the device info synced on api node
		time.Sleep(2 * time.Second)

		timestamp := uint64(time.Now().Unix())
		lastTimestamp := timestamp - 60
		latitude := uint64(1200)
		longitude := uint64(200)
		lastLatitude := latitude - 1100
		lastLongitude := longitude - 1

		data, err := json.Marshal(struct {
			Timestamp uint64 `json:"timestamp"`
			Latitude  uint64 `json:"latitude"`
			Longitude uint64 `json:"longitude"`
		}{
			Timestamp: timestamp,
			Latitude:  latitude,
			Longitude: longitude,
		})
		require.NoError(t, err)
		lastData, err := json.Marshal(struct {
			Timestamp uint64 `json:"timestamp"`
			Latitude  uint64 `json:"latitude"`
			Longitude uint64 `json:"longitude"`
		}{
			Timestamp: lastTimestamp,
			Latitude:  lastLatitude,
			Longitude: lastLongitude,
		})
		require.NoError(t, err)
		taskID := ""
		for i := 0; i < 7; i++ {
			_ = sendMessage(t, lastData, projectID, project.Configs[0], deviceKey, apiNodeUrl)
			taskID = sendMessage(t, data, projectID, project.Configs[0], deviceKey, apiNodeUrl)
			time.Sleep(2 * time.Second)
		}
		for i := 0; i < 3; i++ {
			_ = sendMessage(t, lastData, projectID, project.Configs[0], deviceKey, apiNodeUrl)
			taskID = sendMessage(t, lastData, projectID, project.Configs[0], deviceKey, apiNodeUrl)
			time.Sleep(2 * time.Second)
		}
		// for i := 0; i < 3; i++ {
		// 	_ = sendMessage(t, lastData, projectID, project.Configs[0], deviceKey2, apiNodeUrl)
		// 	taskID = sendMessage(t, data, projectID, project.Configs[0], deviceKey2, apiNodeUrl)
		// 	time.Sleep(2 * time.Second)
		// }
		// for i := 0; i < 1; i++ {
		// 	_ = sendMessage(t, lastData, projectID, project.Configs[0], deviceKey2, apiNodeUrl)
		// 	taskID = sendMessage(t, lastData, projectID, project.Configs[0], deviceKey2, apiNodeUrl)
		// 	time.Sleep(2 * time.Second)
		// }
		waitSettled(t, taskID, apiNodeUrl, 30*time.Second, 10*time.Minute)

		client, err := ethclient.Dial(chainEndpoint)
		require.NoError(t, err)
		mockDappMovementBatchContract, err := mockdappmovementbatch.NewMockdappmovementbatch(
			common.HexToAddress(contracts.MockDappMovementBatch), client)
		require.NoError(t, err)

		n, err := mockDappMovementBatchContract.DeviceTick(nil, crypto.PubkeyToAddress(deviceKey.PublicKey))
		require.NoError(t, err)
		require.Equal(t, n, uint64(7), "unexpect tick count")

		// n, err = mockDappMovementBatchContract.DeviceTick(nil, crypto.PubkeyToAddress(deviceKey2.PublicKey))
		// require.NoError(t, err)
		// require.Equal(t, n, uint64(3), "unexpect tick count")
	})
}

func sendMessage(t *testing.T, dataJson []byte, projectID *big.Int,
	projectConfig *project.Config, deviceKey *ecdsa.PrivateKey, apiNodeUrl string) string {
	reqBody, err := signMesssage(dataJson, projectID.Uint64(), projectConfig, deviceKey)
	require.NoError(t, err)

	taskID, err := createTask(reqBody, apiNodeUrl)
	require.NoError(t, err)

	return taskID
}

func waitSettled(t *testing.T, taskID string, apiNodeUrl string, assigned, settled time.Duration) {
	err := waitUntil(func() (bool, error) {
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
	}, assigned)
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
	}, settled)
	require.NoError(t, err)
}
