package mq

type Msg struct {
	Data []byte `json:"data"`
}

type MQ interface {
	Enqueue(*Msg) error
	Dequeue() (*Msg, error)
}
