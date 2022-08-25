package main

import "fmt"

func main() {
	//% 运算只能用于 整数类型。1 % 2.0，两个操作数都是字面量常量，都是无类型的，这时会以 2.0 的 untype float constant 为准，1 隐式转为 untype float constant，所以编译错误。
	//
	//而 int(1) % 2.0 中，2.0 是无类型的，int(1) 是 int，因此 2.0 会转为 int，因此能正常编译。
	fmt.Println(1 % 2.0)
	fmt.Println(int(1) % 2.0)
}
