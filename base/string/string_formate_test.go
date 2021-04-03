package string

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	str := fmt.Sprintf("%s,%s,%d", "1", "2", 3)
	fmt.Println(str)
}
