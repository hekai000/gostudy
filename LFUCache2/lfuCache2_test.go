package lfucache2

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	cache := Constructor(2)

	// Happy path: Add items
	cache.Put(1, 1)
	if res := cache.Get(1); res != 1 {
		t.Errorf("Expected 1, got %d", res)
	}

	cache.Put(2, 2)
	if res := cache.Get(2); res != 2 {
		t.Errorf("Expected 2, got %d", res)
	}

	// Update value and check frequency
	cache.Put(1, 10)
	if res := cache.Get(1); res != 10 {
		t.Errorf("Expected 10, got %d", res)
	}

	// Adding another item should evict the least frequently used item (2)
	cache.Put(3, 3)
	if res := cache.Get(2); res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}
	if res := cache.Get(3); res != 3 {
		t.Errorf("Expected 3, got %d", res)
	}

	// Edge case: Cache capacity is 0
	emptyCache := Constructor(0)
	emptyCache.Put(1, 1)
	if res := emptyCache.Get(1); res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}

	// Edge case: Getting non-existing item
	if res := cache.Get(4); res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}

	// Edge case: When adding a new item to full capacity cache
	cache.Put(4, 4)
	if res := cache.Get(3); res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}
	if res := cache.Get(4); res != 4 {
		t.Errorf("Expected 4, got %d", res)
	}
}
