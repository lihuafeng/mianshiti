package main

import (
	"fmt"
	//"time"
)

func main() {
	// close channel 后，<-ch 会立马返回
	//ch := make(chan bool)
	//go func() {
	//	<-ch
	//	fmt.Print("Goroutine")
	//}()
	//time.Sleep(2 * time.Second)
	//close(ch)
	//time.Sleep(3 * time.Second)
	//fmt.Print("Main")

	c := make(chan int, 5)
	c <- 5
	c <- 6
	c <- 7
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
}
