package main

import "fmt"

type writer interface {
	myWrite()
}

func main3() {

	//显式地将 nil 赋值给接口时，接口的 type 和 data 都将为 nil。此时，接口与 nil 值判断是相等的。
	var wri writer = nil
	if wri == nil {
		fmt.Println("wri == nil")
	} else {
		fmt.Println("wri != nil")
	}
}
