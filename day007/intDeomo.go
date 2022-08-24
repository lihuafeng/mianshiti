package main

import "fmt"

func main()  {
	//0 开头，表明是八进制，但八进制最大的数字是 7，因此编译不通过。
	fmt.Println(09)
}
