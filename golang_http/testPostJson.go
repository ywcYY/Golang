package main

import (
	"io"
	"log"
	"net/http"
)

func testHttpServer() {
	// 请求处理函数
	f := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "hello world")
	}
	// 响应路径,注意前面要有斜杠 /
	http.HandleFunc("/hello", f)
	// 设置监听端口，并监听，注意前面要有冒号:
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	testHttpServer()
}
