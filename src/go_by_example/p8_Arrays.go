package main

import "fmt"

func main() {
	//数组声明未初始化，初始值均为0
	var arr [4]int
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	//数组声明并初始化,未初始化的值为0
	arr_2 := [5]int{1, 2, 3, 4}
	fmt.Println(arr_2)
}
