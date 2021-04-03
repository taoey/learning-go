package main

import "fmt"

func foo(num int) {
	fmt.Printf("%d", num)
}

func main() {
	num := 10
	foo(num)
}
