package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	LogLevel                slog.Level `env:"LOG_LEVEL,optional"`
	ServiceEndpoint         string     `env:"HTTP_SERVICE_ENDPOINT"`
	BootNodeMultiAddr       string     `env:"BOOTNODE_MULTIADDR"`
	IoTeXChainID            int        `env:"IOTEX_CHAINID"`
	ChainEndpoint           string     `env:"CHAIN_ENDPOINT,optional"`
	OperatorPrvKey          string     `env:"OPERATOR_PRIVATE_KEY,optional"`
	LocalDBDir              string     `env:"LOCAL_DB_DIRECTORY,optional"`
	BeginningBlockNumber    uint64     `env:"BEGINNING_BLOCK_NUMBER,optional"`
	ProverContractAddr      string     `env:"PROVER_CONTRACT_ADDRESS,optional"`
	DaoContractAddr         string     `env:"DAO_CONTRACT_ADDRESS,optional"`
	MinterContractAddr      string     `env:"MINTER_CONTRACT_ADDRESS,optional"`
	TaskManagerContractAddr string     `env:"TASK_MANAGER_CONTRACT_ADDRESS,optional"`
	env                     string     `env:"-"`
}

var (
	defaultTestnetConfig = &Config{
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9001",
		BootNodeMultiAddr:       "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		IoTeXChainID:            2,
		ChainEndpoint:           "https://babel-api.testnet.iotex.io",
		ProverContractAddr:      "0xE5abfa428046642C0b67C412dE2D1468AB195BcC",
		DaoContractAddr:         "0x3AbD770Ee7400F4F067d9FeC41a798764f8DCbC9",
		MinterContractAddr:      "0x5f733ffAE7FdE6797d36eE8Aee02a480eD9b8F85",
		TaskManagerContractAddr: "0x1641C60Ea5a8B251843273518Ce3b2C4D869fBA1",
		LocalDBDir:              "./local_db",
		BeginningBlockNumber:    28000000,
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
