package main

import "fmt"

type integer int

func (i integer) String() string {
	return "hello"
}

func main() {
	fmt.Println(integer(5))
}
