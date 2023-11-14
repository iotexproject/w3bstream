package cmd

import (
	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/common"
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
			return errors.Wrap(err, "get flag node-host failed")
		}
		port, err := cmd.Flags().GetString(NodePort)
		if err != nil {
			return errors.Wrap(err, "get flag node-port failed")
		}

		chainEndpoint, err := cmd.Flags().GetString(ChainEndpoint)
		if err != nil {
			return errors.Wrap(err, "get flag chain-endpoint failed")
		}
		contractAddress, err := cmd.Flags().GetString(ContractAddress)
		if err != nil {
			return errors.Wrap(err, "get flag contract-address failed")
		}

		if host != "" && !govalidator.IsHost(host) {
			return errors.New("invalid host format")
		}
		if port != "" && !govalidator.IsPort(port) {
			return errors.New("invalid host port")
		}
		if chainEndpoint != "" && !govalidator.IsURL(chainEndpoint) {
			return errors.New("invalid endpoint format")
		}
		if contractAddress != "" && !common.IsHexAddress(contractAddress) {
			return errors.New("invalid contract address format")
		}

		if host != "" {
			viper.Set(NodeHost, host)
		}
		if port != "" {
			viper.Set(NodePort, port)
		}
		if chainEndpoint != "" {
			viper.Set(ChainEndpoint, chainEndpoint)
		}
		if contractAddress != "" {
			viper.Set(ContractAddress, contractAddress)
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
	setCmd.Flags().StringP(ChainEndpoint, "c", "", "chain endpoint")
	setCmd.Flags().StringP(ContractAddress, "a", "", "contract address")
}
