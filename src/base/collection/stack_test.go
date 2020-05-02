package collection

import (
	"container/list"
	"fmt"
	"testing"
)

// 遍历list
func printStack(root *list.List) {
	var next *list.Element
	for e := root.Front(); e != nil; e = next {
		fmt.Print(e.Value)
		next = e.Next()
	}
	fmt.Println("\n---------")
}

// 清除链表 有个小坑，详见：https://blog.csdn.net/bluuus/article/details/82774410
func clearList(root *list.List) {
	var next *list.Element
	for cur := root.Front(); cur != nil; cur = next {
		next = cur.Next()
		root.Remove(cur)
	}
}

func Test00(t *testing.T) {
	stack := list.New() //链表创建

	stack.PushBack("a")
	stack.PushBack("b")
	stack.PushBack("c")
	stack.PushBack("d") //向尾部添加元素

	printStack(stack)

	stack.Remove(stack.Front()) // 移除头部元素
	stack.Remove(stack.Back())  // 移除尾部元素
	printStack(stack)

	fmt.Println(stack.Front().Value) //获取首部元素内容
	fmt.Println(stack.Back().Value)  // 获取尾部元素内容
	printStack(stack)

	clearList(stack)
	printStack(stack)

}
