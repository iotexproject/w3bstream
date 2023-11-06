package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "wsctl",
	Short: "W3bstream control cmd",
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	confDir := path.Join(home, ".w3bstream")
	if err := os.MkdirAll(confDir, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigFile(path.Join(confDir, "config.yaml"))
	viper.SetDefault(NodeHost, "localhost")
	viper.SetDefault(NodePort, "9000")

	if err := viper.ReadInConfig(); err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			if err := viper.WriteConfig(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
