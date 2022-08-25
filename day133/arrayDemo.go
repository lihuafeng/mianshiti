package main

import "fmt"

func main() {
	//a[3:4:4] 得到一个 slice，从原数组的第 4 个元素（索引是 3）开始，因此 t[0] 是 4。需要掌握 reslice 的语法。
	a := [5]int{1, 2, 3, 4, 5}
	t1 := a[3:4]
	fmt.Println(t1)
	t := a[3:4:4]
	fmt.Println(t)
	fmt.Println(t[0])
}
