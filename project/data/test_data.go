package data

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/machinefi/w3bstream-mainnet/project"
)

// TODO delete this file

func GetTestData(file string) *project.Config {
	if !strings.HasPrefix(file, "test/data/risc0-project-config.json") {
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
