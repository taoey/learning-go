package main

import "fmt"

// add 函数
func Add(a, b int) int {
	return a + b
}

// 多返回值函数
func vals() (int, int) {
	return 1, 3
}

func sumString(strs ...string) string {
	sum := ""
	for _, str := range strs {
		sum += str
	}
	return sum
}

func main() {
	sum := Add(1, 3)
	fmt.Println(sum)

	a, b := vals()
	fmt.Println(a, b)

	str := sumString("1", "2", "3")
	fmt.Println(str)
}
