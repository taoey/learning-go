package slice

import (
	"fmt"
	"testing"
)

type Member struct {
	Ids []int
}

func Test01(t *testing.T) {
	m := Member{}

	m.Ids = append(m.Ids, 1)
	m.Ids = append(m.Ids, 2)
	m.Ids = append(m.Ids, 3)
	m.Ids = append(m.Ids, 7)

	a := append(m.Ids, 4)
	b := append(m.Ids, 5)
	c := append(m.Ids, 6)

	fmt.Println(a, b, c)

}

func Test02(t *testing.T) {
	m := Member{}

	a1 := append(m.Ids, 1)
	a2 := append(m.Ids, 2)
	a3 := append(m.Ids, 3)

	a := append(m.Ids, 4) // len == cap  此时m.ids再次添加新的元素时，不会进行扩容
	b := append(m.Ids, 5)
	c := append(m.Ids, 6)

	fmt.Println(a1, a2, a3)

	fmt.Println(a, b, c)
}

func Test03(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	arrSub := arr[3:]
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	arr[4] = 0
	fmt.Println(arr, arrSub, arrCopy)
}

func Test04(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}

	newArr := arr[1:2:3]
	newArr2 := arr[1:2:2]

	newArr = append(newArr, 7)
	newArr2 = append(newArr2, 7) // 创建新的底层数据

	arr[1] = 0
	fmt.Println(arr, newArr, newArr2) // [1 0 7 4 5 6] [0 7] [2 7]

}
