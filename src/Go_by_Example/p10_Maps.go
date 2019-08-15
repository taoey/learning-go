package main

import "fmt"

/*
	主要完成map的增删改查
*/
func main() {
	testMap := make(map[string]int)

	//添加元素
	testMap["k1"] = 1
	testMap["k2"] = 2
	fmt.Println(testMap)

	//删除元素
	delete(testMap, "k1")
	fmt.Println(testMap)

	//修改元素
	testMap["k2"] = 3
	fmt.Println(testMap)

	//获取元素
	fmt.Println(testMap["k1"]) // 获取不存在的键，取得value的默认值
	fmt.Println(testMap["k2"])
}
