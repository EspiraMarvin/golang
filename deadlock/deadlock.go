package main

import "sync"

type Collection struct {
	Mutex sync.Mutex
	Data  map[string]string
}

func NewCollection() Collection {
	return Collection{
		Data: make(map[string]string),
	}
}
func (c *Collection) Has(key string) bool {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
}
