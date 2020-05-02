package main

import "fmt"

/*
 range 用于遍历各种数据集合
*/
func main() {
	//遍历数组
	fmt.Println("数组遍历")
	myArray := [3]int{1, 2, 3}
	for index, num := range myArray {
		fmt.Println(index, num)
	}

	for _, num := range myArray {
		fmt.Println(num)
	}

	//遍历Slice
	fmt.Println("切片遍历")
	intList := make([]int, 0)
	for i := 0; i < 5; i++ {
		intList = append(intList, i)
	}
	fmt.Println(intList)
	for index, num := range intList {
		fmt.Println(index, num)
	}

	//遍历map
	stringMap := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	for k, v := range stringMap {
		fmt.Println(k, v)
	}
	fmt.Println(stringMap)

	//遍历字符串  获取的是unicode 编码
	fmt.Println("字符串遍历")
	myStr := "huang"
	for index, str := range myStr {
		fmt.Println(index, str)
	}
}
