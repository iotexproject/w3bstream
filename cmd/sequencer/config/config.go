package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel             slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint      string     `env:"HTTP_SERVICE_ENDPOINT"`
	DatabaseDSN          string     `env:"DATABASE_DSN"`
	BootNodeMultiAddr    string     `env:"BOOTNODE_MULTIADDR"`
	IoTeXChainID         int        `env:"IOTEX_CHAINID"`
	ChainEndpoint        string     `env:"CHAIN_ENDPOINT,optional"`
	ProjectContractAddr  string     `env:"PROJECT_CONTRACT_ADDRESS,optional"`
	ProverContractAddr   string     `env:"PROVER_CONTRACT_ADDRESS,optional"`
	OperatorPriKey       string     `env:"OPERATOR_PRIVATE_KEY,optional"`
	LocalDBDir           string     `env:"LOCAL_DB_DIRECTORY,optional"`
	BeginningBlockNumber uint64     `env:"BEGINNING_BLOCK_NUMBER,optional"`
	env                  string     `env:"-"`
}

var (
	defaultTestnetConfig = &Config{
		LogLevel:             slog.LevelInfo,
		ServiceEndpoint:      ":9001",
		DatabaseDSN:          "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable",
		BootNodeMultiAddr:    "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		IoTeXChainID:         2,
		ChainEndpoint:        "https://babel-api.testnet.iotex.io",
		ProjectContractAddr:  "0x3168A7BE5ba2d9c3aE6309a66152854142c99B26",
		ProverContractAddr:   "0x39d95173C92aadcD47184f770c4a059D8Be66686",
		LocalDBDir:           "./local_db",
		BeginningBlockNumber: 28000000,
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
