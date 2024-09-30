package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel                slog.Level `env:"LOG_LEVEL,optional"`
	VMEndpoints             string     `env:"VM_ENDPOINTS"`
	DatasourceDSN           string     `env:"DATASOURCE_DSN"`
	ChainEndpoint           string     `env:"CHAIN_ENDPOINT,optional"`
	ProjectContractAddr     string     `env:"PROJECT_CONTRACT_ADDRESS,optional"`
	RouterContractAddr      string     `env:"ROUTER_CONTRACT_ADDRESS,optional"`
	TaskManagerContractAddr string     `env:"TASK_MANAGER_CONTRACT_ADDRESS,optional"`
	ProverOperatorPrvKey    string     `env:"PROVER_OPERATOR_PRIVATE_KEY,optional"`
	BeginningBlockNumber    uint64     `env:"BEGINNING_BLOCK_NUMBER,optional"`
	LocalDBDir              string     `env:"LOCAL_DB_DIRECTORY,optional"`
	env                     string     `env:"-"`
}

var (
	defaultTestnetConfig = &Config{
		LogLevel:             slog.LevelInfo,
		VMEndpoints:          `{"1":"risc0:4001","2":"halo2:4001","3":"zkwasm:4001","4":"wasm:4001"}`,
		ChainEndpoint:        "https://babel-api.testnet.iotex.io",
		DatasourceDSN:        "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable",
		ProjectContractAddr:  "0x6B544a7603cead52AdfD99AA64B3d798083cc4CC",
		ProverOperatorPrvKey: "a5f4e99aa80342d5451e8f8fd0dc357ccddb70d3827428fb1fc366f70833497f",
		BeginningBlockNumber: 20000000,
		LocalDBDir:           "./local_db",
		env:                  "TESTNET",
	}
)

func (c *Config) init() error {
	if err := env.ParseEnv(c); err != nil {
		return err
	}
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(c.LogLevel)})
	slog.SetDefault(slog.New(h))
	return nil
}

func Get() (*Config, error) {
	var conf *Config
	env := os.Getenv("ENV")
	switch env {
	case "TESTNET":
		conf = defaultTestnetConfig
	default:
		env = "TESTNET"
		conf = defaultTestnetConfig
	}
	conf.env = env
	if err := conf.init(); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) Print() {
	env.Print(c)
}
