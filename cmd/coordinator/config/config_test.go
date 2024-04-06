package config_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/cmd/coordinator/config"
)

func TestConfig_Init(t *testing.T) {
	r := require.New(t)

	t.Run("UseEnvConfig", func(t *testing.T) {
		os.Clearenv()
		expected := config.Config{
			ServiceEndpoint:           ":1999",
			ChainEndpoint:             "http://iotex.chainendpoint.io",
			DatabaseDSN:               "postgres://username:password@host:port/database?ext=1",
			BootNodeMultiAddr:         "/dns4/a.b.com/tcp/1000/ipfs/123123123",
			IoTeXChainID:              100,
			ProjectContractAddress:    "0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26",
			IPFSEndpoint:              "a.b.com",
			DIDAuthServerEndpoint:     "didkit.com:10001",
			OperatorPrivateKey:        "",
			OperatorPrivateKeyED25519: "",
			ProjectFileDirectory:      "/path/to/project/configs",
		}

		_ = os.Setenv("HTTP_SERVICE_ENDPOINT", expected.ServiceEndpoint)
		_ = os.Setenv("CHAIN_ENDPOINT", expected.ChainEndpoint)
		_ = os.Setenv("DATABASE_DSN", expected.DatabaseDSN)
		_ = os.Setenv("BOOTNODE_MULTIADDR", expected.BootNodeMultiAddr)
		_ = os.Setenv("IOTEX_CHAINID", strconv.Itoa(expected.IoTeXChainID))
		_ = os.Setenv("PROJECT_CONTRACT_ADDRESS", expected.ProjectContractAddress)
		_ = os.Setenv("IPFS_ENDPOINT", expected.IPFSEndpoint)
		_ = os.Setenv("DIDAUTH_SERVER_ENDPOINT", expected.DIDAuthServerEndpoint)
		// missing some env
		// _ = os.Setenv("OPERATOR_PRIVATE_KEY", expected.OperatorPrivateKey)
		// _ = os.Setenv("OPERATOR_PRIVATE_KEY_ED25519", expected.OperatorPrivateKeyED25519)
		_ = os.Setenv("PROJECT_FILE_DIRECTORY", expected.ProjectFileDirectory)

		c := &config.Config{}
		r.Nil(c.Init())
		r.Equal(*c, expected)
	})

	t.Run("CatchPanicCausedByEmptyRequiredEnvVar", func(t *testing.T) {
		os.Clearenv()

		c := &config.Config{}
		defer func() {
			r.NotNil(recover())
		}()
		_ = c.Init()
	})
}
