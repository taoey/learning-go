package main

import (
	"fmt"
	"reflect"
)

type People struct{}

func (p *People) Eat(food string) {
	fmt.Println("eat:", food)
}

func (p *People) Sleep() {
	fmt.Println("sleep")
}

func main() {
	people := People{}

	// 调用无参函数
	f := reflect.ValueOf(&people).MethodByName("Sleep")
	f.Call([]reflect.Value{})

	// 调用有参函数
	f2 := reflect.ValueOf(&people).MethodByName("Eat")
	v := make([]reflect.Value, 1) // 参数个数
	v[0] = reflect.ValueOf("饺子")
	f2.Call(v)
}
