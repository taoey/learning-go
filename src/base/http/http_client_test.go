package main

import (
	"fmt"
	"net/http"
	"testing"
)

// 一个简单的HttpClient-get请求
func TestHttpClientGet(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com")
	fmt.Println(resp, err)
}
