/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/17 16:31
 */
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	cancel() // 提前调用cancel函数，关闭context
	time.Sleep(time.Second)
	fmt.Println("leak run groutine:", runtime.NumGoroutine()-1) // leak run groutine: 0
}
