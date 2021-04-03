package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine") //异步调用

	var input string
	fmt.Scanln(&input) //阻塞线程，为了防止协程未被调用程序就执行完毕
	fmt.Println("done")

}
