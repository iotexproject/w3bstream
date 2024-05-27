package contract

import "sync/atomic"

type Counter interface {
	acquire() int32
	release() int32
	RefCount() int32
}

type counter struct {
	v atomic.Int32
}

func (i *counter) acquire() int32 {
	return i.v.Add(1)
}

func (i *counter) release() int32 {
	return i.v.Add(-1)
}

func (i *counter) RefCount() int32 {
	return i.v.Load()
}
