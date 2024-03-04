package didvc

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/pkg/errors"
)

func VerifyJWTCredential(endpoint string, tok string) error {
	reqv := &VerifyJWTCredentialReq{
		VerifiableCredential: tok,
		LinkedDataProofOptions: &LinkedDataProofOptions{
			ProofFormat: ProofFormatJWT,
		},
	}
	reqbody, err := json.Marshal(reqv)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}
	slog.Info(string(reqbody))

	url := "http://" + endpoint + "/verify/credentials"
	slog.Info(url)

	rsp, err := http.Post(url, "application/json", bytes.NewReader(reqbody))
	if err != nil {
		return errors.Wrap(err, "failed to request verification")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return errors.Errorf("failed to request verification: [status code:%d]", rsp.StatusCode)
	}

	rspbody, err := io.ReadAll(rsp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read verification body")
	}

	slog.Info(string(rspbody))

	rspv := &VerifyCredentialRsp{}
	if err = json.Unmarshal(rspbody, rspv); err != nil {
		return errors.Wrap(err, "failed to parse verification result")
	}
	slog.Info("verification check", "checks", rspv.Checks, "errors", rspv.Errors, "warns", rspv.Warnings)
	return nil
}
