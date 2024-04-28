package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/project"
)

const seqSendUrl = "http://localhost:9000/message"

type projectInstance struct {
	projectID      uint64
	projectVersion string
	vmType         string
	attribute      project.Attribute
}

var messageMap = map[string]string{
	"risc0": "{\"private_input\":\"20\", \"public_input\":\"11,43\", \"receipt_type\":\"Stark\"}",
	"halo2": "{\"private_a\": 3, \"private_b\": 4}",
}

func sendMessage(projects []projectInstance, num int, duration int) {
	ticker := time.NewTicker(time.Duration(duration) * time.Second)

	go func() {
		for range ticker.C {
			for i := 0; i < num; i++ {
				index := rand.Intn(len(projects))
				fmt.Println("Selected project:", projects[index])

				reqbody, err := json.Marshal(&apitypes.HandleMessageReq{
					ProjectID:      projects[index].projectID,
					ProjectVersion: projects[index].projectVersion,
					Data:           messageMap[projects[index].vmType],
				})
				if err != nil {
					slog.Error("failed to marshal request body", "error", err, "project", projects[index])
				}

				_, err = http.Post(seqSendUrl, "application/json", bytes.NewBuffer(reqbody))
				if err != nil {
					slog.Error("failed to send message", "error", err, "project", projects[index])
				}
			}
		}
	}()
}
