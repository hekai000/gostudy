package lfucache

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	cache := Constructor(2)

	// Happy path
	cache.Put(1, 1)
	cache.Put(2, 2)
	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %v, want 1", got)
	}

	// Adding a new key should evict the least frequently used key (2)
	cache.Put(3, 3)
	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) = %v, want -1", got)
	}

	// Adding a new key should evict key 1 because key 3 now has the same frequency
	cache.Put(4, 4)
	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %v, want 1", got)
	}

	if got := cache.Get(3); got != -1 {
		t.Errorf("Get(3) = %v, want -1", got)
	}

	if got := cache.Get(4); got != 4 {
		t.Errorf("Get(4) = %v, want 4", got)
	}

	// Edge case - capacity is 0
	emptyCache := Constructor(0)
	emptyCache.Put(1, 1)
	if got := emptyCache.Get(1); got != -1 {
		t.Errorf("Get(1) = %v, want -1", got)
	}

	// Edge case - overwriting an existing key
	cache.Put(1, 30) // Update value of key 3
	if got := cache.Get(1); got != 30 {
		t.Errorf("Get(1) = %v, want 30", got)
	}
}
