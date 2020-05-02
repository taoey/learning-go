package main

import (
	"fmt"
	"reflect"
)

/*

切片 = Java list
*/

func main() {
	//初始化切片
	arr := make([]string, 0)
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr))

	//追加元素
	arr = append(arr, "Hi")
	arr = append(arr, "Go")
	arr = append(arr, "语言")
	fmt.Println(arr, len(arr))

	// int array demo
	intArr := make([]int, 0)
	for i := 0; i < 10; i++ {
		intArr = append(intArr, i)
	}
	fmt.Println(intArr, len(intArr))

	fmt.Println(intArr[1:3])

}
