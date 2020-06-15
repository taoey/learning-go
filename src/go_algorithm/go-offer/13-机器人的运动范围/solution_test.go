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

	for queue.Len() != 0 {
		item := queue.Front()
		queue.Remove(queue.Front())

		x := item.Value.(*Point).x
		y := item.Value.(*Point).y

		arr[x][y] = 2

		// 上
		if x-1 >= 0 && arr[x-1][y] == 1 {
			queue.PushBack(&Point{x - 1, y})
			arr[x-1][y] = 2
			count++
		}
		// 下
		if x+1 < len(arr) && arr[x+1][y] == 1 {
			queue.PushBack(&Point{x + 1, y})
			arr[x+1][y] = 2
			count++
		}
		// 左
		if y-1 >= 0 && arr[x][y-1] == 1 {
			queue.PushBack(&Point{x, y - 1})
			arr[x][y-1] = 2
			count++
		}
		// 右
		if y+1 < len(arr[0]) && arr[x][y+1] == 1 {
			queue.PushBack(&Point{x, y + 1})
			arr[x][y+1] = 2
			count++
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
	movingCount(3, 2, 17)
}
