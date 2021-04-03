/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/31 14:24
 */
package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here
	// My Example
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d\n", i)
		time.Sleep(1 * time.Second)
	}
	return
}
