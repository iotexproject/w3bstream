package project

import (
	"bytes"
	"crypto/sha256"
	"log/slog"
	"os"
	"path"
	"strconv"

	"github.com/pkg/errors"
)

type cache struct {
	dir string
}

func (c *cache) getPath(projectID uint64) string {
	return path.Join(c.dir, strconv.FormatUint(projectID, 10))
}

func (c *cache) get(projectID uint64, hash []byte) []byte {
	data, err := os.ReadFile(c.getPath(projectID))
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			slog.Error("failed to read project cache file", "error", err, "project_id", projectID)
		}
		return nil
	}
	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		slog.Error("failed to generate cache project file hash", "error", err)
		return nil
	}
	if !bytes.Equal(h.Sum(nil), hash) {
		slog.Error("failed to validate cache project file hash")
		return nil
	}
	return data
}

func (c *cache) set(projectID uint64, data []byte) {
	if err := os.WriteFile(c.getPath(projectID), data, 0666); err != nil {
		slog.Error("failed to write cache project file", "error", err)
	}
}

func newCache(projectCacheDir string) (*cache, error) {
	if err := os.MkdirAll(projectCacheDir, 0777); err != nil {
		return nil, errors.Wrap(err, "failed to create project cache directory")
	}
	return &cache{
		dir: projectCacheDir,
	}, nil
}
