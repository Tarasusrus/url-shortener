package stores

import (
	"testing"
)

// Тест на проверку поведения при передаче несуществующего id в функцию Get.
func TestStore_GetNonExistent(t *testing.T) {
	store := NewStore()
	_, ok := store.Get("nonexistent")
	if ok {
		t.Error("Expected Get to return false for nonexistent id")
	}
}

// Тест на проверку возвращаемого id функцией Set.
func TestStore_SetReturnsID(t *testing.T) {
	store := NewStore()
	url := "http://example.com"
	id := store.Set(url)
	if id == "" {
		t.Error("Expected Set to return a non-empty id")
	}
}

// Тест на проверку уникальности сгенерированных id.
func TestGenerateIDUnique(t *testing.T) {
	idSet := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		id := generateID()
		if idSet[id] {
			t.Errorf("Expected generateID to generate unique IDs, got duplicate for %s", id)
		}
		idSet[id] = true
	}
}

// Тест на проверку длины сгенерированного id.
func TestGenerateIDLength(t *testing.T) {
	id := generateID()
	if len(id) != 10 {
		t.Errorf("Expected generateID to return id of 10 characters, got %d characters", len(id))
	}
}
