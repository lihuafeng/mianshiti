package main

import (
	"fmt"
	"sync"
)

func main() {
	//因为 hello 的参数是 sync.WaitGroup，这会导致 main 中 wg 被复制一份，起不到该有的作用，应该改为 *sync.WaitGroup。
	wg := sync.WaitGroup{}
	wg.Add(1)
	go hello(wg)
	wg.Wait()
	fmt.Println("world")
}

func hello(wg sync.WaitGroup) {
	fmt.Println("hello")
	wg.Done()
}
