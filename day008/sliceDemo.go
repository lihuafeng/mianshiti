package main

import "fmt"

func main()  {
	//150
	//interface类型变量与非interface类型变量判等时，首先要求非interface类型实现了接口，否则编译不通过（本题接口方法集为空，我们认为所有类型都实现了该接口）
	//满足上一条的前提下，interface类型变量的动态类型、值均与非interface类型变量相同时，两个变量判等结果为true，结合array判等规则，答案为true

	var p [100]int
	var m interface{} = [...]int{99: 0}
	fmt.Println(p)
	fmt.Println(m)
	fmt.Println(p == m)

	//array(slice) 初始化时也是可以指定索引的。该题有一个索引 3，根据数组的特点，必然有小于 3 的所有也存在，上题中，v 的值是：[0, 2, 0, 4]。
	v := [...]int{1: 2, 3: 4}
	fmt.Println(len(v))
}
