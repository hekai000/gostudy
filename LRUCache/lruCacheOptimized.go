package lrucache

type LRUCache struct {
	Cap        int
	Keys       map[int]*Node
	head, tail *Node
}

type Node struct {
	Key, Value int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}

}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.Keys[key]; ok {
		cache.Remove(node)
		cache.Add(node)
		return node.Value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.Keys[key]; ok {
		node.Value = value
		cache.Remove(node)
		cache.Add(node)
		return
	} else {
		node = &Node{Key: key, Value: value}
		cache.Keys[key] = node
		cache.Add(node)

	}
	if len(cache.Keys) > cache.Cap {
		delete(cache.Keys, cache.tail.Key)
		cache.Remove(cache.tail)
	}
}

func (cache *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = cache.head
	if cache.head != nil {
		cache.head.Prev = node
	}
	cache.head = node
	if cache.tail == nil {
		cache.tail = node
		cache.tail.Next = nil
	}
}
func (cache *LRUCache) Remove(node *Node) {
	if node == cache.head {
		cache.head = node.Next
		node.Next = nil
		return
	}
	if node == cache.tail {
		cache.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
