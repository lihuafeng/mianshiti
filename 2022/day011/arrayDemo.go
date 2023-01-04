package main

import "fmt"

func main()  {
	//145
	//数组是值类型，长度相等，每个元素也相等，数组就相等
	type pos [2]int
	a := pos{4, 5}
	b := pos{4, 5}
	fmt.Println(a == b)
}
