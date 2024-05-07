package didvc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/machinefi/ioconnect-go/cmd/srv-did-vc/apis"
	"github.com/pkg/errors"
)

func IssueCredential(endpoint string, clientID string) (*apis.IssueTokenRsp, error) {
	body, err := json.Marshal(&apis.IssueTokenReq{ClientID: clientID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse body")
	}

	url := "http://" + endpoint + "/issue"

	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to request verifiable credential")
	}
	defer rsp.Body.Close()

	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response")
	}

	ret := new(apis.IssueTokenRsp)
	if err = json.Unmarshal(content, ret); err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return ret, nil
}
