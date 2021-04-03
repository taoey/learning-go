package main

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
)

var (
	num   int // 要输出的最大值
	line  = 0 // 通道发送计数器
	exit  = make(chan bool)
	chans []chan int // 要初始化的协程数量
)

func main() {
	// 开启4个协程
	chans = []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int)}

	// 多协程启动入口
	go ChanWork(chans[0])

	// 接收要输出的最大数
	fmt.Println("输入要输出的最大数值:")
	_, ok := fmt.Scanf("%d\n", &num)
	if ok == io.EOF {
		return
	}
	// 触发协程同步执行
	chans[0] <- 1

	// 执行结束
	if <-exit {
		return
	}
}

func ChanWork(c chan int) {
	// 协程数
	lens := len(chans)
	for {
		// count为输出计数器
		if count := <-chans[line]; count <= num {
			fmt.Println("channel "+strconv.Itoa(line)+" -> ", count, runtime.NumGoroutine())
			count++

			// 下一个发送通道
			line++
			if line >= lens {
				line = 0 //循环，防止索引越界
			}
			go ChanWork(chans[line])

			chans[line] <- count

		} else {
			// 通道编号问题处理
			id := 0
			if line == 0 {
				id = lens - 1
			} else {
				id = line - 1
			}
			fmt.Println("在通道" + strconv.Itoa(id) + "执行完成")
			close(exit)
			return
		}
	}
}
