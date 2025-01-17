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
	ProverPrvKey            string     `env:"PROVER_OPERATOR_PRIVATE_KEY,optional"`
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
		ProjectContractAddr:     "0x0abec44FC786e8da12267Db5fdeB4311AD1A0A8A",
		RouterContractAddr:      "0x28E0A99A76a467E7418019cBBbF79E4599C73B5B",
		TaskManagerContractAddr: "0xF0714400a4C0C72007A9F910C5E3007614958636",
		ProverPrvKey:            "a5f4e99aa80342d5451e8f8fd0dc357ccddb70d3827428fb1fc366f70833497f",
		BeginningBlockNumber:    28685000,
		LocalDBDir:              "./local_db",
		env:                     "TESTNET",
	}
	defaultMainnetConfig = &Config{
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9002",
		VMEndpoints:             `{"1":"localhost:4001","2":"localhost:4002","3":"zkwasm:4001","4":"wasm:4001"}`,
		ChainEndpoint:           "https://babel-api.mainnet.iotex.io",
		ProjectContractAddr:     "0x6EF4559f2023C93F78d27E0151deF083638478d2",
		RouterContractAddr:      "0x580D9686A7A188746B9f4a06fb5ec9e14E937fde",
		TaskManagerContractAddr: "0x0A422759A8c6b22Ae8B9C4364763b614d5c0CD29",
		ProverPrvKey:            "a5f4e99aa80342d5451e8f8fd0dc357ccddb70d3827428fb1fc366f70833497f",
		BeginningBlockNumber:    31780000,
		LocalDBDir:              "./local_db",
		env:                     "MAINNET",
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
