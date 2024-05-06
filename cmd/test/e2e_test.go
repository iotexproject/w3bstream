package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/machinefi/sprout/apitypes"
	coordinatorconfig "github.com/machinefi/sprout/cmd/coordinator/config"
	"github.com/machinefi/sprout/task"
)

func TestHttpApi(t *testing.T) {
	r := require.New(t)

	coordinatorConf, err := coordinatorconfig.Get()
	r.NoError(err)
	conf := seqConf(coordinatorConf.ServiceEndpoint, coordinatorConf.DIDAuthServerEndpoint)

	t.Run("BadRequest", func(t *testing.T) {
		reqbody, err := json.Marshal(map[string]interface{}{
			"ProjectID":      10000,
			"ProjectVersion": "0.1",
		})
		r.NoError(err)

		resp, err := http.Post(fmt.Sprintf("http://localhost%s/message", conf.endpoint), "application/json", bytes.NewBuffer(reqbody))
		r.NoError(err)
		r.Equal(400, resp.StatusCode)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		r.NoError(err)
		r.Equal(`{"error":"Key: 'HandleMessageReq.Data' Error:Field validation for 'Data' failed on the 'required' tag"}`, string(body))
	})

	t.Run("Risc0Project", func(t *testing.T) {
		var messageID string
		t.Run("SendMessage", func(t *testing.T) {
			reqbody, err := json.Marshal(&apitypes.HandleMessageReq{
				ProjectID:      10000,
				ProjectVersion: "0.1",
				Data:           "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Stark\"}",
			})
			r.NoError(err)

			resp, err := http.Post(fmt.Sprintf("http://localhost%s/message", conf.endpoint), "application/json", bytes.NewBuffer(reqbody))
			r.NoError(err)
			r.Equal(200, resp.StatusCode)
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			r.NoError(err)

			value := gjson.Get(string(body), "messageID")
			r.Equal(true, value.Exists())
			messageID = value.String()
		})

		t.Run("QueryMessage", func(t *testing.T) {
			var finalState string
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for range ticker.C {
				resp, err := http.Get(fmt.Sprintf("http://localhost%s/message/%s", conf.endpoint, messageID))
				r.NoError(err)
				r.Equal(200, resp.StatusCode)
				body, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				r.NoError(err)

				isBreak := false
				states := gjson.Get(string(body), "states")
				states.ForEach(func(_, v gjson.Result) bool {
					state := v.Get("state").String()
					switch state {
					case task.StateOutputted.String():
						finalState = state
						isBreak = true
						return false
					case task.StateFailed.String():
						finalState = state
						isBreak = true
						return false
					default:
						return true
					}
				})
				if isBreak {
					break
				}
			}
			r.Equal(task.StateOutputted.String(), finalState)
		})
	})

	t.Run("Halo2Project", func(t *testing.T) {
		var messageID string
		t.Run("SendMessage", func(t *testing.T) {
			reqbody, err := json.Marshal(&apitypes.HandleMessageReq{
				ProjectID:      10001,
				ProjectVersion: "0.1",
				Data:           "{\"private_a\": 3, \"private_b\": 5}",
			})
			r.NoError(err)

			resp, err := http.Post(fmt.Sprintf("http://localhost%s/message", conf.endpoint), "application/json", bytes.NewBuffer(reqbody))
			r.NoError(err)
			r.Equal(200, resp.StatusCode)
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			r.NoError(err)

			value := gjson.Get(string(body), "messageID")
			r.Equal(true, value.Exists())
			messageID = value.String()
		})

		t.Run("QueryMessage", func(t *testing.T) {
			var finalState string
			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()

			for range ticker.C {
				resp, err := http.Get(fmt.Sprintf("http://localhost%s/message/%s", conf.endpoint, messageID))
				r.NoError(err)
				r.Equal(200, resp.StatusCode)
				body, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				r.NoError(err)

				isBreak := false
				states := gjson.Get(string(body), "states")
				states.ForEach(func(_, v gjson.Result) bool {
					state := v.Get("state").String()
					switch state {
					case task.StateOutputted.String():
						finalState = state
						isBreak = true
						return false
					case task.StateFailed.String():
						finalState = state
						isBreak = true
						return false
					default:
						return true
					}
				})
				if isBreak {
					break
				}
			}
			r.Equal(task.StateOutputted.String(), finalState)
		})
	})
}
