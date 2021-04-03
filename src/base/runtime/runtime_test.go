package runtime

import (
	"fmt"
	"runtime"
	"testing"
)

// 获取cpu数量
func Test01(t *testing.T) {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
}
