package point

import (
	"fmt"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test01(t *testing.T) {
	a := 1
	var b *int = &a // b 是指向a的地址的指针
	var c *int = &a // c 也是指向a的地址的指针

	fmt.Println(a, b, c)
	fmt.Println(reflect.TypeOf(b))
}

func Test02(t *testing.T) {
	//root 节点
	var result = ListNode{
		Val:  -1,
		Next: nil,
	}
	//赋值
	var currentResult *ListNode = &result

	//新节点
	var a = ListNode{
		Val:  1,
		Next: nil,
	}
	// 赋值
	//currentResult = &a  //改变了currentResult的指向
	*currentResult = a // 改变了 currentResult中的值
	fmt.Println(currentResult, result)
}
