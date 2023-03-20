package main

import "fmt"

type A interface {
  ShowA() int
}

type B interface {
  ShowB() int
}

type Work struct {
  i int
}

func (w Work) ShowA() int {
  return w.i + 10
}

func (w Work) ShowB() int {
  return w.i + 20
}

func main() {
  var a A = Work{3}
  s := a.(Work)
  fmt.Printf("%s, %T \n", s ,s)
  fmt.Println(s.ShowA()) //13
  fmt.Println(s.ShowB()) // 23
}
