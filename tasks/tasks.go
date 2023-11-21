package tasks

import (
	"container/list"
	"sync"

	"github.com/machinefi/sprout/message"
)

type Cache struct {
	mtx     *sync.Mutex
	records map[string]*TaskContext
	list    *list.List
	limit   int
}

func (c *Cache) Add(m *message.Message) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	ctx := newTaskContext(m)
	c.records[m.ID] = ctx
	c.list.PushBack(ctx)
	if c.list.Len() > c.limit {
		c.cleanup()
	}
}

func (c *Cache) cleanup() {
	for c.list.Len() > c.limit {
		elem := c.list.Front()
		c.list.Remove(elem)
		m := elem.Value.(*TaskContext)
		delete(c.records, m.ID)
	}
}

func (c *Cache) Query(id string) (*TaskContext, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	t, ok := c.records[id]
	return t, ok
}

func (c *Cache) OnSubmitProving(id string) {
	t, ok := c.Query(id)
	if ok {
		t.OnSubmitProving()
	}
}

func (c *Cache) OnProved(id string, res string) {
	t, ok := c.Query(id)
	if ok {
		t.OnProved(res)
	}
}

func (c *Cache) OnSubmitToBlockchain(id string) {
	t, ok := c.Query(id)
	if ok {
		t.OnSubmitToBlockchain()
	}
}

func (c *Cache) OnSucceeded(id string, res string) {
	t, ok := c.Query(id)
	if ok {
		t.OnSucceeded(res)
	}
}

func (c *Cache) OnFailed(id string, err error) {
	t, ok := c.Query(id)
	if ok {
		t.OnFailed(err)
	}
}

var defaultTasksCache *Cache

func init() {
	defaultTasksCache = &Cache{
		mtx:     &sync.Mutex{},
		records: make(map[string]*TaskContext),
		list:    list.New(),
		limit:   1024,
	}
}

func New(m *message.Message) {
	defaultTasksCache.Add(m)
}

func Query(id string) (*TaskContext, bool) {
	return defaultTasksCache.Query(id)
}

func OnSubmitProving(id string) {
	defaultTasksCache.OnSubmitProving(id)
}

func OnProved(id string, res string) {
	defaultTasksCache.OnProved(id, res)
}

func OnSubmitToBlockchain(id string) {
	defaultTasksCache.OnSubmitToBlockchain(id)
}

func OnSucceeded(id string, res string) {
	defaultTasksCache.OnSucceeded(id, res)
}

func OnFailed(id string, err error) {
	defaultTasksCache.OnFailed(id, err)
}
