package main

import (
	"fmt"
	"runtime"
)

func main() {
	sys := runtime.GOOS
	fmt.Println(sys)
}
