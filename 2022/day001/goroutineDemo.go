package main

//参考文章链接 https://studygolang.com/articles/23333#reply12
import (
	"fmt"
	"runtime"
	"sync"
)
//func main() {
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func() {
//			fmt.Println(i)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}

func main() {
	//WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup //
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}