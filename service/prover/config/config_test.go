package config

import (
	"os"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/util/env"
)

func TestConfig_Init(t *testing.T) {
	r := require.New(t)

	t.Run("UseEnvConfig", func(t *testing.T) {
		os.Clearenv()
		expected := Config{
			ServiceEndpoint:      "test",
			VMEndpoints:          `{"1":"halo2:4001","2":"risc0:4001","3":"zkwasm:4001","4":"wasm:4001"}`,
			ChainEndpoint:        "http://abc.def.com",
			DatasourceDSN:        "postgres://root@localhost/abc?ext=666",
			ProjectContractAddr:  "0x456",
			ProverOperatorPrvKey: "private key",
			LocalDBDir:           "./test",
		}

		_ = os.Setenv("HTTP_SERVICE_ENDPOINT", expected.ServiceEndpoint)
		_ = os.Setenv("VM_ENDPOINTS", expected.VMEndpoints)
		_ = os.Setenv("CHAIN_ENDPOINT", expected.ChainEndpoint)
		_ = os.Setenv("DATASOURCE_DSN", expected.DatasourceDSN)
		_ = os.Setenv("PROJECT_CONTRACT_ADDRESS", expected.ProjectContractAddr)
		_ = os.Setenv("PROVER_OPERATOR_PRIVATE_KEY", expected.ProverOperatorPrvKey)
		_ = os.Setenv("LOCAL_DB_DIRECTORY", expected.LocalDBDir)

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
