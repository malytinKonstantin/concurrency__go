package engine

type Engine interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Del(key string) bool
}

type InMemoryEngine struct {
	data map[string]string
}

func NewEngine() Engine {
	return &InMemoryEngine{
		data: make(map[string]string),
	}
}

func (e *InMemoryEngine) Set(key, value string) {
	e.data[key] = value
}

func (e *InMemoryEngine) Get(key string) (string, bool) {
	value, exists := e.data[key]
	return value, exists
}

func (e *InMemoryEngine) Del(key string) bool {
	_, exists := e.data[key]
	if exists {
		delete(e.data, key)
	}
	return exists
}
