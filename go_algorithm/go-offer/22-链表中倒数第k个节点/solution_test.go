package _2_链表中倒数第k个节点

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针法
// [1,2,3,4,5]
func getKthFromEnd(head *ListNode, k int) *ListNode {
	left, right := head, head

	for right != nil {
		right = right.Next
		if k == 0 {
			left = left.Next
		} else {
			k--
		}
	}
	return left
}

func TestSolution(t *testing.T) {
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2, &c}
	a := ListNode{1, &b}

	end := getKthFromEnd(&a, 2)
	fmt.Println(end)
}
