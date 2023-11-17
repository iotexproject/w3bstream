package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/machinefi/w3bstream-mainnet/cmd/node/apis"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send message to w3bstream node",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID, err := cmd.Flags().GetUint64("project-id")
		if err != nil {
			return errors.Wrap(err, "failed to get flag project-id")
		}
		projectVersion, err := cmd.Flags().GetString("project-version")
		if err != nil {
			return errors.Wrap(err, "failed to get flag project-version")
		}
		data, err := cmd.Flags().GetString("data")
		if err != nil {
			return errors.Wrap(err, "failed to get flag data")
		}

		body := &apis.HandleReq{
			ProjectID:      projectID,
			ProjectVersion: projectVersion,
			Data:           data,
		}
		bodyJson, err := json.Marshal(body)
		if err != nil {
			return errors.Wrap(err, "failed to build call message")
		}

		u := url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%s", viper.Get(NodeHost).(string), viper.Get(NodePort).(string)),
			Path:   "/message",
		}

		resp, err := http.Post(u.String(), "application/json", bytes.NewReader(bodyJson))
		if err != nil {
			return errors.Wrap(err, "call w3bstream node failed")
		}
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "failed to read responded content")
		}
		rspVal := &apis.HandleRsp{}
		if err := json.Unmarshal(content, rspVal); err != nil {
			return errors.Wrap(err, "failed to parse responded content")
		}
		content, _ = json.MarshalIndent(rspVal, "", "  ")
		cmd.Println(string(content))

		return nil
	},
}

func init() {
	messageCmd.AddCommand(sendCmd)

	sendCmd.Flags().Uint64P("project-id", "p", 0, "the projectID which the message will send to")
	sendCmd.Flags().StringP("project-version", "v", "", "the projectVersion")
	sendCmd.Flags().StringP("data", "d", "", "the data which will send to project")

	sendCmd.MarkFlagRequired("project-id")
	sendCmd.MarkFlagRequired("project-version")
	sendCmd.MarkFlagRequired("data")
}
