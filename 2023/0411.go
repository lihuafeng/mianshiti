package main

import "fmt"

type Math struct {
  x, y int
}

var m = map[string]Math{
  "foo": Math{2, 3},
}

func main() {
  m["foo"].x = 4
  fmt.Println(m["foo"].x)
}
/**
答案解析：
参考答案及解析：B，编译报错 cannot assign to struct field m["foo"].x in map。
错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。

修改数据结构
var m = map[string]*Math{
	"foo": &Math{2, 3},
}
 */
