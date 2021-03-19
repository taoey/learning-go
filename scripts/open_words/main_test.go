package main

import (
	"fmt"
	"testing"
)

// 创建记录
func TestCreateWorkList(t *testing.T) {
	CreateWorkList("result_data")
}

// 读取记录
func TestReadWorkList(t *testing.T) {
	list := ReadWorkList("result_data")
	fmt.Println(list)
}

func Test00(t *testing.T) {
	var sarr []string
	sarr = []string{"1"}

	prArr := &sarr

	fmt.Println((*prArr)[0])
}
