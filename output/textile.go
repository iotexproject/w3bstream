package output

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	signing "github.com/machinefi/sprout/output/textile_tmp"
	"github.com/machinefi/sprout/types"
)

const (
	vaultId = "test_signer_impl.data"
)

type textileDB struct {
	endpoint  string
	secretKey *ecdsa.PrivateKey
}

func (t *textileDB) Output(task *types.Task, proof []byte) (string, error) {
	slog.Debug("outputing to textileDB", "chain endpoint", t.endpoint)
	encodedData := t.packData(task, proof)
	txHash, err := t.write(encodedData)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

func (t *textileDB) packData(task *types.Task, proof []byte) []byte {
	// TODO
	return nil
}

func (t *textileDB) write(data []byte) (string, error) {
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)

	// Sign the data
	signer := signing.NewSigner(t.secretKey)
	signature, err := signer.SignData(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign data")
	}

	url := fmt.Sprintf("%s?timestamp=%s&signature=%s", t.endpoint, timestampStr, signature)
	err = writeEvent(url, data)
	return "", err
}

// writeEvent writes a file to a vault via the Basin API.
func writeEvent(url string, fileData []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewReader(fileData))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	hn := sha256.New()
	hr := hn.Sum(fileData)
	req.Header.Set("filename", string(hr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	slog.Debug("Write response", string(responseBody))
	return nil
}

func NewTextileDBAdapter(secretKey string) (Output, error) {
	if len(secretKey) == 0 {
		return nil, errors.New("secretkey is empty")
	}

	pk := crypto.ToECDSAUnsafe(common.FromHex(secretKey))

	return &textileDB{
		endpoint:  fmt.Sprintf("https://basin.tableland.xyz/vaults/%s/events", vaultId),
		secretKey: pk,
	}, nil
}
