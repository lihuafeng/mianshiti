package main

import "fmt"

func main() {
	a := make([]int, 0, 5)
	fmt.Println(a)
	addElem(a, 5)
	fmt.Println(a)
}

func addElem(a []int, i int) {
	a = append(a, 5)
}
