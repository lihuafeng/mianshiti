package main

import "C"

func main() {
	//不会panic,如果把第 3 行注释，会 panic。原因是 cgo 会使死锁检查失效。
	var ch chan struct{}
	<-ch
}
