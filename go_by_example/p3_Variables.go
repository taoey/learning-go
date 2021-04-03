package main

import "fmt"

////////////////////////
//	值只要声明了就要必须要使用，否则编译不通过
///////////////////////
func main() {
	var a string = "hello"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b + c)

	var d = true
	fmt.Println(d)

	f := "你好"
	fmt.Println(f)

}
