package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ExitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) // 第二个参数不设置表示监听所有的信号

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("退出:", s)
				ExitFunc()
			default:
				fmt.Println("其他信号:", s)
				// do something
			}
		}
	}()

	fmt.Println("启动了程序")
	sum := 0
	for {
		sum++
		fmt.Println("休眠了:", sum, "秒")
		time.Sleep(1 * time.Second)
		if sum == 3 {
			c <- syscall.SIGINT
		}
	}

	fmt.Println("main do something")
}
