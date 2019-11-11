package ratelimit

import (
	"fmt"
	"github.com/juju/ratelimit"
	"testing"
	"time"
)

func Test01(t *testing.T) {

	bucket := ratelimit.NewBucket(time.Second, 5)

	fmt.Println(bucket.Capacity(), bucket.Available())
	bucket.Take(10)
	fmt.Println(bucket.Capacity(), bucket.Available())
	time.Sleep(time.Second * 3)
	fmt.Println(bucket.Capacity(), bucket.Available())

}

func Test02(t *testing.T) {
	str := "1234567890"
	strByte := []byte(str)
	resultBytes := []byte{}
	fmt.Println(len(strByte))

	bucket := ratelimit.NewBucket(time.Second, 5)

	fmt.Println(strByte)
	current := 0
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

func Test03(t *testing.T) {
	str := "1234567890"
	strByte := []byte(str)
	resultBytes := []byte{}
	fmt.Println(len(strByte))

	bucket := ratelimit.NewBucket(time.Second, 5)

	fmt.Println(strByte)
	current := 0
	go func() {
		for {
			fmt.Println(len(resultBytes), resultBytes, bucket.Available(), current)
			time.Sleep(time.Millisecond * 200)
		}
	}()
	for len(resultBytes) < len(strByte) {
		bucket.Wait(2)
		resultBytes = append(resultBytes, strByte[current])
		resultBytes = append(resultBytes, strByte[current+1])
		current = current + 2
	}
	fmt.Println(resultBytes)

}
