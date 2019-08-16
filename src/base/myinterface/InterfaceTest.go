package main

import "fmt"

// 定义一个A接口
type Aer interface {
	Aa()
	Ab()
}

// 定义一个B接口
type Ber interface {
	Aa()
	Ab(str string)
}

// 定义一个C结构体（类）
type C struct {
	id int
}

//直接使用相同的方法名 当实现了接口的全部方法时，即实现了该接口
func (c C) Aa() {
	fmt.Println("Aa")
}

func (c C) Ab(str string) {
	fmt.Println("Ab")
}

func main() {

}
