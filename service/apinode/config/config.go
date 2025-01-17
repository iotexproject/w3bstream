package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel                     slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint              string     `env:"HTTP_SERVICE_ENDPOINT"`
	SequencerServiceEndpoint     string     `env:"SEQUENCER_SERVICE_ENDPOINT"`
	TaskAggregatorIntervalSecond uint64     `env:"TASK_AGGREGATOR_INTERVAL_SECOND,optional"`
	ProverServiceEndpoint        string     `env:"PROVER_SERVICE_ENDPOINT"`
	DatabaseDSN                  string     `env:"DATABASE_DSN"`
	ChainEndpoint                string     `env:"CHAIN_ENDPOINT,optional"`
	BeginningBlockNumber         uint64     `env:"BEGINNING_BLOCK_NUMBER,optional"`
	ProjectContractAddr          string     `env:"PROJECT_CONTRACT_ADDRESS,optional"`
	TaskManagerContractAddr      string     `env:"TASK_MANAGER_CONTRACT_ADDRESS,optional"`
	IoIDContractAddr             string     `env:"IOID_CONTRACT_ADDRESS,optional"`
	LocalDBDir                   string     `env:"LOCAL_DB_DIRECTORY,optional"`
	env                          string     `env:"-"`
}

var defaultTestnetConfig = &Config{
	LogLevel:                     slog.LevelInfo,
	ServiceEndpoint:              ":9000",
	SequencerServiceEndpoint:     "localhost:9001",
	TaskAggregatorIntervalSecond: 1,
	ProverServiceEndpoint:        "localhost:9002",
	ChainEndpoint:                "https://babel-api.testnet.iotex.io",
	BeginningBlockNumber:         28685000,
	TaskManagerContractAddr:      "0xF0714400a4C0C72007A9F910C5E3007614958636",
	IoIDContractAddr:             "0x45Ce3E6f526e597628c73B731a3e9Af7Fc32f5b7",
	LocalDBDir:                   "./local_db",
	env:                          "TESTNET",
}

var defaultMainnetConfig = &Config{
	LogLevel:                     slog.LevelInfo,
	ServiceEndpoint:              ":9000",
	SequencerServiceEndpoint:     "localhost:9001",
	TaskAggregatorIntervalSecond: 300,
	ProverServiceEndpoint:        "localhost:9002",
	ChainEndpoint:                "https://babel-api.mainnet.iotex.io",
	BeginningBlockNumber:         31780000,
	TaskManagerContractAddr:      "0x0A422759A8c6b22Ae8B9C4364763b614d5c0CD29",
	IoIDContractAddr:             "0x1FCB980eD0287777ab05ADc93012332e11300e54",
	LocalDBDir:                   "./local_db",
	env:                          "MAINNET",
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
	case "MAINNET":
		conf = defaultMainnetConfig
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
