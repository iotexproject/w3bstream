package didvc

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/pkg/errors"
)

// VerifyJWTCredential verify client token and retrieve client DID
func VerifyJWTCredential(endpoint string, tok string) (string, error) {
	reqv := &VerifyJWTCredentialReq{
		VerifiableCredential: tok,
		LinkedDataProofOptions: &LinkedDataProofOptions{
			ProofFormat: ProofFormatJWT,
		},
	}
	reqbody, err := json.Marshal(reqv)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal request")
	}
	slog.Info(string(reqbody))

	// TODO @fangjian here will call new did vc service to validate token, and retrieve client DID
	url := "http://" + endpoint + "/verify/credentials"
	slog.Info(url)

	rsp, err := http.Post(url, "application/json", bytes.NewReader(reqbody))
	if err != nil {
		return "", errors.Wrap(err, "failed to request verification")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return "", errors.Errorf("failed to request verification: [status code:%d]", rsp.StatusCode)
	}

	rspbody, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read verification body")
	}

	slog.Info(string(rspbody))

	rspv := &VerifyCredentialRsp{}
	if err = json.Unmarshal(rspbody, rspv); err != nil {
		return "", errors.Wrap(err, "failed to parse verification result")
	}
	slog.Info("verification check", "checks", rspv.Checks, "errors", rspv.Errors, "warns", rspv.Warnings)
	// TODO @fangjian return the real client DID
	return "did:key:z6MkeeChrUs1EoKkNNzoy9FwJJb9gNQ92UT8kcXZHMbwj67B", nil
}
