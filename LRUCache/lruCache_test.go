package lrucache

// import (
// 	"testing"
// )

// func TestLRUCache(t *testing.T) {
// 	cache := Constructor(2)

// 	// Happy path
// 	cache.Put(1, 1)
// 	cache.Put(2, 2)
// 	if got := cache.Get(1); got != 1 {
// 		t.Errorf("Get(1) = %v; want 1", got)
// 	}

// 	cache.Put(3, 3) // evicts key 2
// 	if got := cache.Get(2); got != -1 {
// 		t.Errorf("Get(2) = %v; want -1", got)
// 	}

// 	cache.Put(4, 4) // evicts key 1
// 	if got := cache.Get(1); got != -1 {
// 		t.Errorf("Get(1) = %v; want -1", got)
// 	}
// 	if got := cache.Get(3); got != 3 {
// 		t.Errorf("Get(3) = %v; want 3", got)
// 	}
// 	if got := cache.Get(4); got != 4 {
// 		t.Errorf("Get(4) = %v; want 4", got)
// 	}

// 	// Edge cases

// 	if got := cache.Get(3); got != 3 {
// 		t.Errorf("Get(3) = %v; want 3", got)
// 	}
// 	cache.Put(3, 6) // update key 3
// 	if got := cache.Get(3); got != 6 {
// 		t.Errorf("Get(3) = %v; want 6", got)
// 	}
// 	cache.Put(6, 6) // evicts key 4
// 	if got := cache.Get(4); got != -1 {
// 		t.Errorf("Get(4) = %v; want -1", got)
// 	}
// }
