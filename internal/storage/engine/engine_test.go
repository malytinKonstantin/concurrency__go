package engine

import "testing"

func TestInMemoryEngine(t *testing.T) {
	engine := NewEngine()

	// Тестируем Set и Get
	engine.Set("key1", "value1")
	value, exists := engine.Get("key1")
	if !exists || value != "value1" {
		t.Errorf("Ожидалось значение 'value1' для 'key1', получено '%s'", value)
	}

	// Тестируем Del
	deleted := engine.Del("key1")
	if !deleted {
		t.Errorf("Ожидалось успешное удаление 'key1'")
	}

	_, exists = engine.Get("key1")
	if exists {
		t.Errorf("Не ожидалось, что 'key1' будет существовать после удаления")
	}

	// Тестируем Del для несуществующего ключа
	deleted = engine.Del("nonexistent")
	if deleted {
		t.Errorf("Не ожидалось успешного удаления для несуществующего ключа")
	}
}
