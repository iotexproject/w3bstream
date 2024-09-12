package config

import (
	"os"
	"strconv"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/iotexproject/w3bstream/util/env"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestConfig_Init(t *testing.T) {
	r := require.New(t)

	t.Run("UseEnvConfig", func(t *testing.T) {
		os.Clearenv()
		expected := Config{
			ServiceEndpoint:       ":1999",
			ChainEndpoint:         "http://iotex.chainendpoint.io",
			DatabaseDSN:           "postgres://username:password@host:port/database?ext=1",
			DefaultDatasourceURI:  "postgres://username:password@host:port/database?ext=1",
			BootNodeMultiAddr:     "/dns4/a.b.com/tcp/1000/ipfs/123123123",
			IoTeXChainID:          100,
			ProjectContractAddr:   "0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26",
			ProverContractAddr:    "0x",
			OperatorPriKey:        "",
			OperatorPriKeyED25519: "",
			ProjectFileDir:        "/path/to/project/configs",
			SchedulerEpoch:        10,
			LocalDBDir:            "./test",
		}

		_ = os.Setenv("HTTP_SERVICE_ENDPOINT", expected.ServiceEndpoint)
		_ = os.Setenv("CHAIN_ENDPOINT", expected.ChainEndpoint)
		_ = os.Setenv("DATABASE_DSN", expected.DatabaseDSN)
		_ = os.Setenv("DEFAULT_DATASOURCE_URI", expected.DefaultDatasourceURI)
		_ = os.Setenv("BOOTNODE_MULTIADDR", expected.BootNodeMultiAddr)
		_ = os.Setenv("IOTEX_CHAINID", strconv.Itoa(expected.IoTeXChainID))
		_ = os.Setenv("PROJECT_CONTRACT_ADDRESS", expected.ProjectContractAddr)
		_ = os.Setenv("PROVER_CONTRACT_ADDRESS", expected.ProverContractAddr)
		_ = os.Setenv("LOCAL_DB_DIRECTORY", expected.LocalDBDir)
		// missing some env
		// _ = os.Setenv("OPERATOR_PRIVATE_KEY", expected.OperatorPrivateKey)
		// _ = os.Setenv("OPERATOR_PRIVATE_KEY_ED25519", expected.OperatorPrivateKeyED25519)
		_ = os.Setenv("PROJECT_FILE_DIRECTORY", expected.ProjectFileDir)
		_ = os.Setenv("SCHEDULER_EPOCH", strconv.FormatUint(expected.SchedulerEpoch, 10))

		c := &Config{}
		r.Nil(c.init())
		r.Equal(*c, expected)
	})

	t.Run("CatchPanicCausedByEmptyRequiredEnvVar", func(t *testing.T) {
		os.Clearenv()

		c := &Config{}
		defer func() {
			r.NotNil(recover())
		}()
		_ = c.init()
	})

	t.Run("FailedToParseEnv", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(env.ParseEnv, errors.New(t.Name()))

		c := &Config{}
		err := c.init()
		r.ErrorContains(err, t.Name())
	})
}

func TestGet(t *testing.T) {
	r := require.New(t)

	t.Run("GetDefaultTestConfig", func(t *testing.T) {
		_ = os.Setenv("COORDINATOR_ENV", "INTEGRATION_TEST")

		conf, err := Get()
		r.NoError(err)
		r.Equal(":19001", conf.ServiceEndpoint)
	})

	t.Run("GetDefaultDebugConfig", func(t *testing.T) {
		_ = os.Setenv("COORDINATOR_ENV", "LOCAL_DEBUG")

		conf, err := Get()
		r.NoError(err)
		r.Equal(":9001", conf.ServiceEndpoint)
	})

	t.Run("GetDefaultConfig", func(t *testing.T) {
		_ = os.Setenv("COORDINATOR_ENV", "PROD")

		conf, err := Get()
		r.NoError(err)
		r.Equal(":9001", conf.ServiceEndpoint)
	})
}