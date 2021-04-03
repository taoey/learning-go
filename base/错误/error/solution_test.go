package error

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	//test()    //defer panic
	ret := divisionIntRecover(1, 0)
	fmt.Println(ret)
	fmt.Println("main")
}

func divisionIntRecover(a, b int) (ret int) {
	defer func() {
		if err := recover(); err != nil {
			// 打印异常，关闭资源，退出此函数
			fmt.Println(err)
			ret = -1
		}
	}()
	return a / b
}
