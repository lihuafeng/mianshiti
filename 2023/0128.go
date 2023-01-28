package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Print(s)
	fmt.Printf("%p", s)
	s = append(s, 1,2,3)
	fmt.Print(s)
	fmt.Printf("%p", s)
}
