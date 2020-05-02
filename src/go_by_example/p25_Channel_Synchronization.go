package main

import (
	"fmt"
	"time"
)

/*
done 通道将被用于通知其他 Go协程这个函数已经工作完毕。
*/
func work(done chan bool) {
	fmt.Println("working......")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go work(done)
	<-done //阻塞其他协程，直到其他协程执行完毕
}
