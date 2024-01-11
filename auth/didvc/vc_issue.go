package didvc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func IssueCredential(endpoint string, r *IssueCredentialReq, jwtFormat bool) (*IssueCredentialJWTRsp, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse body")
	}

	if jwtFormat {
		if r.Options == nil {
			r.Options = new(LinkedDataProofOptions)
		}
		r.Options.ProofFormat = ProofFormatJWT
	}

	url := "http://" + endpoint + "/issue/credentials"

	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to request verifiable credential")
	}
	defer rsp.Body.Close()

	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response")
	}

	ret := new(IssueCredentialJWTRsp)
	if err = json.Unmarshal(content, ret); err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return ret, nil
}
