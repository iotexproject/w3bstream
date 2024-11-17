package e2e

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/e2e/utils"
	"github.com/iotexproject/w3bstream/service/apinode/api"
)

const (
	// private keys in Anvil local chain
	payerHex        = "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
	projectFilePath = "./testdata/risc0"
)

func TestE2E(t *testing.T) {
	if os.Getenv("TEST_E2E") != "true" {
		t.Skip("Skipping E2E tests.")
	}
	var chainEndpoint string
	if runtime.GOARCH == "arm64" {
		chainEndpoint = "http://localhost:8545"
		log.Printf("Using local chain at %s", chainEndpoint)
		if err := utils.TestChain(chainEndpoint); err != nil {
			t.Fatalf("failed to test chain %s: %v", chainEndpoint, err)
		}
	} else {
		// Setup local chain
		chainContainer, endpoint, err := utils.SetupLocalChain()
		t.Cleanup(func() {
			if err := chainContainer.Terminate(context.Background()); err != nil {
				t.Logf("failed to terminate chain container: %v", err)
			}
		})
		require.NoError(t, err)
		chainEndpoint = endpoint
	}

	// Deploy contract to local chain
	contracts, err := utils.DeployContract(chainEndpoint, payerHex)
	require.NoError(t, err)

	// Setup postgres
	dbName := "users"
	pgContainer, PGURI, err := utils.SetupPostgres(dbName)
	t.Cleanup(func() {
		if err := pgContainer.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate postgres container: %v", err)
		}
	})
	require.NoError(t, err)

	// Setup IPFS
	ipfsContainer, ipfsEndpoint, err := utils.SetupIPFS()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := ipfsContainer.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate ipfs container: %v", err)
		}
	})

	// Setup VM
	vmContainer, vmEndpoint, err := utils.SetupRisc0VM()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := vmContainer.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate vm container: %v", err)
		}
	})

	// APINode init
	apiNode, apiNodeUrl, err := apiNodeInit(PGURI, chainEndpoint, contracts.TaskManager)
	require.NoError(t, err)
	err = apiNode.Start()
	require.NoError(t, err)
	defer apiNode.Stop()

	// Sequencer init
	tempSequencerDB, err := os.CreateTemp("", "sequencer.db")
	require.NoError(t, err)
	defer os.Remove(tempSequencerDB.Name())
	defer tempSequencerDB.Close()
	sequencer, err := sequencerInit(PGURI, tempSequencerDB.Name(), chainEndpoint, contracts)
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
	prover, proverKey, err := proverInit(PGURI, tempProverDB.Name(), chainEndpoint, vmEndpoint, contracts)
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
	projectID, err := registerProject(t, chainEndpoint, ipfsEndpoint, projectFilePath, contracts, projectOwnerKey)
	require.NoError(t, err)

	// Register prover
	proverAddr := crypto.PubkeyToAddress(proverKey.PublicKey)
	err = sendETH(t, chainEndpoint, payerHex, proverAddr, 20)
	require.NoError(t, err)
	err = registerProver(t, chainEndpoint, contracts, proverKey)
	require.NoError(t, err)

	// Send message
	msgData := struct {
		PrivateInput string `json:"private_input"`
		PublicInput  string `json:"public_input"`
		ReceiptType  string `json:"receipt_type"`
	}{
		PrivateInput: "14",
		PublicInput:  "3,34",
		ReceiptType:  "Stark",
	}
	dataJson, err := json.Marshal(msgData)
	require.NoError(t, err)

	senderKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	reqBody, err := signMesssage(dataJson, projectID.Uint64(), senderKey)
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

func sendETH(t *testing.T, chainEndpoint string, payerHex string, toAddress common.Address, amount uint64) error {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return err
	}
	defer client.Close()

	// 2. Load the sender's private key
	privateKey, err := crypto.HexToECDSA(payerHex) // Replace with actual private key
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// 3. Get the sender's address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 4. Get the current nonce for the sender's account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	require.NoError(t, err)

	// 5. Define transaction parameters
	value := big.NewInt(0).Mul(big.NewInt(int64(amount)), big.NewInt(1e18)) // Amount in Wei (1 ETH = 10^18 Wei)
	gasLimit := uint64(21000)                                               // Gas limit for simple ETH transfer
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get suggested gas price: %v", err)
	}

	// 6. Create the transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// 7. Sign the transaction with the sender's private key
	chainID, err := client.NetworkID(context.Background())
	require.NoError(t, err)
	signedTx, err := types.SignTx(tx, types.NewEIP2930Signer(chainID), privateKey)
	require.NoError(t, err)

	// 8. Send the signed transaction
	err = client.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)

	_, err = utils.WaitForTransactionReceipt(client, signedTx.Hash())
	require.NoError(t, err)

	return nil
}

func signMesssage(data []byte, projectID uint64, key *ecdsa.PrivateKey) ([]byte, error) {
	req := &api.CreateTaskReq{
		Nonce:          uint64(time.Now().Unix()),
		ProjectID:      projectID,
		ProjectVersion: "v1.0.0",
		Payloads:       []string{hexutil.Encode(data)},
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	h := crypto.Keccak256Hash(reqJson)
	sig, err := crypto.Sign(h.Bytes(), key)
	if err != nil {
		return nil, err
	}

	req.Signature = hexutil.Encode(sig)

	return json.Marshal(req)
}

func createTask(body []byte, apiurl string) (string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/task", apiurl), bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrapf(err, "failed to create task, status code: %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var handleMessageResp api.CreateTaskResp
	if err := json.Unmarshal(buf.Bytes(), &handleMessageResp); err != nil {
		return "", errors.Wrap(err, "failed to deserialize response body")
	}

	return handleMessageResp.TaskID, nil
}

func queryTask(taskID string, apiurl string) (*api.QueryTaskResp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/task/%s", apiurl, taskID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("failed to query task, status code: %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var queryTaskResp api.QueryTaskResp
	if err := json.Unmarshal(buf.Bytes(), &queryTaskResp); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize response body")
	}

	return &queryTaskResp, nil
}

func waitUntil(f func() (bool, error), timeOut time.Duration) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutCh := time.After(timeOut)

	for {
		select {
		case <-timeoutCh:
			return fmt.Errorf("timeout waiting for condition")
		case <-ticker.C:
			ok, err := f()
			if err != nil {
				return err
			}
			if ok {
				return nil
			}
		}
	}
}
