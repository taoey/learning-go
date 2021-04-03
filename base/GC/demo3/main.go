/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/31 14:31
 */
package main

import (
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

var cache = map[interface{}]interface{}{}

func keepalloc() {
	for i := 0; i < 10; i++ {
		m := make([]byte, 1<<10)
		cache[i] = m
	}
}

func keepalloc2() {
	for i := 0; i < 100000; i++ {
		go func() {
			select {}
		}()
	}
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	keepalloc()
	keepalloc2()
	runtime.GC()
	time.Sleep(time.Second)
}
