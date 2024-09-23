package lrucache

// import "container/list"

// type LRUCache struct {
// 	Cap  int
// 	Keys map[int]*list.Element
// 	List *list.List
// }

// type pair struct {
// 	k, v int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		Cap:  capacity,
// 		Keys: make(map[int]*list.Element),
// 		List: list.New(),
// 	}
// }

// func (c *LRUCache) Get(key int) int {
// 	if ele, ok := c.Keys[key]; ok {
// 		c.List.MoveToFront(ele)
// 		return ele.Value.(pair).v
// 	}
// 	return -1
// }

// func (c *LRUCache) Put(key int, value int) {
// 	if ele, ok := c.Keys[key]; ok {
// 		ele.Value = pair{key, value}
// 		c.List.MoveToFront(ele)
// 	} else {
// 		if c.List.Len() == c.Cap { // cache is full
// 			delete(c.Keys, c.List.Back().Value.(pair).k)
// 			c.List.Remove(c.List.Back())
// 		}
// 		ele := c.List.PushFront(pair{key, value})
// 		c.Keys[key] = ele
// 	}
// }
