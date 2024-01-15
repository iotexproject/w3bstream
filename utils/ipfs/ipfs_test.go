package ipfs_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/utils/ipfs"
)

func DISABLE_TestIPFS(t *testing.T) {
	endpoint := "ipfs.mainnet.iotex.io"
	r := require.New(t)

	content := []byte("test data")

	s := ipfs.NewIPFS(endpoint)
	cid, err := s.AddContent(content)
	r.NoError(err)
	r.NotEqual(cid, "")
	t.Log(cid)

	content2, err := s.Cat(cid)
	r.NoError(err)
	r.True(bytes.Equal(content2, content))

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
