package output

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/tablelandnetwork/basin-cli/pkg/signing"
	"github.com/tidwall/gjson"

	"github.com/machinefi/sprout/task"
)

type textileDB struct {
	endpoint  string
	secretKey *ecdsa.PrivateKey
}

func (t *textileDB) Output(proverID uint64, task *task.Task, proof []byte) (string, error) {
	slog.Debug("outputing to textileDB", "chain endpoint", t.endpoint)
	encodedData, err := t.packData(proof)
	if err != nil {
		return "", err
	}
	txHash, err := t.write(encodedData)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

func (t *textileDB) packData(proof []byte) ([]byte, error) {
	proof, err := hex.DecodeString(string(proof))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode proof")
	}

	var (
		result string
		values []gjson.Result
	)

	if valueJournal := gjson.GetBytes(proof, "Stark.journal.bytes"); valueJournal.Exists() {
		values = valueJournal.Array()
	} else if valueJournal = gjson.GetBytes(proof, "Snark.journal"); valueJournal.Exists() {
		values = valueJournal.Array()
	} else {
		return nil, errors.New("proof does not contain journal")
	}

	// get result from proof
	for _, value := range values {
		result += fmt.Sprint(value.Int())
	}

	data := map[string]string{
		"result": result,
		"proof":  string(proof),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "marshal pack data error")
	}

	return jsonData, nil
}

func (t *textileDB) write(data []byte) (string, error) {
	signatureBytes, err := signing.NewSigner(t.secretKey).SignBytes(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign data")
	}

	txHash := hex.EncodeToString(signatureBytes)

	url := fmt.Sprintf("%s?timestamp=%s&signature=%s",
		t.endpoint,
		strconv.FormatInt(time.Now().Unix(), 10),
		txHash)
	err = writeTextileEvent(url, data)
	if err != nil {
		return "", err
	}
	return txHash, err
}

// writeTextileEvent writes a file to a vault via the Basin API.
func writeTextileEvent(url string, fileData []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewReader(fileData))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	hn := sha256.New()
	hr := hn.Sum(fileData)
	req.Header.Set("filename", string(hr)[0:7])

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

	slog.Debug("Write event", "response", string(responseBody))
	return nil
}

// TODO: refactor textile with a KV database adapter
func newTextileDBAdapter(conf TextileConfig, secretKey string) (*textileDB, error) {
	if secretKey == "" {
		return nil, errors.New("secret key is empty")
	}
	return &textileDB{
		endpoint:  fmt.Sprintf("https://basin.tableland.xyz/vaults/%s/events", conf.VaultID),
		secretKey: crypto.ToECDSAUnsafe(common.FromHex(secretKey)),
	}, nil
}
