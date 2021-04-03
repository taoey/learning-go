package ratelimit

import (
	"fmt"
	"github.com/juju/ratelimit"
	"testing"
	"time"
)

// 使用NewBucket创建令牌桶
func TestNewBucket(t *testing.T) {

	// 创建令牌桶：流入速度 1个/s，容量 5 ，令牌桶声明后为满容量
	bucket := ratelimit.NewBucket(time.Second, 5)

	fmt.Println(bucket.Capacity(), bucket.Available())
	bucket.Take(10)
	fmt.Println(bucket.Capacity(), bucket.Available())
	time.Sleep(time.Second * 3)
	fmt.Println(bucket.Capacity(), bucket.Available())

}

// 使用NewBucketWithRate创建令牌桶
func TestNewBucketWithRate(t *testing.T) {
	// 创建令牌桶：流入速度 2个/s，容量 5 ，令牌桶声明后为满容量
	bucket := ratelimit.NewBucketWithRate(2, 5)
	fmt.Println(bucket.Capacity(), bucket.Available())
	bucket.Take(10)
	fmt.Println(bucket.Capacity(), bucket.Available()) // 5 -5
	time.Sleep(time.Second * 3)                        // 期间生成6个令牌
	fmt.Println(bucket.Capacity(), bucket.Available()) // 5 1
}

// 令牌桶应用：控制array填充速度
func TestTake(t *testing.T) {
	str := "1234567890"
	strByte := []byte(str)
	resultBytes := []byte{}
	fmt.Println(len(strByte))

	bucket := ratelimit.NewBucket(time.Second, 5)

	fmt.Println(strByte)
	current := 0

	// 令牌桶容量及字符数组监控
	go func() {
		for {
			fmt.Println(len(resultBytes), resultBytes, bucket.Available(), current)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	for len(resultBytes) < len(strByte) {
		if bucket.Available() > 0 {
			bucket.TakeAvailable(1)
			resultBytes = append(resultBytes, strByte[current])
			current = current + 1
		} else {
			time.Sleep(time.Second * 3)
		}
	}
	fmt.Println(resultBytes)

}

// wait 函数的使用
func TestWait(t *testing.T) {
	str := "1234567890"
	strByte := []byte(str)
	resultBytes := []byte{}
	fmt.Println(len(strByte))

	bucket := ratelimit.NewBucket(time.Second, 5)

	fmt.Println(strByte)
	current := 0

	//打印监控
	go func() {
		for {
			fmt.Println(len(resultBytes), resultBytes, bucket.Available(), current)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	for len(resultBytes) < len(strByte) {
		bucket.Wait(2) // wait函数，会进行阻塞，直到令牌桶中有足够的令牌才会继续执行
		resultBytes = append(resultBytes, strByte[current])
		resultBytes = append(resultBytes, strByte[current+1])
		current = current + 2
	}
	fmt.Println(resultBytes)

}
