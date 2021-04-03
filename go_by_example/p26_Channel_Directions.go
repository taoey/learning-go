package main

import "fmt"

func poster(channel chan<- string, message string) {
	channel <- message
}

func main() {
	post := make(chan string, 1)
	poster(post, "succeed")

	fmt.Println(<-post)
}
