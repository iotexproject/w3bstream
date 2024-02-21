package testutil

import (
	"encoding/json"
	"io"
	"net/http"

	. "github.com/agiledragon/gomonkey/v2"
)

func JsonMarshal(p *Patches, data []byte, err error) *Patches {
	return p.ApplyFunc(
		json.Marshal,
		func(v any) ([]byte, error) {
			return data, err
		},
	)
}

func JsonUnmarshal(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		json.Unmarshal,
		func([]byte, any) error {
			return err
		},
	)
}

func HttpPost(p *Patches, rsp *http.Response, err error) *Patches {
	return p.ApplyFunc(
		http.Post,
		func(string, string, io.Reader) (*http.Response, error) {
			return rsp, err
		},
	)
}

func IoReadAll(p *Patches, data []byte, err error) *Patches {
	return p.ApplyFunc(
		io.ReadAll,
		func(io.Reader) ([]byte, error) {
			return data, err
		},
	)
}
