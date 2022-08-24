package main

import "fmt"

func main() {
	//real 是一个内建函数，虽然返回的类型是 FloatType，但这里返回的是常量（见 Go 规范）。这里要提醒大家，并非 slice 的索引可以是 float 类型。
	//常量值由 rune、 integer、 floating-point、 imaginary或 string字面量、表示常量的标识符、常量表达式、结果为常量的转换或某些内置的结果值表示函数，例如 unsafe.Sizeof应用于任何值， cap或len应用于 某些表达式， real以及imag应用于复数常量和complex应用于数值常量。布尔真值由预先声明的常量 true和表示false。预先声明的标识符 iota表示整数常数。
	//参考文章 https://docs.studygolang.com/ref/spec#Constants
	a := []int{7, 8, 9}
	fmt.Println(a[real(2)])
}