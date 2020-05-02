package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("first for end ")

	for i := 2; i <= 4; i++ {
		fmt.Println(i)
	}

	count := 0
	fmt.Println("while begin")
	for {
		if count == 2 {
			break
		}
		fmt.Println(count)
		count++
	}
}
