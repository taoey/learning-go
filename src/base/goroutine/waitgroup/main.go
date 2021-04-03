/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/11 9:59
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func testFunc(spend int) {
	defer wp.Done()

	time.Sleep(time.Second * time.Duration(spend))
	fmt.Println(spend, "done")
}

var wp sync.WaitGroup

func main() {
	wp = sync.WaitGroup{}

	start := time.Now().Second()

	for i := 1; i < 10; i++ {
		wp.Add(1)
		go testFunc(i)
	}

	wp.Wait()

	end := time.Now().Second()
	allSpend := end - start
	fmt.Println("all spend:", allSpend)
}
