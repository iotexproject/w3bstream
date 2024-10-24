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
	// if os.Getenv("TEST_E2E") != "true" {
	// 	t.Skip("Skipping E2E tests.")
	// }
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
	pgContainer, URI, err := utils.SetupPostgres(dbName)
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

	// Bootnode init
	bootnode, err := bootNodeInit()
	require.NoError(t, err)
	bootnode.Start()
	defer bootnode.Stop()

	// APINode init
	apiNode, err := apiNodeInit(URI, chainEndpoint, bootnode.Addrs()[1], contracts.TaskManager)
	require.NoError(t, err)
	apiNode.Start()
	defer apiNode.Stop()
	apiNodeUrl := fmt.Sprintf("http://localhost%s", apiNode.Config.ServiceEndpoint)

	// Sequencer init
	tempSequencerDB, err := os.CreateTemp("", "sequencer.db")
	require.NoError(t, err)
	defer os.Remove(tempSequencerDB.Name())
	defer tempSequencerDB.Close()
	sequencer, err := sequencerInit(tempSequencerDB.Name(), chainEndpoint, bootnode.Addrs()[1],
		contracts)
	require.NoError(t, err)
	err = sendETH(t, chainEndpoint, payerHex, sequencer.Address(), 200)
	require.NoError(t, err)

	sequencer.Start()
	defer sequencer.Stop()

	// Register project
	projectID, err := registerProject(t, chainEndpoint, ipfsEndpoint, projectFilePath, contracts, payerHex)
	require.NoError(t, err)

	// Register prover
	proverKey, err := crypto.GenerateKey()
	require.NoError(t, err)
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
		ReceiptType:  "Snark",
	}
	dataJson, err := json.Marshal(msgData)
	require.NoError(t, err)

	senderKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	reqBody, err := signMesssage(dataJson, projectID.Uint64(), senderKey)
	require.NoError(t, err)

	err = sendMessage(reqBody, apiNodeUrl)
	require.NoError(t, err)

	time.Sleep(20 * time.Second)

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

func sendMessage(body []byte, apiurl string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/message", apiurl), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}

func signMesssage(data []byte, projectID uint64, key *ecdsa.PrivateKey) ([]byte, error) {
	req := &api.HandleMessageReq{
		ProjectID:      projectID,
		ProjectVersion: "v1.0.0",
		Data:           string(data),
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

	fmt.Printf("Signature: %x", sig)
	req.Signature = hexutil.Encode(sig)

	return json.Marshal(req)
}
