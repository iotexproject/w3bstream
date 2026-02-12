package event

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventInterface interface {
	Contract() common.Address
	Topic() common.Hash
	HandleEvent(log types.Log) error
	Channel() <-chan interface{}
}
