package main

import "fmt"

/*
- 通道：用于协程之间的通信
- 通道类似于队列结构，先进先出
- 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
*/
func main() {
	message := make(chan string) //创建通道

	go func() {
		message <- "OK"
		message <- "Hello"
	}()

	msg1 := <-message
	fmt.Println(msg1)

	msg2 := <-message
	fmt.Println(msg2)

}
