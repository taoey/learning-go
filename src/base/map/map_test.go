package _map

import (
	"fmt"
	"testing"
)

// 使用map[string]bool来模拟 set
func Test01(t *testing.T) {
	set := make(map[string]bool)

	result := set["hello"]
	fmt.Println(result)
}