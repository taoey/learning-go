package _7_build_tree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var dict map[int]int
var preorderGlobal []int

func buildTree(preorder []int, inorder []int) *TreeNode {
	preorderGlobal = preorder
	dict := map[int]int{}
	for index, item := range inorder {
		dict[item] = index
	}
	return helper(0, 0, len(inorder)-1)
}

func helper(preRootIndex, inLeftIndex, InRightIndex int) *TreeNode {
	if inLeftIndex > InRightIndex {
		return nil
	}
	inRootIndex := dict[preorderGlobal[preRootIndex]]
	root := TreeNode{
		Val:   preorderGlobal[preRootIndex],
		Left:  helper(preRootIndex+1, inLeftIndex, inRootIndex-1),
		Right: helper(preRootIndex+(inRootIndex-inLeftIndex+1), inRootIndex+1, InRightIndex),
	}
	return &root
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	treeRoot := buildTree(preorder, inorder)
	fmt.Println(treeRoot)
}
