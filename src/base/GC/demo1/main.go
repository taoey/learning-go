/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/31 14:12
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		for {
		}
	}()
	fmt.Println(3 + 1)
	time.Sleep(time.Millisecond)
	runtime.GC()
	println("OK")
}
