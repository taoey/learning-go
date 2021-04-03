package main

import "fmt"

type rect struct {
	width, height int
}

//方法直接操作对象内部数据  使用 （r *rect） 用来说明这个方法可以直接被rect类调用
func (r *rect) area() int {
	return r.height * r.width
}

func main() {
	r := rect{width: 10, height: 12}
	area := r.area()

	fmt.Println(area)
}
