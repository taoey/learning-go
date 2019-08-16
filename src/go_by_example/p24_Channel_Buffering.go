package main

import "fmt"

/*

默认通道是无缓冲的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。
可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值。
*/
func main() {
	message := make(chan string, 2) //可缓存两个消息的通道

	message <- "Hello"
	message <- "Go"

	fmt.Println(<-message)
	fmt.Println(<-message)

}
