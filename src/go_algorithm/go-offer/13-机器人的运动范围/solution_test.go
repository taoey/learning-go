package _3_机器人的运动范围

import (
	"container/list"
	"fmt"
	"testing"
)

func movingCount(m int, n int, k int) int {
	// 生成一个二维数组（golang的二维数组是真麻烦）
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}

	// 遍历全图，找出能够走的点 置为1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if getSum(i)+getSum(j) <= k {
				arr[i][j] = 1
			}
		}
	}
	// 使用 bfs进行广度优先遍历，找出可走路径，统计路径上点的个数
	queue := list.New()
	queue.PushBack(&Point{0, 0})
	count := 1

	// 方向数组 上下左右
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for queue.Len() != 0 {
		item := queue.Front()
		queue.Remove(queue.Front())

		x := item.Value.(*Point).x
		y := item.Value.(*Point).y

		arr[x][y] = 2

		//上下左右四个点判断
		for i := 0; i < 4; i++ {
			x = item.Value.(*Point).x + dx[i]
			y = item.Value.(*Point).y + dy[i]

			//越界判断
			if x >= 0 && x < len(arr) && y >= 0 && y < len(arr[0]) && arr[x][y] == 1 {
				queue.PushBack(&Point{x, y})
				arr[x][y] = 2
				count++
			}
		}

	}

	return count
}

//坐标点
type Point struct {
	x int
	y int
}

// 获得数位之和
func getSum(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func TestGetSum(t *testing.T) {
	sum := getSum(20)
	fmt.Println(sum)
}

func TestSolution(t *testing.T) {
	count := movingCount(3, 2, 17)
	fmt.Println(count)
}
