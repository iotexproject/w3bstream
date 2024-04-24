package didvc

import (
	"bytes"
	"encoding/json"
	"github.com/machinefi/ioconnect-go/cmd/srv-did-vc/apis"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// VerifyJWTCredential verify client token and retrieve client DID
func VerifyJWTCredential(endpoint string, tok string) (string, error) {
	body, err := json.Marshal(&apis.VerifyTokenReq{Token: tok})
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal request")
	}

	rsp, err := http.Post("http://"+endpoint+"/verify", "application/json", bytes.NewReader(body))
	if err != nil {
		return "", errors.Wrap(err, "failed to request verification")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return "", errors.Errorf("failed to request verification: [status code:%d]", rsp.StatusCode)
	}

	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read verification body")
	}

	ret := &apis.VerifyTokenRsp{}
	if err = json.Unmarshal(content, ret); err != nil {
		return "", errors.Wrap(err, "failed to parse verification result")
	}
	return ret.ClientID, nil
}
