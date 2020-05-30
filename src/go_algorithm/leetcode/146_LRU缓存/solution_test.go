package _46_LRU缓存

import (
	"container/list"
	"fmt"
	"testing"
)

// 键值对
type Pair struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	list     *list.List
	cacheMap map[int]*list.Element
}

// 构造函数
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cacheMap: make(map[int]*list.Element),
	}
}

// 获取数据
func (this *LRUCache) Get(key int) int {
	if element, ok := this.cacheMap[key]; ok {
		this.list.MoveToFront(element)
		return element.Value.(Pair).value
	}
	return -1
}

// 放置数据
func (this *LRUCache) Put(key int, value int) {
	// 判断是否在缓存中，已存在则直接调整数据位置,更新key-value
	if element, ok := this.cacheMap[key]; ok {
		this.list.MoveToFront(element)
		element.Value = Pair{key, value}
	} else {
		// 判断缓存是否已满，已满则进行数据淘汰
		if this.list.Len() >= this.capacity {
			// 删除map中的数据
			delete(this.cacheMap, this.list.Back().Value.(Pair).key)
			// 删除list中的数据
			this.list.Remove(this.list.Back())
		}
		front := this.list.PushFront(Pair{key, value})
		this.cacheMap[key] = front
	}
}

func TestLRU(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)             // 返回  1
	cache.Put(3, 3)          // 该操作会使得关键字 2 作废
	element1 := cache.Get(2) // 返回 -1 (未找到)
	fmt.Println(element1)
	cache.Put(4, 4)     // 该操作会使得关键字 1 作废
	ele := cache.Get(1) // 返回 -1 (未找到)
	fmt.Println(ele)
	cache.Get(3) // 返回  3
	cache.Get(4) // 返回  4

}
