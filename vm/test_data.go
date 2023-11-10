package vm

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/machinefi/w3bstream-mainnet/project"
)

// TODO delete this file

type testData struct {
	Project  string `json:"project"`
	MD5      string `json:"md5"`
	Content  []byte `json:"content"`
	ExpParam string `json:"expParam"`
}

func getTestData(file string) *project.Config {
	if !strings.HasPrefix(file, "test/data/create.json") {
		file = path.Join("/data", file)
	}
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var payload project.Config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		panic(err)
	}
	return &payload
}
