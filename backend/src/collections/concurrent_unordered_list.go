package collections

import (
	"sync"
)

func NewUnOrderedList[T any]() ConcurrentUnorderedList[T] {
	return ConcurrentUnorderedList[T]{
		data: []*T{},
		lock: sync.RWMutex{},
	}
}

type ConcurrentUnorderedList[T any] struct {
	data []*T
	lock sync.RWMutex
}

func (c *ConcurrentUnorderedList[T]) Append(value *T) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data = append(c.data, value)
}

func (c *ConcurrentUnorderedList[T]) Delete(value *T) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for i, v := range c.data {
		if v == value {
			c.data[i], c.data[len(c.data)-1] = c.data[len(c.data)-1], c.data[i]
			c.data = c.data[:len(c.data)-1]
			break
		}
	}
}

func (c *ConcurrentUnorderedList[T]) Apply(fn func(*T)) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	for _, v := range c.data {
		fn(v)
	}
}
