package main

import (
	"fmt"
)

type obj struct{}

func main() {
	a := &obj{}
	fmt.Printf("%p\n", a)
	c := &obj{}
	fmt.Printf("%p\n", c)
	fmt.Println(a == c)
}
