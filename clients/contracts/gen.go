package contracts

//go:generate abigen --abi abis/ioIDRegistry.json --pkg contracts --type IoIDRegistry -out ./ioIDRegistry.go
//go:generate abigen --abi abis/ioID.json --pkg contracts --type IoID -out ./ioID.go
