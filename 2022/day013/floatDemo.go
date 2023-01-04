package main

import "fmt"

func main()  {
	//143
	var a = 0.0
	const b = 0.0
	fmt.Println(a / b)
	//142
	//Slice, map, and function values are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier nil.
	fmt.Println(func() {} == func() {})
}
