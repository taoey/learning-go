package _6_树的子结构

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1 找出A树中所有以B根节点的子节点
// 2 判断子树与B是否相同
// 特殊测试用例：B 空值判断

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	for _, item := range trave(A, B) {
		if isSub(item, B) {
			return true
		}
	}
	return false
}

func trave(A *TreeNode, B *TreeNode) []*TreeNode {
	result := []*TreeNode{}
	if A != nil {
		if A.Val == B.Val {
			result = append(result, A)
		}
		trave(A.Left, B)
		trave(A.Right, B)
	}
	return result
}

func isSub(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil {
		return false
	} else if A.Val != B.Val {
		return false
	}
	return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A == nil || B == nil {
		return false
	} else {
		if A.Val == B.Val {
			return isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
		}
	}
	return true
}
