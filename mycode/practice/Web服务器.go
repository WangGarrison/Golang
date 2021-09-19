package main

import (
	"io/ioutil"
	"net/http"
)

func main9() {
	//注册处理请求的函数，第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求
	http.HandleFunc("/index", index)

	//开始监听：http.ListenAndServe监听，参数是服务器要监听的主机地址和端口号
	http.ListenAndServe("localhost:8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}
