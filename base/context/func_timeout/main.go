package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.LstdFlags | log.Lshortfile)
}
func main() {
	log.Println("Program started...")
	err := operationWithTimeout(time.Second * 5)
	if err != nil {
		log.Println(err)
	}
	log.Println("Program finished...")
}

// 新建方法：包装slowOperation，进行超时控制
func operationWithTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	exit := make(chan bool)
	// 起一个协程执行slowOperation
	go func() {
		slowOperation(exit)
	}()

	select {
	case <-ctx.Done():
		//log.Println(ctx.Err())
		return errors.New("send timeout") // 返回自定义错误
	case result := <-exit:
		log.Println(result)
		return nil
	}
	log.Println("operationWithTimeout finished...")
	return nil
}

func slowOperation(exit chan bool) {
	log.Println("slowOperation started...")
	time.Sleep(60 * time.Second) // 模拟耗时程序， 其中可以包含 exit <- false
	exit <- true
	log.Println("slowOperation finished...")
}
