package main

import (
	"fmt"
	"math"
)

// text2 元素做行，text1元素 做列 len(text1) 个 text2
func longestCommonSubsequence(text1 string, text2 string) int {
	row := len(text2) + 1
	col := len(text1) + 1

	arr := make([][]int, row)
	for i := range arr {
		arr[i] = make([]int, col)
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if text2[i-1] == text1[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else {
				arr[i][j] = int(math.Max(float64(arr[i][j-1]), float64(arr[i-1][j])))
			}
		}
	}
	return arr[row-1][col-1]
}

func makArr(row, column, initNum int) [][]int {
	arr := make([][]int, row)
	for i := range arr {
		arr[i] = make([]int, column)
	}

	if initNum == 0 {
		return arr
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			arr[i][j] = initNum
		}
	}
	return arr
}

func main() {
	res := longestCommonSubsequence("bsbininm", "jmjkbkjkv")
	fmt.Println(res)
}
