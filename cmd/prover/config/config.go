package config

import (
	"log/slog"
	"os"

	"github.com/machinefi/sprout/cmd/internal"
)

type Config struct {
	Risc0ServerEndpoint    string `env:"RISC0_SERVER_ENDPOINT"`
	Halo2ServerEndpoint    string `env:"HALO2_SERVER_ENDPOINT"`
	ZKWasmServerEndpoint   string `env:"ZKWASM_SERVER_ENDPOINT"`
	WasmServerEndpoint     string `env:"WASM_SERVER_ENDPOINT"`
	ChainEndpoint          string `env:"CHAIN_ENDPOINT"`
	ProjectContractAddress string `env:"PROJECT_CONTRACT_ADDRESS,optional"`
	DatabaseDSN            string `env:"DATABASE_DSN"`
	BootNodeMultiAddr      string `env:"BOOTNODE_MULTIADDR"`
	ProverContractAddress  string `env:"PROVER_CONTRACT_ADDRESS,optional"`
	ProverPrivateKey       string `env:"PROVER_PRIVATE_KEY,optional"`
	IoTeXChainID           int    `env:"IOTEX_CHAINID"`
	SchedulerEpoch         uint64 `env:"SCHEDULER_EPOCH"`
	IPFSEndpoint           string `env:"IPFS_ENDPOINT"`
	ProjectFileDirectory   string `env:"PROJECT_FILE_DIRECTORY,optional"`
	ProjectCacheDirectory  string `env:"PROJECT_CACHE_DIRECTORY,optional"`
	LogLevel               int    `env:"LOG_LEVEL,optional"`
	SequencerPubKey        string `env:"SEQUENCER_PUBKEY,optional"`
	ProverPubKey           string `env:"PROVER_PUBKEY,optional"` // TODO del, get from contract
	env                    string `env:"-"`
}

var (
	defaultConfig = &Config{
		Risc0ServerEndpoint:    "risc0:4001",
		Halo2ServerEndpoint:    "halo2:4001",
		ZKWasmServerEndpoint:   "zkwasm:4001",
		WasmServerEndpoint:     "wasm:4001",
		ChainEndpoint:          "https://babel-api.testnet.iotex.io",
		ProjectContractAddress: "0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26",
		DatabaseDSN:            "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable",
		BootNodeMultiAddr:      "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		ProverContractAddress:  "0xB2b3f3c8BB00493c6f12232C2cb3e20A65698939",
		ProverPrivateKey:       "did:key:z6MkmF1AgufHf8ASaxDcCR8iSZjEsEbJMp7LkqyEHw6123",
		IoTeXChainID:           2,
		SchedulerEpoch:         720,
		IPFSEndpoint:           "ipfs.mainnet.iotex.io",
		LogLevel:               int(slog.LevelDebug),
		SequencerPubKey:        "",
		ProverPubKey:           "",
	}

	defaultDebugConfig = &Config{
		Risc0ServerEndpoint:    "localhost:4001",
		Halo2ServerEndpoint:    "localhost:4002",
		ZKWasmServerEndpoint:   "localhost:4003",
		WasmServerEndpoint:     "localhost:4004",
		ChainEndpoint:          "https://babel-api.testnet.iotex.io",
		ProjectContractAddress: "0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26",
		DatabaseDSN:            "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable",
		BootNodeMultiAddr:      "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		ProverContractAddress:  "",
		ProverPrivateKey:       "",
		IoTeXChainID:           2,
		SchedulerEpoch:         720,
		IPFSEndpoint:           "ipfs.mainnet.iotex.io",
		ProjectCacheDirectory:  "./project_cache",
		LogLevel:               int(slog.LevelDebug),
		SequencerPubKey:        "",
		ProverPubKey:           "",
	}

	defaultTestConfig = &Config{
		Risc0ServerEndpoint:    "localhost:14001",
		Halo2ServerEndpoint:    "localhost:14002",
		ZKWasmServerEndpoint:   "localhost:14003",
		WasmServerEndpoint:     "localhost:14004",
		ChainEndpoint:          "https://babel-api.testnet.iotex.io",
		ProjectContractAddress: "",
		DatabaseDSN:            "postgres://test_user:test_passwd@localhost:15432/test?sslmode=disable",
		BootNodeMultiAddr:      "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o",
		ProverContractAddress:  "",
		ProverPrivateKey:       "",
		IoTeXChainID:           2,
		SchedulerEpoch:         720,
		IPFSEndpoint:           "ipfs.mainnet.iotex.io",
		ProjectFileDirectory:   "./testdata",
		LogLevel:               int(slog.LevelDebug),
		SequencerPubKey:        "",
		ProverPubKey:           "",
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
	if err := conf.Init(); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) Print() {
	internal.Print(c)
}
