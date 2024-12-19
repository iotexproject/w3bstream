package e2e

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iotexproject/w3bstream/e2e/services"
	"github.com/iotexproject/w3bstream/service/apinode/api"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func signMesssage(data []byte, projectID uint64, key *ecdsa.PrivateKey) ([]byte, error) {
	req := &api.CreateTaskReq{
		Nonce:     uint64(time.Now().Unix()),
		ProjectID: strconv.Itoa(int(projectID)),
		Payload:   data,
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	h := sha256.Sum256(reqJson)
	// TODO: uncomment once project config can be loaded
	// value := gjson.GetBytes(data, "timestamp")
	// buf := new(bytes.Buffer)
	// if err := binary.Write(buf, binary.LittleEndian, value.Uint()); err != nil {
	// 	return nil, errors.New("failed to convert uint64 to bytes array")
	// }
	// d := []byte{}
	// d = append(d, h[:]...)
	// d = append(d, buf.Bytes()...)
	// nh := sha256.Sum256(d)
	sig, err := crypto.Sign(h[:], key)
	if err != nil {
		return nil, err
	}
	sig = sig[:len(sig)-1]

	fmt.Printf("signature: %s, hash: %s\n", hexutil.Encode(sig), hexutil.Encode(h[:]))

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

	_, err = services.WaitForTransactionReceipt(client, signedTx.Hash())
	require.NoError(t, err)

	return nil
}
