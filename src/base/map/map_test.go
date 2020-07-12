package _map

import (
	"fmt"
	"testing"
)

// 使用map[string]struct{}来模拟 set
func Test01(t *testing.T) {
	set := make(map[string]struct{})

	result := set["hello"]
	fmt.Println(result)
}
