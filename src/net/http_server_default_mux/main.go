/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/7 10:24
 */
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// 创建路由规则
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	// 进行监听
	http.ListenAndServe(":8090", nil)
}
