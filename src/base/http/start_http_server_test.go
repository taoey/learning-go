package main

import (
	"fmt"
	"net/http"
	"testing"
)

// 操作response
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// 获取header内容，并写入response中
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}


// http.HandleFunc 方法将路由和处理方法相对应 ，本例子使用的是http.ListenAndServe方法进行的启动
func Test00(t *testing.T) {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8100",nil)
}



// DemoHandle server handle示例
type DemoHandle struct {

}
// ServeHTTP 匹配到路由后执行的方法
func (DemoHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("执行了DemoHandle处理方法")
}

func Test01(t *testing.T) {
	// 初始化一个http.Server
	server := &http.Server{}
	// 初始化handler并赋值给server.Handler
	server.Handler = DemoHandle{}
	// 绑定地址
	server.Addr = ":8888"

	// 所有路由都会执行ServeHTTP
	server.ListenAndServe()
}

