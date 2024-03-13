package apitypes

import "time"

type ErrRsp struct {
	Error string `json:"error,omitempty"`
}

func NewErrRsp(err error) *ErrRsp {
	return &ErrRsp{Error: err.Error()}
}

type HandleMessageReq struct {
	ProjectID      uint64 `json:"projectID"        binding:"required"`
	ProjectVersion string `json:"projectVersion"   binding:"required"`
	Data           string `json:"data"             binding:"required"`
}

type HandleMessageRsp struct {
	MessageID string `json:"messageID"`
}

type LivenessRsp struct {
	Status string `json:"status"`
}

type StateLog struct {
	State   string    `json:"state"`
	Time    time.Time `json:"time"`
	Comment string    `json:"comment"`
	Result  string    `json:"result"`
}

type QueryTaskStateLogRsp struct {
	TaskID    uint64      `json:"taskID"`
	ProjectID uint64      `json:"projectID"`
	States    []*StateLog `json:"states"`
}

type QueryMessageStateLogRsp struct {
	MessageID string      `json:"messageID"`
	States    []*StateLog `json:"states"`
}

type ENodeConfigRsp struct {
	ProjectContractAddress string `json:"projectContractAddress"`
	OperatorETHAddress     string `json:"OperatorETHAddress,omitempty"`
	OperatorSolanaAddress  string `json:"operatorSolanaAddress,omitempty"`
}
