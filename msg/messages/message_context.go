package messages

import (
	"github.com/machinefi/w3bstream-mainnet/enums"
	"github.com/machinefi/w3bstream-mainnet/msg"
	"time"
)

func newMessageContext(m *msg.Msg) *MessageContext {

	return &MessageContext{
		Msg:        m,
		Status:     enums.MessageStatusReceived,
		ReceivedAt: time.Now(),
	}
}

type MessageContext struct {
	*msg.Msg
	Status               enums.MessageStatus `json:"status"`
	ReceivedAt           time.Time           `json:"receivedAt"`
	SubmitProvingAt      time.Time           `json:"submitProvingAt,omitempty"`
	ProofResult          string              `json:"proofResult,omitempty"`
	SubmitToBlockchainAt time.Time           `json:"SubmitToBlockchainAt,omitempty"`
	TxHash               string              `json:"txHash,omitempty"`
	Succeed              bool                `json:"succeed"`
	ErrorMessage         string              `json:"errorMessage,omitempty"`
}

func (mc *MessageContext) OnSubmitProving() {
	mc.Status = enums.MessageStatusSubmitProving
	mc.SubmitProvingAt = time.Now()
}

func (mc *MessageContext) OnProved(res string) {
	mc.Status = enums.MessageStatusProved
	mc.ProofResult = res
}

func (mc *MessageContext) OnSubmitToBlockchain() {
	mc.Status = enums.MessageStatusSubmitToBlockchain
	mc.SubmitToBlockchainAt = time.Now()
}

func (mc *MessageContext) OnSucceeded(txHash string) {
	mc.Succeed = true
	mc.TxHash = txHash
}

func (mc *MessageContext) OnFailed(err error) {
	mc.ErrorMessage = err.Error()
}
