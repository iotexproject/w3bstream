package handler

import (
	"context"
	"github.com/machinefi/w3bstream-mainnet/enums"
	"github.com/machinefi/w3bstream-mainnet/msg/messages"
	"github.com/spf13/viper"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/output/chain/eth"
	"github.com/machinefi/w3bstream-mainnet/project/data"
	"github.com/machinefi/w3bstream-mainnet/test/contract"
	"github.com/machinefi/w3bstream-mainnet/util/mq"
	"github.com/machinefi/w3bstream-mainnet/util/mq/gochan"
	"github.com/machinefi/w3bstream-mainnet/vm"
)

type Handler struct {
	mq                    mq.MQ
	vmHandler             *vm.Handler
	chainEndpoint         string
	operatorPrivateKey    string
	projectConfigFilePath string
}

func New(vmHandler *vm.Handler, chainEndpoint, operatorPrivateKey, projectConfigFilePath string) *Handler {
	q := gochan.New()
	h := &Handler{
		mq:                    q,
		vmHandler:             vmHandler,
		chainEndpoint:         chainEndpoint,
		operatorPrivateKey:    operatorPrivateKey,
		projectConfigFilePath: projectConfigFilePath,
	}
	go q.Watch(h.asyncHandle)
	return h
}

func (r *Handler) Handle(msg *msg.Msg) error {
	slog.Debug("push message into sequencer")
	messages.New(msg)
	return r.mq.Enqueue(msg)
}

func (r *Handler) asyncHandle(m *msg.Msg) {
	slog.Debug("message popped")
	project := data.GetTestData(r.projectConfigFilePath)

	messages.OnSubmitProving(m.ID)
	res, err := r.vmHandler.Handle(m, project.VMType, project.Code, project.CodeExpParam)
	if err != nil {
		slog.Error("proof failed: ", err)
		messages.OnFailed(m.ID, err)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	messages.OnProved(m.ID, string(res))

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		messages.OnFailed(m.ID, err)
		return
	}
	slog.Debug("writing proof to chain")
	messages.OnSubmitToBlockchain(m.ID)
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x190Cc9af23504ac5Dc461376C1e2319bc3B9cD29", data)
	if err != nil {
		slog.Error(err.Error())
		messages.OnFailed(m.ID, err)
		return
	}
	messages.OnSucceeded(m.ID, txHash)
	slog.Debug("transaction hash", "tx_hash", txHash)
}

var DefaultHandler *Handler

func init() {
	DefaultHandler = New(
		vm.DefaultHandler,
		viper.GetString(enums.EnvKeyChainEndpoint),
		viper.GetString(enums.EnvKeyOperatorPrivateKey),
		viper.GetString(enums.EnvKeyProjectConfigPath),
	)
}
