package goroutine

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.LstdFlags | log.Lshortfile)
}
func TestMainTest2(t *testing.T) {
	log.Println("Program started...")
	err := operationWithTimeout2(time.Second * 5)
	if err != nil {
		log.Println(err)
	}
	log.Println("Program finished...")
}

// 新建方法：包装slowOperation，进行超时控制
func operationWithTimeout2(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := make(chan bool)
	// 起一个协程执行slowOperation
	go func() {
		slowOperation2(c)
	}()

	select {
	case <-ctx.Done():
		//log.Println(ctx.Err())
		return errors.New("send timeout") // 返回自定义错误
	case result := <-c:
		log.Println(result)
		return nil
	}
	log.Println("operationWithTimeout finished...")
	return nil
}

func slowOperation2(c chan bool) {
	log.Println("slowOperation started...")
	time.Sleep(60 * time.Second) // 模拟耗时程序， 其中可以包含 c <- false
	c <- true
	log.Println("slowOperation finished...")
}
