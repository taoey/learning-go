package _map

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"testing"
	"unsafe"
)

// 使用map[string]struct{}来模拟 set
func Test01(t *testing.T) {
	set := make(map[string]struct{})

	result := set["hello"]
	fmt.Println(result)
}

func Test02(t *testing.T) {
	myMap := sync.Map{}

	myMap.Store("hello", struct{}{})
}

func Test03(t *testing.T) {
	var s struct{}
	var s1 struct{}
	fmt.Println("空结构体占用内存的情况：", unsafe.Sizeof(s))
	fmt.Printf("空结构体指针指向情况:s = %p, s1 = %p,两个指针的比较结果：%v", &s, &s1, &s == &s1)
}

// 底层map指向同一个内存地址
func Test04(t *testing.T) {
	mapA := bson.M{}
	mapB := mapA

	mapA["tao"] = 18
	fmt.Println(mapA, mapB)                //map[tao:18] map[tao:18]
	fmt.Printf("针指向情况:%p, %p", mapA, mapB) //针指向情况:0xc00006c6f0, 0xc00006c6f0
}

// 使用序列化和反序列化实现map的深度拷贝
func Test05(t *testing.T) {
	mapA := bson.M{}
	mapB := bson.M{}
	bytes, _ := json.Marshal(mapA)
	json.Unmarshal(bytes, &mapB)

	mapB["tao"] = 18

	fmt.Println(mapA, mapB)                //map[] map[tao:18]
	fmt.Printf("针指向情况:%p, %p", mapA, mapB) //针指向情况:0xc00006c6f0, 0xc00006c720
}
