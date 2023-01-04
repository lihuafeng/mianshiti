package main

import "fmt"

func main()  {
	//149
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(m)
	fmt.Println(len(m))
}
