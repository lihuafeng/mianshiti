package main

import "fmt"

func increaseA() int {
  var i int
  defer func() {
    i++
  }()
  return i
}

func increaseB() (r int) {
  defer func() {
    r++
  }()
  return r
}

func main() {
  fmt.Println(increaseA())
  fmt.Println(increaseB())
}
/**
 输出 0 1
知识点：defer、返回值。注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。关于 defer 与返回值的知识点
defer 修饰的匿名函数，只能更新具名返回值
 */
