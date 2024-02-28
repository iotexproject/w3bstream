package eth

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinefi/sprout/project/contracts"
)

func EthclientDial(p *gomonkey.Patches, cli *ethclient.Client, err error) *gomonkey.Patches {
	return p.ApplyFunc(
		ethclient.Dial,
		func(string) (*ethclient.Client, error) {
			return cli, err
		},
	)
}

func ProjectRegistrarContract(p *gomonkey.Patches, instance *contracts.Contracts, err error) *gomonkey.Patches {
	return p.ApplyFunc(
		contracts.NewContracts,
		func(common.Address, bind.ContractBackend) (*contracts.Contracts, error) {
			return instance, err
		},
	)
}
