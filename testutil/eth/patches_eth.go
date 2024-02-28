package eth

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EthclientDial(p *gomonkey.Patches, cli *ethclient.Client, err error) *gomonkey.Patches {
	return p.ApplyFunc(
		ethclient.Dial,
		func(string) (*ethclient.Client, error) {
			return cli, err
		},
	)
}
