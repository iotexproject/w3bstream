package mq

import "errors"

var (
	ErrMQFull  = errors.New("message queue is full")
	ErrMQEmpty = errors.New("message queue is empty")
)
