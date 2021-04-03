package golang_queue

// queue 相关接口
// 一个队列常见的功能
// 入队，出队，队列长度

/*
参考java的队列相关方法
　　add        增加一个元索                     如果队列已满，则抛出一个IIIegaISlabEepeplian异常
　　remove   移除并返回队列头部的元素    如果队列为空，则抛出一个NoSuchElementException异常
　　element  返回队列头部的元素             如果队列为空，则抛出一个NoSuchElementException异常

　　offer       添加一个元素并返回true       如果队列已满，则返回false
　　poll         移除并返问队列头部的元素    如果队列为空，则返回null
　　peek       返回队列头部的元素             如果队列为空，则返回null

*/

type Queue interface {
	// 向队列末尾添加一个元素
	Add(data interface{})

	// 移除并返回队列头部的元素
	Remove() (interface{}, error)

	// 返问队列头部的元素
	Element() (interface{}, error)
}
