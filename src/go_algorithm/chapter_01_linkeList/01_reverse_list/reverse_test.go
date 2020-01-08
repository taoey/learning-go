package _1_reverse_list

import (
	"fmt"
	. "github.com/isdamir/gotype"
	"testing"
)

//实现链表的逆序 ： 就地逆序

func Reverse01(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode
	var cur *LNode
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func Test01(t *testing.T) {
	head := &LNode{}
	fmt.Println("就地逆序")
	CreateNode(head, 8)
	PrintNode("逆序前：", head)
	Reverse01(head)
	PrintNode("逆序后：", head)
}
