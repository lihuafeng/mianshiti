package main

import "fmt"

func main() {
	f := func() { fmt.Print("A") }
	defer f()
	f = func() { fmt.Print("B") }
	defer f()
	f = func() { fmt.Print("C") }
	defer f()
}
