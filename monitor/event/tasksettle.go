package event

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/iotexproject/w3bstream/smartcontracts/go/taskmanager"
)

type settleTask struct {
	contractAddr common.Address
	contract     *taskmanager.Taskmanager
	contractABI  abi.ABI
	channel      chan interface{}
}

type TaskSettleEvent struct {
	ProjectId   uint64
	TaskId      [32]byte
	TxHash      [32]byte
	BlockNumber uint64
}

func NewSettleTask(address common.Address, client *ethclient.Client) EventInterface {
	ta, err := taskmanager.NewTaskmanager(address, client)
	if err != nil {
		panic(err)
	}

	tokenSOLCashierABI, err := abi.JSON(strings.NewReader(taskmanager.TaskmanagerABI))
	if err != nil {
		panic(err)
	}
	return &settleTask{
		contract:     ta,
		contractAddr: address,
		contractABI:  tokenSOLCashierABI,
		channel:      make(chan interface{}, 1),
	}
}

func (s *settleTask) Contract() common.Address {
	return s.contractAddr
}

func (s *settleTask) Topic() common.Hash {
	return s.contractABI.Events["TaskSettled"].ID
}

func (s *settleTask) Channel() <-chan interface{} {
	return s.channel
}

func (s *settleTask) HandleEvent(log types.Log) error {
	ret, err := s.contract.ParseTaskSettled(log)
	if err != nil {
		return err
	}
	select {
	case s.channel <- TaskSettleEvent{
		ProjectId:   ret.ProjectId.Uint64(),
		TaskId:      ret.TaskId,
		TxHash:      log.TxHash,
		BlockNumber: log.BlockNumber,
	}:
		return nil
	default:
		return errors.New("channel is full")
	}
}
