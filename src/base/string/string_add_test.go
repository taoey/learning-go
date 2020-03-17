package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// 最简单的方法：直接使用 +
func TestAddString00(t *testing.T) {
	var s string
	str := "hello"
	str += "-"
	str += "world"
	fmt.Println(s)
	fmt.Println(str)
}

// 一种比较常用的方法： fmt.Sprintf
func TestAddString01(t *testing.T) {
	string1 := "hello"
	string2 := "world"
	result := fmt.Sprintf("%s-%s", string1, string2)
	fmt.Println(result)
}

// 对于 string 数组：strings.Join
func TestAddString02(t *testing.T) {
	strList := []string{"hello", "world", "taoey"}
	join := strings.Join(strList, "-")
	fmt.Println(join)
}

// 频繁进行拼接，推荐使用：bytes.Buffer  注意:WriteString()参数一定是string，不能是byte
func TestAddString03(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("hello")
	buf.WriteString("-")
	buf.WriteString("world")
	result := buf.String()

	fmt.Println(result)
	buf.Reset()
}
