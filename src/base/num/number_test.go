package num

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestInt(t *testing.T) {
	fmt.Println(math.Ceil(4.5), math.Floor(4.5))
	fmt.Println(math.Ceil(9/2), math.Floor(9/2))
	fmt.Println(math.Ceil(9.0/2.0), math.Floor(9.0/2.0))
	fmt.Println(math.Ceil(9/2.0), math.Floor(9/2.0))
}

//数据类型相互转化
func TestIntFloat(t *testing.T) {
	a := 1
	b := 2.6
	c := 2.0 / 3

	fmt.Println(reflect.TypeOf(a), a)
	fmt.Println(reflect.TypeOf(b), b)
	fmt.Println(reflect.TypeOf(c), c)

	// int -> float
	d := float64(a)
	fmt.Println("d", reflect.TypeOf(d), d)
	// float -> int
	f := int(b)
	fmt.Println("f", reflect.TypeOf(f), f) //统一向下取整舍弃小数部分
	// int -> string
	g := strconv.Itoa(a)
	fmt.Println("g", reflect.TypeOf(g), g)
	// float -> string

	h := strconv.FormatFloat(b, 'f', -1, 32) //2.6
	i := strconv.FormatFloat(b, 'e', -1, 32) //2.6
	fmt.Println("h", reflect.TypeOf(h), h)
	fmt.Println("i", reflect.TypeOf(i), i)
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)

	// string -> int
	j := "86"
	k, _ := strconv.Atoi(j)
	k2, _ := strconv.ParseInt(j, 10, 32)
	fmt.Println("k", reflect.TypeOf(k), k)
	fmt.Println("k2", reflect.TypeOf(k2), k2)

	// string -> float
	l := "16.6"
	m, _ := strconv.ParseFloat(l, 64)
	fmt.Println("m", reflect.TypeOf(m), m)
}
