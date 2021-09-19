package main

import "fmt"

// 一个服务器需要满足能够开启和写日志的功能
type Service interface {
	Start()
	Log(string)
}

// 日志器
type Logger struct {
}

// 实现Service的Log方法
func (g *Logger) Log(l string) {
	fmt.Println(l)

}

// 游戏服务
type GameService struct {
	Logger //嵌入日志器
}

// 实现Service的start方法
func (g *GameService) Start() {

}

func main2() {
	var s Service = new(GameService)
	s.Start()
	s.Log("Hello")
}
