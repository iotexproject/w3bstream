package config

import (
	"log/slog"
	"os"

	"github.com/machinefi/sprout/cmd/internal"
)

type Config struct {
	ServiceEndpoint         string `env:"HTTP_SERVICE_ENDPOINT"`
	DatabaseDSN             string `env:"DATABASE_DSN"`
	DefaultDatasourceURI    string `env:"DEFAULT_DATASOURCE_URI"`
	BootNodeMultiAddr       string `env:"BOOTNODE_MULTIADDR"`
	IoTeXChainID            int    `env:"IOTEX_CHAINID"`
	ChainEndpoint           string `env:"CHAIN_ENDPOINT,optional"`
	ProjectContractAddr     string `env:"PROJECT_CONTRACT_ADDRESS,optional"`
	ProverContractAddr      string `env:"PROVER_CONTRACT_ADDRESS,optional"`
	IPFSEndpoint            string `env:"IPFS_ENDPOINT"`
	OperatorPriKey          string `env:"OPERATOR_PRIVATE_KEY,optional"`
	OperatorPriKeyED25519   string `env:"OPERATOR_PRIVATE_KEY_ED25519,optional"`
	ProjectFileDir          string `env:"PROJECT_FILE_DIRECTORY,optional"`
	ProjectCacheDir         string `env:"PROJECT_CACHE_DIRECTORY,optional"`
	LocalDBDir              string `env:"LOCAL_DB_DIRECTORY,optional"`
	SchedulerEpoch          uint64 `env:"SCHEDULER_EPOCH,optional"`
	BeginningBlockNumber    uint64 `env:"BEGINNING_BLOCK_NUMBER,optional"`
	LogLevel                int    `env:"LOG_LEVEL,optional"`
	DefaultDatasourcePubKey string `env:"DEFAULT_DATASOURCE_PUBLIC_KEY,optional"`
	ContractWhitelist       string `env:"CONTRACT_WHITELIST,optional"`
	env                     string `env:"-"`
}

var (
	// prod default config for coordinator; all config elements come from docker-compose.yaml in root of project
	defaultConfig = &Config{
		ServiceEndpoint:         ":9001",
		DatabaseDSN:             "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable",
		DefaultDatasourceURI:    "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable",
		BootNodeMultiAddr:       "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		IoTeXChainID:            2,
		ChainEndpoint:           "https://babel-api.testnet.iotex.io",
		ProjectContractAddr:     "0xCBb7a80983Fd3405972F700101A82DB6304C6547",
		ProverContractAddr:      "0x6B544a7603cead52AdfD99AA64B3d798083cc4CC",
		IPFSEndpoint:            "ipfs.mainnet.iotex.io",
		DefaultDatasourcePubKey: "0x04df6acbc5b355aabfb2145b36b20b7942c831c245c423a20b189fab4cf3a3dba3d564080841f2eb4890c118ca5e0b80b25f81269621c5e28273a962996c109afa",
		LogLevel:                int(slog.LevelDebug),
		LocalDBDir:              "./local_db",
		SchedulerEpoch:          720,
		BeginningBlockNumber:    20000000,
		ContractWhitelist:       "0x1AA325E5144f763a520867c56FC77cC1411430d0,0xC9D7D9f25b98119DF5b2303ac0Df6b15C982BbF5",
	}
	// local debug default config for coordinator; all config elements come from docker-compose-dev.yaml in root of project
	defaultDebugConfig = &Config{
		ServiceEndpoint:         ":9001",
		DatabaseDSN:             "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable",
		DefaultDatasourceURI:    "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable",
		BootNodeMultiAddr:       "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		IoTeXChainID:            2,
		IPFSEndpoint:            "ipfs.mainnet.iotex.io",
		DefaultDatasourcePubKey: "0x04df6acbc5b355aabfb2145b36b20b7942c831c245c423a20b189fab4cf3a3dba3d564080841f2eb4890c118ca5e0b80b25f81269621c5e28273a962996c109afa",
		LogLevel:                int(slog.LevelDebug),
		ContractWhitelist:       "0x1AA325E5144f763a520867c56FC77cC1411430d0,0xC9D7D9f25b98119DF5b2303ac0Df6b15C982BbF5",
	}
	// integration default config for coordinator; all config elements come from Makefile in `integration_test` entry
	defaultTestConfig = &Config{
		ServiceEndpoint:         ":19001",
		ChainEndpoint:           "https://babel-api.testnet.iotex.io",
		DatabaseDSN:             "postgres://test_user:test_passwd@localhost:15432/test?sslmode=disable",
		DefaultDatasourceURI:    "postgres://test_user:test_passwd@localhost:15432/test?sslmode=disable",
		BootNodeMultiAddr:       "/dns4/localhost/tcp/18000/p2p/12D3KooWJkfxZL1dx74yM1afWof6ka4uW5jMsoGasCSBwGyCUJML",
		IoTeXChainID:            2,
		ProjectContractAddr:     "", //"0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26",
		IPFSEndpoint:            "ipfs.mainnet.iotex.io",
		ProjectFileDir:          "./testdata",
		DefaultDatasourcePubKey: "0x04df6acbc5b355aabfb2145b36b20b7942c831c245c423a20b189fab4cf3a3dba3d564080841f2eb4890c118ca5e0b80b25f81269621c5e28273a962996c109afa",
		LogLevel:                int(slog.LevelDebug),
	}
)

func (c *Config) Init() error {
	if err := internal.ParseEnv(c); err != nil {
		return err
	}
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(c.LogLevel)})
	slog.SetDefault(slog.New(h))
	return nil
}

func (c *Config) Env() string {
	return c.env
}

func Get() (*Config, error) {
	var conf *Config
	env := os.Getenv("COORDINATOR_ENV")
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
	if err := conf.Init(); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) Print() {
	internal.Print(c)
}
