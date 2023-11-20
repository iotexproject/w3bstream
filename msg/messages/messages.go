package messages

import (
	"container/list"
	"sync"

	"github.com/machinefi/sprout/msg"
)

type MessageCache struct {
	mtx     *sync.Mutex
	records map[string]*MessageContext
	list    *list.List
	limit   int
}

func (ms *MessageCache) Add(m *msg.Msg) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	c := newMessageContext(m)
	ms.records[m.ID] = c
	ms.list.PushBack(c)
	if ms.list.Len() > ms.limit {
		ms.cleanup()
	}
}

func (ms *MessageCache) cleanup() {
	for ms.list.Len() > ms.limit {
		elem := ms.list.Front()
		ms.list.Remove(elem)
		m := elem.Value.(*MessageContext)
		delete(ms.records, m.ID)
	}
}

func (ms *MessageCache) Query(id string) (*MessageContext, bool) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	c, ok := ms.records[id]
	return c, ok
}

func (ms *MessageCache) OnSubmitProving(id string) {
	c, ok := ms.Query(id)
	if ok {
		c.OnSubmitProving()
	}
}

func (ms *MessageCache) OnProved(id string, res string) {
	c, ok := ms.Query(id)
	if ok {
		c.OnProved(res)
	}
}

func (ms *MessageCache) OnSubmitToBlockchain(id string) {
	c, ok := ms.Query(id)
	if ok {
		c.OnSubmitToBlockchain()
	}
}

func (ms *MessageCache) OnSucceeded(id string, res string) {
	c, ok := ms.Query(id)
	if ok {
		c.OnSucceeded(res)
	}
}

func (ms *MessageCache) OnFailed(id string, err error) {
	c, ok := ms.Query(id)
	if ok {
		c.OnFailed(err)
	}
}

var defaultMessageCache *MessageCache

func init() {
	defaultMessageCache = &MessageCache{
		mtx:     &sync.Mutex{},
		records: make(map[string]*MessageContext),
		list:    list.New(),
		limit:   1024,
	}
}

func New(m *msg.Msg) {
	defaultMessageCache.Add(m)
}

func Query(id string) (*MessageContext, bool) {
	return defaultMessageCache.Query(id)
}

func OnSubmitProving(id string) {
	defaultMessageCache.OnSubmitProving(id)
}

func OnProved(id string, res string) {
	defaultMessageCache.OnProved(id, res)
}

func OnSubmitToBlockchain(id string) {
	defaultMessageCache.OnSubmitToBlockchain(id)
}

func OnSucceeded(id string, res string) {
	defaultMessageCache.OnSucceeded(id, res)
}

func OnFailed(id string, err error) {
	defaultMessageCache.OnFailed(id, err)
}
