package tasks

import (
	"time"

	"github.com/machinefi/sprout/enums"
	"github.com/machinefi/sprout/message"
)

func newTaskContext(m *message.Message) *TaskContext {
	now := time.Now()
	return &TaskContext{
		Message:    m,
		Status:     enums.ZKProofTaskStatusReceived,
		ReceivedAt: &now,
	}
}

type TaskContext struct {
	*message.Message
	Status               enums.ZKProofTaskStatus `json:"status"`
	ReceivedAt           *time.Time              `json:"receivedAt"`
	SubmitProvingAt      *time.Time              `json:"submitProvingAt,omitempty"`
	ProofResult          string                  `json:"proofResult,omitempty"`
	SubmitToBlockchainAt *time.Time              `json:"SubmitToBlockchainAt,omitempty"`
	TxHash               string                  `json:"txHash,omitempty"`
	Succeed              bool                    `json:"succeed"`
	ErrorMessage         string                  `json:"errorMessage,omitempty"`
}

func (mc *TaskContext) OnSubmitProving() {
	mc.Status = enums.ZKProofTaskStatusSubmitProving
	mc.SubmitProvingAt = new(time.Time)
	*mc.SubmitProvingAt = time.Now()
}

func (mc *TaskContext) OnProved(res string) {
	mc.Status = enums.ZKProofTaskStatusProved
	mc.ProofResult = res
}

func (mc *TaskContext) OnSubmitToBlockchain() {
	mc.Status = enums.ZKProofTaskStatusSubmitToBlockchain
	mc.SubmitToBlockchainAt = new(time.Time)
	*mc.SubmitToBlockchainAt = time.Now()
}

func (mc *TaskContext) OnSucceeded(txHash string) {
	mc.Succeed = true
	mc.TxHash = txHash
}

func (mc *TaskContext) OnFailed(err error) {
	mc.ErrorMessage = err.Error()
}
