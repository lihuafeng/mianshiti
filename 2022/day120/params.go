package main

import "fmt"

func f(a ...int) {
	//首先，a 的类型是 []int，调用 f 时，没有传递任何参数，因此相当于值是 nil，即 a 的类型是 []int，值是 nil。而 fmt.Printf 的动词 %#v 会同时打印类型和值
	fmt.Printf("%#v\n", a)
}

func main() {
	f()
}
