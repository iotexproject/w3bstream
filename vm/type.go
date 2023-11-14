package vm

import "github.com/machinefi/w3bstream-mainnet/enums"

type Type string

const (
	Risc0 Type = "risc0"
	Halo2 Type = "halo2"
)

var vmEndpointConfigEnvKeyMap = map[string]Type{
	enums.EnvKeyRisc0ServerEndpoint: Risc0,
	enums.EnvKeyHalo2ServerEndpoint: Halo2,
}
