package test02

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/**
select 多队列的随机选择器
*/

func sender(messages chan string) {
	for i := 0; i < 5; i++ {
		messages <- "s" + strconv.Itoa(i)
		fmt.Println("ch1 push", "s"+strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
}

// 用于模拟高延时程序
func sender2(messages chan int) {
	for i := 0; i < 10; i++ {
		messages <- i
		fmt.Println("ch2 push", i)
		time.Sleep(time.Second * 20)
	}
	close(messages)
}

func Test01(t *testing.T) {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go sender(ch1)
		go sender2(ch2)
	}

	for i := 0; i < 1000; i++ { // 运行1000次 select
		// 每次只运行一次case
		// 多队列的随机选择器
		select {
		case str, ch1Check := <-ch1: // 当channel有值时，进行输出操作
			if !ch1Check {
				fmt.Println("ch1Check false")
			}
			fmt.Println("ch1 pop", str)
		case num, ch2Check := <-ch2:
			if !ch2Check {
				fmt.Println("ch2Check false")
			}
			fmt.Println("ch2 pop", num)
		}
	}
	fmt.Println("select block end")
	time.Sleep(time.Second * 60)
}

// select 等待机制
func Test02(t *testing.T) {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go sender(ch1)
		go sender2(ch2)
	}

	for {
		select {
		case str, ch1Check := <-ch1: // 当channel有值时，进行输出操作
			if !ch1Check {
				fmt.Println("ch1Check false")
			}
			fmt.Println("ch1 pop", str)
		case num, ch2Check := <-ch2:
			if !ch2Check {
				fmt.Println("ch2Check false")
			}
			fmt.Println("ch2 pop", num)
		}
	}
}
