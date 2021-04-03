package _go

import (
	"log"
	"testing"
)

// 使用go自带的log包
func init() {
	log.SetFlags(log.Ldate | log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[前缀测试] ")
}

func Test01(t *testing.T) {
	//[前缀测试] 2019/09/26 14:30:51 go_log_test.go:15: hello
	log.Println("hello")
}
