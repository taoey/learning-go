package _46_LRU缓存

import (
	"container/list"
	"fmt"
	"testing"
)

// LRU结构体
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

// key-value 结构体
type Pair struct {
	key   int
	value int
}

// 创建一个LRU对象
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

// 根据key获取value,没获取到返回-1
// 通过map获取元素，获取到后把元素移动到头部
func (this *LRUCache) Get(key int) int {
	if item, ok := this.cache[key]; ok {
		this.list.MoveToFront(item)
		return item.Value.(Pair).value
	}
	return -1
}

// 放置一个key-value对
// 需要注意容量放满时的情况
func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if this.list.Len() >= this.capacity {
			delete(this.cache, this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(Pair{key, value})
		this.cache[key] = this.list.Front()
	}
}

func TestLRU(t *testing.T) {
	lruCache := NewLRUCache(3)

	lruCache.Put(1, 11)
	lruCache.Put(2, 22)
	lruCache.Put(3, 33)
	lruCache.Put(4, 44)

	aa := lruCache.Get(3)
	fmt.Println("aa:", aa)

	bb := lruCache.Get(4)
	fmt.Println("bb", bb)

	cc := lruCache.Get(1)
	fmt.Println("cc", cc)
}
