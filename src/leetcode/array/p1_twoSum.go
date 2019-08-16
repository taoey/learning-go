package main

import "fmt"

//使用字典，一次遍历（注意result = append(result,intMap[another],index)，两个index的位置）
func twoSum(nums []int, target int) []int {
	var result []int
	intMap := make(map[int]int)

	for index, num := range nums {
		another := target - num
		if _, ok := intMap[another]; ok {
			result = append(result, intMap[another], index)
			return result
		} else {
			intMap[num] = index
		}
	}
	result = append(result, 0, 0)
	return result
}

func main() {
	sum := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(sum)
}
