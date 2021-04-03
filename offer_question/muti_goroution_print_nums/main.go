package main

import (
	"log"
)

//WorkChanOut 工作流
func WorkChanOut(i int, ch chan int, stopC chan<- bool) {
	for {
		sum := <-ch
		if sum == 10 {
			stopC <- true
			return
		}
		sum++
		log.Println(i, ":", sum)
		ch <- sum
	}
}

//ServeAuto 服务
func main() {

	ch := make(chan int)
	stopC := make(chan bool)
	for i := 0; i < 3; i++ {
		go WorkChanOut(i, ch, stopC)
	}
	ch <- 0 // 初始值 0
	log.Println("stop", <-stopC)
}
