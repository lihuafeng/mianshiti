package main

import (
	"fmt"
	"net/url"
)

// 其中 url.Values 的定义：type Values map[string][]string
type Query struct {
	url.Values
}

func main() {
	q := Query{}
	q.Values = make(map[string][]string) //映射的初值是nil 必须实例化后赋值
	q.Values["name"] = []string{"polarisxu"}
	fmt.Println(q.Get("name"))
}