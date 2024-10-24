package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel                slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint         string     `env:"HTTP_SERVICE_ENDPOINT"`
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
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9002",
		VMEndpoints:             `{"1":"localhost:4001","2":"localhost:4002","3":"zkwasm:4001","4":"wasm:4001"}`,
		ChainEndpoint:           "https://babel-api.testnet.iotex.io",
		DatasourceDSN:           "postgres://postgres:mysecretpassword@postgres:5432/w3bstream?sslmode=disable",
		ProjectContractAddr:     "0x2e45132c8fFeBa7490d57A6118Bd060E55161564",
		RouterContractAddr:      "0xBAB5D88AbECd06c3969fa3CE2597DDD31d13e3C3",
		TaskManagerContractAddr: "0x7AEF1Ed51c1EF3f9e118e25De5D65Ff9F7E2fd29",
		ProverOperatorPrvKey:    "a5f4e99aa80342d5451e8f8fd0dc357ccddb70d3827428fb1fc366f70833497f",
		BeginningBlockNumber:    28685000,
		LocalDBDir:              "./local_db",
		env:                     "TESTNET",
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
