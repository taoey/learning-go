package main

import "fmt"

// 值传递和指针
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1

	zeroval(1)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)
}
