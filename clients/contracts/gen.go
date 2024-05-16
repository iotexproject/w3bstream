package contracts

//go:generate abigen --abi abis/ioIDRegistry.json --pkg contracts --type IoIDRegistry -out ./ioid_registry.go
//go:generate abigen --abi abis/ProjectDevice.json --pkg contracts --type ProjectDevice -out ./project_device.go
