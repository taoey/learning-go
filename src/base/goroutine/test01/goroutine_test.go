package test01

import (
	"fmt"
	"testing"
	"time"
)

func sender(messages chan string) {
	messages <- "hello 1"
	messages <- "hello 2"
	messages <- "hello 3"
	messages <- "hello 4"
}

func sender2(messages chan string) {
	time.Sleep(time.Second * 2)
	str := <-messages
	str = str + "world"
	messages <- str

	// messages 写入端已经结束，可以在此处进行关闭
	// 编程建议：如何合理的关闭channel？ 答:我们需要明确自己的程序的运行机制，在合理位置关闭channel
	close(messages)
}

/**
 * goroutine遵循FIFO理念
 */
func Test01(t *testing.T) {
	messages := make(chan string, 3)
	go sender(messages)
	go sender2(messages)
	time.Sleep(time.Second * 3)

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

/**
 * range 缓冲队列的逐个获取方法
 */
func Test02(t *testing.T) {
	messages := make(chan string, 3)
	go sender(messages)
	go sender2(messages)
	time.Sleep(time.Second * 3)

	// 使用 range 消费channel ，只有channel==nil时才会停止，所以写程序是一定不要忘记关闭channel（面试点）
	// 否则会出现死锁问题 fatal error: all goroutines are asleep - deadlock!
	for message := range messages {
		fmt.Println(message)
	}

	fmt.Println("software stop !!!")
}
