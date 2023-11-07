package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		projectID, err := cmd.Flags().GetString("project-id")
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

		body := struct {
			ProjectID      string `json:"projectID"`
			ProjectVersion string `json:"projectVersion"`
			Data           string `json:"data"`
		}{
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

		return nil
	},
}

func init() {
	messageCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("project-id", "p", "", "the projectID which the message will send to")
	sendCmd.Flags().StringP("project-version", "v", "", "the projectVersion")
	sendCmd.Flags().StringP("data", "d", "", "the data which will send to project")

	sendCmd.MarkFlagRequired("project-id")
	sendCmd.MarkFlagRequired("project-version")
	sendCmd.MarkFlagRequired("data")
}
