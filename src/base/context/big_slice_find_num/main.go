package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

const (
	TIME_OUT_SECOND = 5
)

func FindNum(arr []int, target int) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*TIME_OUT_SECOND)
	defer func() {
		fmt.Println("cancel before", runtime.NumGoroutine())
		cancelFunc()
		time.Sleep(time.Second)
		fmt.Println("cancel after", runtime.NumGoroutine())
	}()

	// 确定goroution个数
	goroutineNum := runtime.NumCPU()
	exitChan := make(chan int)
	if len(arr) < goroutineNum {
		go subFindNum(ctx, arr, target, exitChan)
	} else {
		// 子切片划分
		subLength := len(arr) / goroutineNum
		for i := 0; i < goroutineNum; i++ {
			index := i
			var subArr []int
			if index == goroutineNum-1 { //剩余元素都归到最后一个切片中
				subArr = arr[index*subLength:]
			} else {
				subArr = arr[index*subLength : (index+1)*subLength]
			}
			go subFindNum(ctx, subArr, target, exitChan)
		}
	}
	fmt.Println("current cpu num:", runtime.NumCPU())
	fmt.Println("current cpu num:", runtime.NumGoroutine())
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
		return
	case <-exitChan:
		fmt.Println("find it")
		return
	}
	return
}

// 子切片查询
func subFindNum(ctx context.Context, subArr []int, targetNum int, exitChan chan int) {
	for _, val := range subArr {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if val == targetNum {
			exitChan <- 1
			return
		}
	}
}

func main() {
	maxNum := 100000
	arr := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		arr[i] = i
	}
	targetNum := 1000
	FindNum(arr, targetNum)
}
