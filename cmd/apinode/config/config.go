package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel          slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint   string     `env:"HTTP_SERVICE_ENDPOINT"`
	AggregationAmount int        `env:"AGGREGATION_AMOUNT,optional"`
	DatabaseDSN       string     `env:"DATABASE_DSN"`
	PrvKey            string     `env:"PRIVATE_KEY,optional"`
	BootNodeMultiAddr string     `env:"BOOTNODE_MULTIADDR"`
	IoTeXChainID      int        `env:"IOTEX_CHAINID"`
	ChainEndpoint     string     `env:"CHAIN_ENDPOINT,optional"`
	LocalDBDir        string     `env:"LOCAL_DB_DIRECTORY,optional"`
	env               string     `env:"-"`
}

var defaultTestnetConfig = &Config{
	LogLevel:          slog.LevelInfo,
	ServiceEndpoint:   ":9000",
	AggregationAmount: 1,
	DatabaseDSN:       "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable",
	PrvKey:            "dbfe03b0406549232b8dccc04be8224fcc0afa300a33d4f335dcfdfead861c85",
	BootNodeMultiAddr: "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
	IoTeXChainID:      2,
	ChainEndpoint:     "https://babel-api.testnet.iotex.io",
	LocalDBDir:        "./local_db",
	env:               "TESTNET",
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
