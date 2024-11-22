package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel                slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint         string     `env:"HTTP_SERVICE_ENDPOINT"`
	TaskProcessingBandwidth int        `env:"TASK_PROCESSING_BANDWIDTH"`
	ClickhouseEndpoint      string     `env:"CLICKHOUSE_ENDPOINT"`
	ClickhousePasswd        string     `env:"CLICKHOUSE_PASSWORD"`
	ClickhouseTLS           bool       `env:"CLICKHOUSE_TLS"`
	ChainEndpoint           string     `env:"CHAIN_ENDPOINT,optional"`
	OperatorPrvKey          string     `env:"OPERATOR_PRIVATE_KEY,optional"`
	LocalDBDir              string     `env:"LOCAL_DB_DIRECTORY,optional"`
	BeginningBlockNumber    uint64     `env:"BEGINNING_BLOCK_NUMBER,optional"`
	ProverContractAddr      string     `env:"PROVER_CONTRACT_ADDRESS,optional"`
	MinterContractAddr      string     `env:"MINTER_CONTRACT_ADDRESS,optional"`
	TaskManagerContractAddr string     `env:"TASK_MANAGER_CONTRACT_ADDRESS,optional"`
	env                     string     `env:"-"`
}

var (
	defaultTestnetConfig = &Config{
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9001",
		TaskProcessingBandwidth: 20,
		ChainEndpoint:           "https://babel-api.testnet.iotex.io",
		ClickhouseTLS:           true,
		OperatorPrvKey:          "33e6ba3e033131026903f34dfa208feb88c284880530cf76280b68d38041c67b",
		ProverContractAddr:      "0xab6836908d15E42D30bdEf14cbFA4ad45dCAF3a3",
		MinterContractAddr:      "0x49C096AE869A3054Db06ffF221b917b41f94CEf3",
		TaskManagerContractAddr: "0xF0714400a4C0C72007A9F910C5E3007614958636",
		LocalDBDir:              "./local_db",
		BeginningBlockNumber:    28685000,
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
