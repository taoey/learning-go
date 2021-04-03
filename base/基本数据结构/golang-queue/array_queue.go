package golang_queue

import (
	"errors"
	"sync"
)

type ArrayQueue struct {
	array []interface{} // 底层切片
	size  int           // 队列的元素数量
	lock  sync.Mutex    // 为了并发安全使用的锁
}

// 向队列末尾添加一个元素
func (this *ArrayQueue) Add(data interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	this.array = append(this.array, data)

	// 队中元素数量+1
	this.size = this.size + 1
}

// 移除并返回队列头部的元素
func (this *ArrayQueue) Remove() (interface{}, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.size == 0 {
		err := errors.New("queue is empty")
		return nil, err
	}

	head := this.array[0]
	//this.array = this.array[1:]
	//this.size = this.size-1

	newArray := make([]interface{}, this.size-1, this.size-1)
	for i := 1; i < this.size; i++ {
		// 从老数组的第一位开始进行数据移动
		newArray[i-1] = this.array[i]
	}
	this.array = newArray
	return head, nil
}

// 移除并返问队列头部的元素
func (this *ArrayQueue) Element() (interface{}, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.size == 0 {
		err := errors.New("queue is empty")
		return nil, err
	}

	head := this.array[0]
	return head, nil
}
