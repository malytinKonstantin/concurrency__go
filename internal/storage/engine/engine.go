package engine

import (
	"sync"
)

type Engine interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Del(key string) bool
}

type InMemoryEngine struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewEngine() Engine {
	return &InMemoryEngine{
		data: make(map[string]string),
	}
}

func (e *InMemoryEngine) Set(key, value string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.data[key] = value
}

func (e *InMemoryEngine) Get(key string) (string, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	value, exists := e.data[key]
	return value, exists
}

func (e *InMemoryEngine) Del(key string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	_, exists := e.data[key]
	if exists {
		delete(e.data, key)
	}
	return exists
}
