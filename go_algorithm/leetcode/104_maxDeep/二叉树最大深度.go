package _04_maxDeep

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// dfs 队列
	queue := []*TreeNode{}
	deep := 0
	queue = append(queue, root)

	for len(queue) > 0 {
		// 每次获取一层的结点，deep数量+1
		deep += 1
		len := len(queue) // 注意需要提前获取，否则在运行过程中queue的数量会发生变化
		for i := 0; i < len; i++ {
			top := queue[0]
			queue = queue[1:]

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}

	}
	return deep
}
