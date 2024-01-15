package ipfs

import (
	"bytes"
	"io"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
)

func NewIPFS(endpoint string) *IPFS {
	sh := shell.NewShell(endpoint)
	return &IPFS{
		endpoint: endpoint,
		sh:       sh,
	}
}

type IPFS struct {
	endpoint string
	sh       *shell.Shell
}

func (s *IPFS) AddFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read file: %s", path)
	}
	return s.AddContent(content)
}

func (s *IPFS) AddContent(content []byte) (string, error) {
	reader := bytes.NewReader(content)

	return s.sh.Add(reader, shell.Pin(true))
}

func (s *IPFS) Cat(cid string) ([]byte, error) {
	reader, err := s.sh.Cat(cid)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read content from ipfs: %s", cid)
	}
	defer reader.Close()

	return io.ReadAll(reader)
}
