package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel                         slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint                  string     `env:"HTTP_SERVICE_ENDPOINT"`
	BootNodeMultiAddr                string     `env:"BOOTNODE_MULTIADDR"`
	IoTeXChainID                     int        `env:"IOTEX_CHAINID"`
	DatasourceDSN                    string     `env:"DATASOURCE_DSN"`
	ChainEndpoint                    string     `env:"CHAIN_ENDPOINT,optional"`
	OperatorPrvKey                   string     `env:"OPERATOR_PRIVATE_KEY,optional"`
	LocalDBDir                       string     `env:"LOCAL_DB_DIRECTORY,optional"`
	BeginningBlockNumber             uint64     `env:"BEGINNING_BLOCK_NUMBER,optional"`
	ProverContractAddr               string     `env:"PROVER_CONTRACT_ADDRESS,optional"`
	DaoContractAddr                  string     `env:"DAO_CONTRACT_ADDRESS,optional"`
	MinterContractAddr               string     `env:"MINTER_CONTRACT_ADDRESS,optional"`
	TaskManagerContractAddr          string     `env:"TASK_MANAGER_CONTRACT_ADDRESS,optional"`
	BlockHeaderValidatorContractAddr string     `env:"BLOCK_HEADER_VALIDATOR_CONTRACT_ADDRESS,optional"`
	env                              string     `env:"-"`
}

var (
	defaultTestnetConfig = &Config{
		LogLevel:                         slog.LevelInfo,
		ServiceEndpoint:                  ":9001",
		BootNodeMultiAddr:                "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		DatasourceDSN:                    "postgres://postgres:mysecretpassword@postgres:5432/w3bstream?sslmode=disable",
		IoTeXChainID:                     2,
		ChainEndpoint:                    "https://babel-api.testnet.iotex.io",
		OperatorPrvKey:                   "33e6ba3e033131026903f34dfa208feb88c284880530cf76280b68d38041c67b",
		ProverContractAddr:               "0x92aE72A5f15ee8cF61f950A60a600e17875644b2",
		DaoContractAddr:                  "0xAF2C7967BD575C5a4b9a19333faC8d24744775f0",
		MinterContractAddr:               "0xeC0e0749Ec1434C6B23A7175B1309C8AEBa29da9",
		TaskManagerContractAddr:          "0x7AEF1Ed51c1EF3f9e118e25De5D65Ff9F7E2fd29",
		BlockHeaderValidatorContractAddr: "0xd74263530C4e555CEEb3901d4165D83B4071A5e7",
		LocalDBDir:                       "./local_db",
		BeginningBlockNumber:             28685000,
		env:                              "TESTNET",
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
