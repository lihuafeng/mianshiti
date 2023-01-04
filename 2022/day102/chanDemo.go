package main

import "fmt"

func main() {
	//After the last value has been received from a closed channel c, any receive from c will succeed without blocking, returning the zero value for the channel element.
	//
	//即从关闭了的 channel 接收不会堵塞，并返回零值。
	c := make(chan int)
	close(c)
	val, _ := <-c
	fmt.Println(val)
}
