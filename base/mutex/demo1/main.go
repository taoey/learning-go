package main

import "sync"

func add(a, b int) {
	m := sync.Mutex{}
	m.Lock()
}

func main() {

}
