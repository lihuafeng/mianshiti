package main

import "fmt"

func main()  {
	//0 开头，表明是八进制，但八进制最大的数字是 7，因此编译不通过。
	//fmt.Println(09)

	//关键在于 3/2 计算的结果，3、2 这是整型字面值常量。根据 Go 的规则，3/2 结果也是整型，因此是 1，最后会隐式转换为 float64。
	var i float64 = 3.0 / 2.0
	var j float64 = 3 / 2
	fmt.Println(i)
	fmt.Println(j)
}
