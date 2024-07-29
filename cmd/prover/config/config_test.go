package config_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/cmd/internal"
	"github.com/iotexproject/w3bstream/cmd/prover/config"
)

func TestConfig_Init(t *testing.T) {
	r := require.New(t)

	t.Run("UseEnvConfig", func(t *testing.T) {
		os.Clearenv()
		expected := config.Config{
			VMEndpoints:          `{"1":"halo2:4001","2":"risc0:4001","3":"zkwasm:4001","4":"wasm:4001"}`,
			ChainEndpoint:        "http://abc.def.com",
			ProjectContractAddr:  "0x123",
			DatabaseDSN:          "postgres://root@localhost/abc?ext=666",
			BootNodeMultiAddr:    "/dsn4/abc/123",
			ProverContractAddr:   "0x456",
			IoTeXChainID:         5,
			SchedulerEpoch:       720,
			IPFSEndpoint:         "abc.ipfs.net",
			ProverOperatorPriKey: "private key",
			ProjectFileDir:       "/path/to/project/configs",
			LocalDBDir:           "./test",
		}

		_ = os.Setenv("VM_ENDPOINTS", expected.VMEndpoints)
		_ = os.Setenv("CHAIN_ENDPOINT", expected.ChainEndpoint)
		_ = os.Setenv("DATABASE_DSN", expected.DatabaseDSN)
		_ = os.Setenv("BOOTNODE_MULTIADDR", expected.BootNodeMultiAddr)
		_ = os.Setenv("PROVER_CONTRACT_ADDRESS", expected.ProverContractAddr)
		_ = os.Setenv("IOTEX_CHAINID", strconv.Itoa(expected.IoTeXChainID))
		_ = os.Setenv("SCHEDULER_EPOCH", strconv.FormatUint(expected.SchedulerEpoch, 10))
		_ = os.Setenv("PROJECT_CONTRACT_ADDRESS", expected.ProjectContractAddr)
		_ = os.Setenv("IPFS_ENDPOINT", expected.IPFSEndpoint)
		_ = os.Setenv("PROVER_OPERATOR_PRIVATE_KEY", expected.ProverOperatorPriKey)
		_ = os.Setenv("PROJECT_FILE_DIRECTORY", expected.ProjectFileDir)
		_ = os.Setenv("LOCAL_DB_DIRECTORY", expected.LocalDBDir)

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

	t.Run("FailedToParseEnv", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(internal.ParseEnv, errors.New(t.Name()))

		c := &config.Config{}
		err := c.Init()
		r.ErrorContains(err, t.Name())
	})
}

func TestGet(t *testing.T) {
	r := require.New(t)

	t.Run("GetDefaultTestConfig", func(t *testing.T) {
		_ = os.Setenv("PROVER_ENV", "INTEGRATION_TEST")

		conf, err := config.Get()
		r.NoError(err)
		r.True(strings.Contains(conf.VMEndpoints, "localhost:14001"))
	})

	t.Run("GetDefaultDebugConfig", func(t *testing.T) {
		_ = os.Setenv("PROVER_ENV", "LOCAL_DEBUG")

		conf, err := config.Get()
		r.NoError(err)
		r.True(strings.Contains(conf.VMEndpoints, "localhost:4001"))
	})

	t.Run("GetDefaultConfig", func(t *testing.T) {
		_ = os.Setenv("PROVER_ENV", "PROD")

		conf, err := config.Get()
		r.NoError(err)
		r.True(strings.Contains(conf.VMEndpoints, "risc0:4001"))
	})
}
