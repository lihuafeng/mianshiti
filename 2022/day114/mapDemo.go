package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Println(key, val)
		m[key] = &val
	}
	fmt.Printf("%T\n",m)
	fmt.Printf("%v\n",m)
	fmt.Println(*m[0])
	fmt.Println(*m[1])
	fmt.Println(*m[2])
	fmt.Println(*m[3])
}
