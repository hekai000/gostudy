package lfucache

import "container/list"

type LFUCache struct {
	nodesDict     map[int]*list.Element
	frequencyDict map[int]*list.List
	capacity      int
	min           int
}

type node struct {
	key, value int
	frequency  int
}

func Constructor(capacity int) LFUCache {
	lfu := LFUCache{
		nodesDict:     make(map[int]*list.Element),
		frequencyDict: make(map[int]*list.List),
		capacity:      capacity,
		min:           0,
	}
	return lfu
}

// Get方法用于获取key对应的值
func (cache *LFUCache) Get(key int) int {

	// 如果cache的容量为0，则返回-1
	if cache.capacity == 0 {
		return -1
	}

	// 从nodesDict中获取key对应的node
	nd, ok := cache.nodesDict[key]
	// 如果key不存在，则返回-1
	if !ok {
		return -1
	}
	// 获取node的值
	currentNode := nd.Value.(*node)
	// 从frequencyDict中获取当前node的频率对应的链表，并移除该node
	cache.frequencyDict[currentNode.frequency].Remove(nd)
	// 增加node的频率
	currentNode.frequency += 1

	// 如果frequencyDict中不存在当前node的频率，则新建一个链表
	if _, ok := cache.frequencyDict[currentNode.frequency]; !ok {
		cache.frequencyDict[currentNode.frequency] = list.New()
	}
	// 将当前node加入frequencyDict中对应频率的链表
	cn := cache.frequencyDict[currentNode.frequency].PushFront(currentNode)
	// 更新nodesDict中key对应的node
	cache.nodesDict[key] = cn
	// 如果min的值等于当前node的频率减1，且frequencyDict中min对应的链表为空，则更新min的值
	if cache.min == currentNode.frequency-1 && cache.frequencyDict[cache.min].Len() == 0 {
		cache.min = currentNode.frequency
	}
	// 返回当前node的值
	return currentNode.value

}

// This function adds a key-value pair to the LFUCache.
// If the key already exists in the cache, it updates the value and returns.
// If the cache's capacity is reached, it removes the least frequently used item.
// Otherwise, it creates a new node with the given key and value, and adds it to the cache.
func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}

	// If the key already exists in the cache, update its value and return.
	if currentNode, ok := cache.nodesDict[key]; ok {
		currentNode.Value.(*node).value = value
		cache.Get(key)
		return
	}

	// If the cache's capacity is reached, remove the least frequently used item.
	if cache.capacity == len(cache.nodesDict) {
		backNode := cache.frequencyDict[cache.min].Back()
		delete(cache.nodesDict, backNode.Value.(*node).key)
		cache.frequencyDict[cache.min].Remove(backNode)
	}

	// Create a new node with the given key and value, and add it to the cache.
	cache.min = 1
	NewNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}

	// If the frequency of 1 does not exist in the frequencyDict, create a new list.
	if _, ok := cache.frequencyDict[1]; !ok {
		cache.frequencyDict[1] = list.New()
	}
	// Add the new node to the front of the list with frequency 1.
	new := cache.frequencyDict[1].PushFront(NewNode)
	cache.nodesDict[key] = new

}
