package _1_转数组的最小数字

import (
	"fmt"
	"testing"
)

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)>>1
		if numbers[mid] > numbers[right] { //mid 取到了左侧数组中的数据，需要左侧边界需要右移
			left = mid + 1
		} else if numbers[mid] < numbers[right] { // mid 取到了右侧边界中的数据，右侧边界需要替换
			right = mid
			//right = mid-1 不这么写是因为 mid可能为旋点
		} else {
			right--
		}
	}
	return numbers[left]
}

// 二分解法

func TestMinArray(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	num := minArray(numbers)
	fmt.Println(num)

}
