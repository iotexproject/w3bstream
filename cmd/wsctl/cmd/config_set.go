package cmd

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "update config",
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString(NodeHost)
		if err != nil {
			return errors.Wrap(err, "get flag nodeHost failed")
		}
		port, err := cmd.Flags().GetString(NodePort)
		if err != nil {
			return errors.Wrap(err, "get flag nodePort failed")
		}

		if host != "" && !govalidator.IsHost(host) {
			return errors.New("invalid host format")
		}
		if port != "" && !govalidator.IsPort(port) {
			return errors.New("invalid host port")
		}

		if host != "" {
			viper.Set(NodeHost, host)
		}
		if port != "" {
			viper.Set(NodePort, port)
		}

		if err := viper.WriteConfig(); err != nil {
			return errors.Wrap(err, "failed to write config file")
		}

		return nil
	},
}

func init() {
	configCmd.AddCommand(setCmd)

	setCmd.Flags().StringP(NodeHost, "n", "", "node host")
	setCmd.Flags().StringP(NodePort, "p", "", "node port")
}
