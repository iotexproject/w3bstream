package config

import (
	"log/slog"
	"os"

	"github.com/iotexproject/w3bstream/util/env"
)

type Config struct {
	VMEndpoints             string `env:"VM_ENDPOINTS"`
	DatabaseDSN             string `env:"DATABASE_DSN"`
	BootNodeMultiAddr       string `env:"BOOTNODE_MULTIADDR"`
	ChainEndpoint           string `env:"CHAIN_ENDPOINT,optional"`
	ProjectContractAddr     string `env:"PROJECT_CONTRACT_ADDRESS,optional"`
	ProverContractAddr      string `env:"PROVER_CONTRACT_ADDRESS,optional"`
	ProverOperatorPriKey    string `env:"PROVER_OPERATOR_PRIVATE_KEY,optional"`
	IoTeXChainID            int    `env:"IOTEX_CHAINID"`
	SchedulerEpoch          uint64 `env:"SCHEDULER_EPOCH,optional"`
	BeginningBlockNumber    uint64 `env:"BEGINNING_BLOCK_NUMBER,optional"`
	ProjectFileDir          string `env:"PROJECT_FILE_DIRECTORY,optional"`
	LocalDBDir              string `env:"LOCAL_DB_DIRECTORY,optional"`
	LogLevel                int    `env:"LOG_LEVEL,optional"`
	DefaultDatasourcePubKey string `env:"DEFAULT_DATASOURCE_PUBLIC_KEY,optional"`
	env                     string `env:"-"`
}

var (
	defaultConfig = &Config{
		VMEndpoints:             `{"1":"risc0:4001","2":"halo2:4001","3":"zkwasm:4001","4":"wasm:4001"}`,
		ChainEndpoint:           "https://babel-api.testnet.iotex.io",
		DatabaseDSN:             "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable",
		BootNodeMultiAddr:       "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		ProjectContractAddr:     "0xCBb7a80983Fd3405972F700101A82DB6304C6547",
		ProverContractAddr:      "0x6B544a7603cead52AdfD99AA64B3d798083cc4CC",
		ProverOperatorPriKey:    "a5f4e99aa80342d5451e8f8fd0dc357ccddb70d3827428fb1fc366f70833497f",
		IoTeXChainID:            2,
		SchedulerEpoch:          720,
		BeginningBlockNumber:    20000000,
		LocalDBDir:              "./local_db",
		LogLevel:                int(slog.LevelDebug),
		DefaultDatasourcePubKey: "0x04df6acbc5b355aabfb2145b36b20b7942c831c245c423a20b189fab4cf3a3dba3d564080841f2eb4890c118ca5e0b80b25f81269621c5e28273a962996c109afa",
	}

	defaultDebugConfig = &Config{
		VMEndpoints:             `{"1":"localhost:4001","2":"localhost:4002","3":"localhost:4003","4":"localhost:4004"}`,
		DatabaseDSN:             "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable",
		BootNodeMultiAddr:       "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		ProverOperatorPriKey:    "a5f4e99aa80342d5451e8f8fd0dc357ccddb70d3827428fb1fc366f70833497f",
		IoTeXChainID:            2,
		LogLevel:                int(slog.LevelDebug),
		DefaultDatasourcePubKey: "0x04df6acbc5b355aabfb2145b36b20b7942c831c245c423a20b189fab4cf3a3dba3d564080841f2eb4890c118ca5e0b80b25f81269621c5e28273a962996c109afa",
	}

	defaultTestConfig = &Config{
		VMEndpoints:             `{"1":"localhost:14001","2":"localhost:14002","3":"localhost:14003","4":"localhost:14004"}`,
		ChainEndpoint:           "https://babel-api.testnet.iotex.io",
		ProjectContractAddr:     "",
		DatabaseDSN:             "postgres://test_user:test_passwd@localhost:15432/test?sslmode=disable",
		BootNodeMultiAddr:       "/dns4/localhost/tcp/18000/p2p/12D3KooWJkfxZL1dx74yM1afWof6ka4uW5jMsoGasCSBwGyCUJML",
		ProverContractAddr:      "",
		ProverOperatorPriKey:    "a5f4e99aa80342d5451e8f8fd0dc357ccddb70d3827428fb1fc366f70833497f",
		IoTeXChainID:            2,
		SchedulerEpoch:          20,
		ProjectFileDir:          "./testdata",
		LogLevel:                int(slog.LevelDebug),
		DefaultDatasourcePubKey: "0x04df6acbc5b355aabfb2145b36b20b7942c831c245c423a20b189fab4cf3a3dba3d564080841f2eb4890c118ca5e0b80b25f81269621c5e28273a962996c109afa",
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
	env := os.Getenv("PROVER_ENV")
	switch env {
	case "INTEGRATION_TEST":
		conf = defaultTestConfig
	case "LOCAL_DEBUG":
		conf = defaultDebugConfig
	default:
		env = "PROD"
		conf = defaultConfig
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
