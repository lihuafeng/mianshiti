package main

var n = -99

func main() {
	//注意，从 map 获取数据，即使 map 是 nil，也不会 panic。通过 make 创建 map 时，第 2 个参数可以为负数，以下是等价的：
	//make(map[string]int, -99)
	//make(map[string]int, 0)
	//make(map[string]int)
	//参考 https://github.com/golang/go/issues/46909
	m := make(map[string]int, n)
	println(m["Go"])
}
