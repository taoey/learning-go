package _2_矩阵中的路径

import "testing"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, word string, k int) bool {
	// 出口判断
	// 遍历的值和当前的字符不相等
	if board[i][j] != word[k] {
		return false
	}
	// 遍历的值为word最后一个字符
	if k == len(word)-1 {
		return true
	}

	// 遍历过的路径置为空
	temp := board[i][j]
	board[i][j] = byte(' ')

	// board[i][j]==word[k] && k != len(word)-1 说明已经匹配到了字符串，但未匹配完全

	// 上
	if i-1 >= 0 && dfs(board, i-1, j, word, k+1) {
		return true
	}
	// 下
	if i+1 < len(board) && dfs(board, i+1, j, word, k+1) {
		return true
	}
	// 左
	if j-1 >= 0 && dfs(board, i, j-1, word, k+1) {
		return true
	}
	// 右
	if j+1 < len(board[0]) && dfs(board, i, j+1, word, k+1) {
		return true
	}
	board[i][j] = temp
	return false
}

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	exist(board, "ABCCED")
}
