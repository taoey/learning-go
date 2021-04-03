package test04

import (
	"fmt"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	amount := 20
	array := make([]int, amount+1)
	for i := 0; i < amount; i++ {
		go func() {
			array[i] = i
			fmt.Println(array)
		}()

	}
	time.Sleep(time.Minute)
}

func Test02(t *testing.T) {
	amount := 20
	array := make([]int, amount)
	for i := 0; i < amount; i++ {
		aa := i
		go func() {
			array[aa] = aa
			fmt.Println(array)
		}()

	}
	time.Sleep(time.Minute)
}

func Test03(t *testing.T) {
	amount := 20
	array := make([]int, amount)
	for i := 0; i < amount; i++ {
		aa := i
		fmt.Println(&aa, &i)
		go func() {
			array[aa] = aa
		}()

	}
	time.Sleep(time.Minute)
}
