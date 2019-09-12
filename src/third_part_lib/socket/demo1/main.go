package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	//监听msg事件，打印消息，并触发客户端replay事件
	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	//监听关闭事件
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		//s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("E:\\projects\\go\\src\\Lets-GO\\src\\third_part_lib\\socket\\demo1\\asset")))
	log.Println("Serving at localhost:8100...")
	log.Fatal(http.ListenAndServe(":8100", nil))
}
