package lfucache2

import "container/heap"

type LFUCache struct {
	capacity int
	pq       PriorityQueue
	hash     map[int]*Item
	counter  int
}

type Item struct {
	key       int
	value     int
	frequency int
	count     int
	index     int
}

func Constructor(capacity int) LFUCache {
	lfu := LFUCache{
		pq:       PriorityQueue{},
		hash:     make(map[int]*Item, capacity),
		capacity: capacity,
	}
	return lfu
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].frequency == pq[j].frequency {
		return pq[i].count < pq[j].count
	}
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  //avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, frequency int, count int) {
	item.value = value
	item.count = count
	item.frequency = frequency
	heap.Fix(pq, item.index)
}

func (cache *LFUCache) Get(key int) int {
	if cache.capacity == 0 {
		return -1
	}
	if item, ok := cache.hash[key]; ok {
		cache.counter++
		cache.pq.update(item, item.value, item.frequency+1, cache.counter)
		return item.value
	}
	return -1
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}
	cache.counter++
	if item, ok := cache.hash[key]; ok {

		cache.pq.update(item, value, item.frequency+1, cache.counter)
	}
	if len(cache.pq) == cache.capacity {
		item := heap.Pop(&cache.pq).(*Item)
		delete(cache.hash, item.key)
	}

	item := &Item{key: key, value: value, frequency: 1, count: cache.counter, index: -1}
	cache.hash[key] = item
	heap.Push(&cache.pq, item)

}
