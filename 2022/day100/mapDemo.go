package main

import "fmt"

func main() {
	//因为 map 中元素顺序是随机的，因此结果不确定，每次运行结果可能不一样。
	m := map[string]int{"foo": 0, "bar": 1, "baz": 2}
	for k := range m {
		if k == "foo" {
			delete(m, "bar")
		}
		if k == "bar" {
			delete(m, "foo")
		}
	}
	fmt.Println(m)
}
