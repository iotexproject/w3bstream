package data

import (
	"encoding/json"
	"os"

	"github.com/machinefi/w3bstream-mainnet/project"
)

// TODO delete this file

func GetTestData(file string) *project.Config {
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
