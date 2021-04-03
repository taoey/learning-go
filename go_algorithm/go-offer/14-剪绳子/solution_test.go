package _4_剪绳子

import (
	"fmt"
	"testing"
)

func cuttingRope(n int) int {

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1, n+1)

	// 子问题：可以剪，也可以不剪
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 4
	for i := 4; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			if j*dp[i-j] > curMax {
				curMax = j * dp[i-j]
			}
		}
		dp[i] = curMax
	}

	return dp[n]
}

func Test01(t *testing.T) {
	max := cuttingRope(9)
	fmt.Println(max)
}
