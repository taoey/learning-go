package _1_reverse_list

import (
	"fmt"
	. "github.com/isdamir/gotype"
	"testing"
)

//实现链表的逆序 ： 就地逆序(不带头结点)
func Reverse01(head *LNode) *LNode {
	var pre *LNode
	cur := head
	var next *LNode

	for cur != nil {
		cur = cur.Next

		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

func Test01(t *testing.T) {
	head := &LNode{}
	fmt.Println("就地逆序")
	CreateNode(head, 8)
	PrintNode("逆序前：", head)
	head01 := Reverse01(head)
	PrintNode("逆序后：", head01)
}
