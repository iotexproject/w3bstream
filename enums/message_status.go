package enums

type MessageStatus uint8

const (
	MessageStatusReceived MessageStatus = iota
	MessageStatusSubmitProving
	MessageStatusProved
	MessageStatusSubmitToBlockchain
)
