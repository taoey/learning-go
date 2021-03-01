package main

import (
	"fmt"
	"testing"
	"time"
)

// 交替打印
// golang 两个协程交替打印1-20
// 使用无缓冲channel实现

func TwoGoroutinePrintNum(max int) {
	ch := make(chan int)

	go func() {
		for i := 0; i <= max; i++ {
			ch <- 1
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	go func() {
		for i := 0; i <= max; i++ {
			<-ch
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()

	time.Sleep(time.Second * 3)
}

func TestTwoGoroutinePrintNum(t *testing.T) {
	TwoGoroutinePrintNum(20)
}
