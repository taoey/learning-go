package main

// 切片的坑
// https://blog.csdn.net/vs974532452/article/details/101111671

// Findlinks1 prints the links in an HTML document read from standard input.
// leetcode 78. 子集

import "fmt"

func helper(nums []int, res [][]int, tmp []int, level int) {
	if len(tmp) <= len(nums) {
		//长度一样的tmp用的是同一个地址，故可能会存在覆盖值得情况，
		// 长度不一样时重新开辟空间，将已有得元素复制进去
		//*res = append(*res, tmp) 如此处，最终长度为1的tmp会被最后3这个元素全部覆盖
		//以下相当于每次重新申请内存，使其指向的地址不一样，解决了最后地址一样的元素值被覆盖的状态状态
		var a []int
		a = append(a, tmp[:]...)
		//res = append(res, a) 如果此处不是指针引用传递，在append后，res重新分配内存，与之前传进来的res地址不一样，最终res仍为空值
		res = append(res, a)
	}
	//fmt.Println(*res, "--->", tmp)
	for i := level; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		helper(nums, res, tmp, i+1)
		tmp = tmp[:len(tmp)-1] //相当于删除tmp末位的最后一个元素
	}
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var res [][]int
	var tmp []int
	helper(nums, res, tmp, 0)
	return res
}

func main() {
	pre := []int{1, 2, 3}
	fmt.Println(subsets(pre))
}
