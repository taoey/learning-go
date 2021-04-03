/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/17 16:17
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

// 打印1-5
func main() {
	gen := func() <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				dst <- n
				n++
			}
		}()
		return dst
	}
	for n := range gen() {
		fmt.Println(n)
		if n >= 5 {
			break
		}
	}

	time.Sleep(time.Second)
	fmt.Println("leak run groutine:", runtime.NumGoroutine()-1) // leak run groutine: 1
}
