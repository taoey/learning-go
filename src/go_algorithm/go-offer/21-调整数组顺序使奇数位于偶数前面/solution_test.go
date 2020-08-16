package _1_调整数组顺序使奇数位于偶数前面

import (
	"fmt"
	"testing"
)

// 调整数组顺序使奇数位于偶数前面
// nums = [1,2,3,4]
// 答案之一：[1,3,2,4]

// 方法一
func exchange_00(nums []int) []int {
	a := []int{}
	b := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			a = append(a, num)
		} else {
			b = append(b, num)
		}
	}

	return append(b, a...)
}

// 方法二
func exchange(nums []int) []int {
	head := 0
	tail := len(nums) - 1

	for head < tail {
		for nums[head]%2 != 0 { // 前面是奇数
			if head < tail {
				head++
			} else {
				break
			}
		}
		for nums[tail]%2 == 0 { // 后面是偶数
			if head < tail {
				tail--
			} else {
				break
			}
		}
		nums[head], nums[tail] = nums[tail], nums[head]
	}

	return nums
}

func Test00(t *testing.T) {
	arr := []int{1, 3, 5, 7}
	arr2 := exchange(arr)
	fmt.Println(arr2)
}
