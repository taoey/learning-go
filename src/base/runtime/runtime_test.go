package runtime

import (
	"fmt"
	"runtime"
	"testing"
)

func Test01(t *testing.T) {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
}
