package main

import (
  "fmt"
  "time"
)

func main() {

  var m = [...]int{1, 2, 3}

  for i, v := range m {
    go func() {
      fmt.Println(i, v)
    }()
  }

  //fix
  //for ii, vv :=range m{
  //  go func(ii,vv int) {
  //    fmt.Println(ii, vv)
  //  }(ii,vv)
  //}

  //for iii,vvv :=range m{
  //  mm := iii
  //  nn := vvv
  //  go func() {
  //    fmt.Println(mm,nn)
  //  }()
  //}

  time.Sleep(time.Second * 3)
}
