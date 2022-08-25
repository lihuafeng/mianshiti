package main

import "fmt"

func main() {
	//编译错误 float64 不支持 | 操作。
	var a, b = 1.0, 2.0
	fmt.Println(a | b)
}
