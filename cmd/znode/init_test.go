package main

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func Test_bindEnvConfig(t *testing.T) {
	r := require.New(t)

	err := os.Unsetenv(IPFSEndpoint)
	initConfig()
	r.NoError(err)
	r.Equal(viper.Get(IPFSEndpoint), "ipfs.mainnet.iotex.io")

	err = os.Setenv(IPFSEndpoint, "any")
	initConfig()
	r.NoError(err)
	r.Equal(viper.Get(IPFSEndpoint), "any")
}
