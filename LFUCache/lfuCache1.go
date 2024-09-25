package lfucache

// import "container/list"

// type LFUCache struct {
// 	capacity int
// 	lists    map[int]*list.List
// 	min      int
// 	nodes    map[int]*list.Element
// }

// type node struct {
// 	key       int
// 	value     int
// 	frequency int
// }

// func Constructor(capacity int) LFUCache {
// 	return LFUCache{
// 		capacity: capacity,
// 		lists:    make(map[int]*list.List),
// 		min:      0,
// 		nodes:    make(map[int]*list.Element),
// 	}
// }

// func (cache *LFUCache) Get(key int) int {
// 	value, ok := cache.nodes[key]
// 	if !ok {
// 		return -1
// 	}
// 	currentNode := value.Value.(*node)
// 	cache.lists[currentNode.frequency].Remove(value)
// 	currentNode.frequency += 1
// 	if _, ok := cache.lists[currentNode.frequency]; !ok {
// 		cache.lists[currentNode.frequency] = list.New()
// 	}
// 	newList := cache.lists[currentNode.frequency]
// 	newNode := newList.PushFront(currentNode)
// 	cache.nodes[key] = newNode
// 	if cache.min == currentNode.frequency-1 && cache.lists[currentNode.frequency-1].Len() == 0 {
// 		cache.min++
// 	}

// 	return currentNode.value
// }

// func (cache *LFUCache) Put(key int, value int) {
// 	if cache.capacity == 0 {
// 		return
// 	}
// 	if currentValue, ok := cache.nodes[key]; ok {
// 		currentNode := currentValue.Value.(*node)
// 		currentNode.value = value
// 		cache.Get(key)
// 		return
// 	}
// 	if len(cache.nodes) == cache.capacity {
// 		currentList := cache.lists[cache.min]
// 		backNode := currentList.Back()
// 		delete(cache.nodes, backNode.Value.(*node).key)
// 		currentList.Remove(backNode)

// 	}
// 	cache.min = 1
// 	currentNode := &node{
// 		key:       key,
// 		value:     value,
// 		frequency: 1,
// 	}

// 	if _, ok := cache.lists[1]; !ok {
// 		cache.lists[1] = list.New()
// 	}
// 	newList := cache.lists[1]
// 	newNode := newList.PushFront(currentNode)
// 	cache.nodes[key] = newNode
// }
