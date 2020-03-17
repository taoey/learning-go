package goroutine

import (
	"fmt"
	"testing"
	"time"
)

var pool = NewGoroutinePool()

func TestGoroutinePool01(t *testing.T) {
	pool.Run(&ConvertTask{
		CallBack: func() {
			time.Sleep(time.Second * 10)
			fmt.Println("hello1")

		},
	})

	pool.Run(&ConvertTask{
		CallBack: func() {
			fmt.Println("hello2")
		},
	})

	pool.Run(&ConvertTask{
		CallBack: func() {
			fmt.Println("hello3")
		},
	})

	time.Sleep(time.Second * 60)
}
