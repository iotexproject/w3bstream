package ipfs_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/util/ipfs"
)

func TestIPFS(t *testing.T) {
	r := require.New(t)
	s := ipfs.NewIPFS("ipfs.mainnet.iotex.io")

	content := []byte("test data")
	f, err := os.CreateTemp("", "test")
	r.Nil(err)
	r.NotNil(f)
	defer os.RemoveAll(f.Name())

	cid, err := s.AddFile(f.Name())
	r.NoError(err)
	r.NotEqual(cid, "")
	t.Log(cid)

	_, err = s.AddFile("not_exists")
	r.Error(err)

	cid, err = s.AddContent(content)
	r.NoError(err)
	r.NotEqual(cid, "")
	t.Log(cid)

	content2, err := s.Cat(cid)
	r.NoError(err)
	r.True(bytes.Equal(content2, content))

	_, err = s.Cat("notexists")
	r.Error(err)

	content, err = s.Cat("QmSiDPZH2xCpuC9g2RLSrVYzfWRxLX4Mgu1cNPUAkZe696")
	r.NoError(err)

	h := sha256.New()
	_, err = h.Write(content)
	r.NoError(err)

	sha256sum := h.Sum(nil)
	hashv, err := hex.DecodeString("a543eb8a3e8e1551b0041fb4b7964dba8d07af3cc82490827f8b90d005dd842d")
	r.NoError(err)
	r.True(bytes.Equal(sha256sum, hashv))
}
