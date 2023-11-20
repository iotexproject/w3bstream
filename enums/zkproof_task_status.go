package enums

type ZKProofTaskStatus uint8

const (
	ZKProofTaskStatusReceived ZKProofTaskStatus = iota
	ZKProofTaskStatusSubmitProving
	ZKProofTaskStatusProved
	ZKProofTaskStatusSubmitToBlockchain
)
