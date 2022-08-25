package main

import "fmt"

func main() {
	//编译错误 float64 不支持 | 操作。
	var a, b = 1.0, 2.0
	fmt.Println(a | b)

	var m, n float64 = 1.0, 4.0
	fmt.Println(m | n)
}
