package main

import "fmt"

func main() {
	m := map[string]int{"uno": 1}
	p := &m
	fmt.Println(*p["uno"])
	fmt.Println(m["uno"])
}