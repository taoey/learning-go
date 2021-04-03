package LRU缓存

import (
	"container/list"
	"fmt"
	"testing"
)

// LRU结构体
type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}

// key-value 结构体
type Kv struct {
	k int
	v int
}

// 创建一个LRU对象
func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		cap:   cap,
		list:  list.New(),
		cache: make(map[int]*list.Element),
	}
}

// 根据key获取value,没获取到返回-1
// 通过map获取元素，获取到后把元素移动到头部
func (this *LRUCache) Get(k int) int {
	if item, ok := this.cache[k]; ok {
		this.list.MoveToFront(item)
		return item.Value.(Kv).v
	}
	return -1
}

// 放置一个key-value对
// 需要注意容量放满时的情况
func (this *LRUCache) Put(k, v int) {
	if elem, ok := this.cache[k]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Kv{k, v} // update value
	} else {
		if this.list.Len() >= this.cap {
			this.removeEleOldest()
		}
		this.addEleToCache(k, v)
	}
}

func (this *LRUCache) addEleToCache(k, v int) error {
	this.list.PushFront(Kv{k, v})
	this.cache[k] = this.list.Front()
	return nil
}

func (this *LRUCache) removeEleOldest() error {
	backEle := this.list.Back()
	delete(this.cache, backEle.Value.(Kv).k) // delete element form map
	this.list.Remove(backEle)                // remove element from list
	return nil
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

func TestList(t *testing.T) {
	myList := list.New()

	myList.PushFront(Kv{1, 2})

	// update demo1
	//myList.Front().Value.(Kv).v = 10  // update error: cannot assign to myList.Front().Value.(Kv).v

	// update demo2
	ele := myList.Front().Value.(Kv)
	ele.v = 10
	fmt.Println(ele, myList.Front().Value) // debug log : {1 10} {1 2}

	myList.Front().Value = ele        // only update value
	fmt.Println(myList.Front().Value) // update succeed {1 10}
}
