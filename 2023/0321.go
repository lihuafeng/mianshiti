package main

import "fmt"

func f1() (r int) {
  defer func() {
    r++
  }()
  return 0
}


func f2() (r int) {
  t := 5
  defer func() {
    t = t + 5
  }()
  return t
}


func f3() (r int) {
  fmt.Printf("r1:%v \n", r)
  defer func(r int) {
    r = r + 5
    fmt.Printf("r2:%v \n", r)
  }(r)
  return 1
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}
