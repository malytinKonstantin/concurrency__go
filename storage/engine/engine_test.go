package engine

import "testing"

func TestInMemoryEngine(t *testing.T) {
	engine := NewEngine()

	// Test Set and Get
	engine.Set("key1", "value1")
	value, exists := engine.Get("key1")
	if !exists || value != "value1" {
		t.Errorf("Expected to get 'value1' for 'key1', got '%s'", value)
	}

	// Test Del
	deleted := engine.Del("key1")
	if !deleted {
		t.Errorf("Expected 'key1' to be deleted")
	}

	_, exists = engine.Get("key1")
	if exists {
		t.Errorf("Did not expect 'key1' to exist after deletion")
	}

	// Test Del non-existing key
	deleted = engine.Del("nonexistent")
	if deleted {
		t.Errorf("Did not expect deletion success for non-existing key")
	}
}
