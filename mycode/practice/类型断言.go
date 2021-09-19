package main

import (
	"fmt"
)

func main5() {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Print(value, ",", ok)
}
