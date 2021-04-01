package _5_合并两个排序的链表

import "testing"

func TestSolution(t *testing.T) {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用递归解法虽然比较抽象，但是逻辑还是非常简单的
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

}
