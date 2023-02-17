package main

import "fmt"

func hello(num ...int) {
	fmt.Printf("num:%v \n", num)
	num[0] = 18
}
/**
答案解析：
参考答案及解析：A.18。知识点：可变参数函数。
 */
func main() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
	hello(1,2,3)
	fmt.Println(i[0])
}