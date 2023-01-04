package main

import (
	"fmt"
	"math"
)

func main() {
	//首先，math.Log(-1)，这是高中数学，负数不能取对数，所以这里的结果是 NaN，即不是一个数。这是关于计算机浮点数的知识了。
	//
	//关键知识点：任何一个 NaN 都不相等。因此 m[v] 会查找不到，所以返回 int 的默认值 0，而 m 中三个元素，key 都是 NaN，因此不相等，所有 len(m) 等于 3。
	v := math.Log(-1)
	fmt.Println(v)
	m := map[float64]int{v: 1, v: 2, v: 3}
	fmt.Printf("%T %v\n", m, m)
	fmt.Println(m[v], len(m))
}
