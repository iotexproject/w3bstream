package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/machinefi/w3bstream-mainnet/msg/messages"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

var messageQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "query message by message id from w3bstream node",
	RunE: func(cmd *cobra.Command, args []string) error {
		messageID, err := cmd.Flags().GetString("message-id")
		if err != nil {
			return errors.Wrap(err, "failed to get flag message-id")
		}

		u := url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%s", viper.Get(NodeHost).(string), viper.Get(NodePort).(string)),
			Path:   "/message/" + messageID,
		}

		rsp, err := http.Get(u.String())
		if err != nil {
			return errors.Wrap(err, "call w3bstream node failed")
		}

		switch sc := rsp.StatusCode; sc {
		case http.StatusNotFound:
			return errors.Errorf("the message [%s] is not found or expired", messageID)
		case http.StatusOK:
		default:
			return errors.Errorf("responded status code: %d", sc)
		}

		defer rsp.Body.Close()
		content, err := io.ReadAll(rsp.Body)
		if err != nil {
			return errors.Wrap(err, "read responded body failed")
		}

		rspVal := &messages.MessageContext{}
		if err := json.Unmarshal(content, rspVal); err != nil {
			return errors.Wrap(err, "parse responded body failed")
		}
		content, _ = json.MarshalIndent(rspVal, "", "  ")

		cmd.Print(string(content))
		return nil
	},
}

func init() {
	messageCmd.AddCommand(messageQueryCmd)

	messageQueryCmd.Flags().StringP("message-id", "m", "", "the message id you want to query")
	_ = messageQueryCmd.MarkFlagRequired("message-id")
}
