package filefetcher

import (
	"bytes"
	"crypto/sha256"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/iotexproject/w3bstream/util/ipfs"
	"github.com/pkg/errors"
)

type Filedescriptor struct {
	Uri  string
	Hash [32]byte
}

func (fd *Filedescriptor) FetchFile(skipHash ...bool) ([]byte, error) {
	u, err := url.Parse(fd.Uri)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse project file uri %s", fd.Uri)
	}

	var data []byte
	switch u.Scheme {
	case "http", "https":
		resp, _err := http.Get(fd.Uri)
		if _err != nil {
			return nil, errors.Wrapf(_err, "failed to fetch project file, uri %s", fd.Uri)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
	case "ipfs":
		// ipfs url: ipfs://${endpoint}/${cid}
		sh := ipfs.NewIPFS(u.Host)
		cid := strings.Split(strings.Trim(u.Path, "/"), "/")
		data, err = sh.Cat(cid[0])
	default:
		return nil, errors.Errorf("invalid project file uri %s", fd.Uri)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read project file, uri %s", fd.Uri)
	}

	if len(skipHash) == 0 || !skipHash[0] {
		h := sha256.New()
		if _, err := h.Write(data); err != nil {
			return nil, errors.Wrap(err, "failed to generate project file hash")
		}
		if !bytes.Equal(h.Sum(nil), fd.Hash[:]) {
			return nil, errors.New("failed to validate project file hash")
		}
	}

	return data, nil
}
