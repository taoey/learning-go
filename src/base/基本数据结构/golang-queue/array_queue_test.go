package golang_queue

import (
	"fmt"
	"testing"
	"time"
)

func TestArrayQueue_Remove(t *testing.T) {
	queue := new(ArrayQueue)
	queue.Add("hello")
	queue.Add("world")
	fmt.Println(queue)

	queue.Remove()
	fmt.Println(queue)
}

func TestArrayQueue_Element(t *testing.T) {

}

func TestArrayQueue_Add(t *testing.T) {
	queue := new(ArrayQueue)
	queue.Add("hello")
	queue.Add("world")

	fmt.Println(queue)

}

func TestGo(t *testing.T) {
	amount := 20
	queue := new(ArrayQueue)

	for i := 0; i < amount; i++ {
		go func() {
			aa := i
			queue.Add(aa)
			fmt.Println("add", queue)
		}()
	}
	//for i:=0;i<amount/2;i++{
	//	go func() {
	//		queue.Remove()
	//		fmt.Println("remove",queue)
	//	}()
	//}
	time.Sleep(time.Minute)
}
