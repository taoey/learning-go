/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/11 16:26
 */

package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	t1 int8  // 1
	t2 int64 // 8
	t3 int32 // 4
}

type TT struct {
	t1 int8  // 1
	t3 int32 // 4
	t2 int64 // 8
}

func main() {
	t := T{}
	fmt.Println(unsafe.Sizeof(t)) // 24

	tt := TT{}
	fmt.Println(unsafe.Sizeof(tt)) //16
}
