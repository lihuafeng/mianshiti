package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//在 main goroutine 第二次请求 Lock 时，会堵塞。这样另一个 goroutine 会运行，释放锁。如果没有另一个 goroutine 释放锁，则会报 fatal error，所有的 goroutine 都处于 sleep 状态，死锁！
	var m sync.Mutex
	fmt.Print("A, ")
	m.Lock()

	go func() {
		time.Sleep(200 * time.Millisecond)
		m.Unlock()
	}()

	m.Lock()
	fmt.Print("B ")
}
