package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel                slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint         string     `env:"HTTP_SERVICE_ENDPOINT"`
	ProverServiceEndpoint   string     `env:"PROVER_SERVICE_ENDPOINT"`
	DatabaseDSN             string     `env:"DATABASE_DSN"`
	PrvKey                  string     `env:"PRIVATE_KEY,optional"`
	BootNodeMultiAddr       string     `env:"BOOTNODE_MULTIADDR"`
	IoTeXChainID            int        `env:"IOTEX_CHAINID"`
	ChainEndpoint           string     `env:"CHAIN_ENDPOINT,optional"`
	BeginningBlockNumber    uint64     `env:"BEGINNING_BLOCK_NUMBER,optional"`
	TaskManagerContractAddr string     `env:"TASK_MANAGER_CONTRACT_ADDRESS,optional"`
	env                     string     `env:"-"`
}

var defaultTestnetConfig = &Config{
	LogLevel:                slog.LevelInfo,
	ServiceEndpoint:         ":9000",
	ProverServiceEndpoint:   "localhost:9002",
	DatabaseDSN:             "postgres://postgres:mysecretpassword@postgres:5432/w3bstream?sslmode=disable",
	PrvKey:                  "dbfe03b0406549232b8dccc04be8224fcc0afa300a33d4f335dcfdfead861c85",
	BootNodeMultiAddr:       "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
	IoTeXChainID:            2,
	ChainEndpoint:           "https://babel-api.testnet.iotex.io",
	BeginningBlockNumber:    28685000,
	TaskManagerContractAddr: "0xF0714400a4C0C72007A9F910C5E3007614958636",
	env:                     "TESTNET",
}

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
