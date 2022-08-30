package main

import (
	"fmt"
	"reflect"
)

func main() {
	//用来判断两个值是否深度一致：除了类型相同；在可以时（主要是基本类型）会使用==；但还会比较array、slice的成员，map的键值对，结构体字段进行深入比对。map的键值对，对键只使用==，但值会继续往深层比对。
	// DeepEqual函数可以正确处理循环的类型。函数类型只有都会nil时才相等；空切片不等于nil切片；还会考虑array、slice的长度、map键值对数。
	i := func() {}
	j := func() {}
	fmt.Println(i ==nil)
	fmt.Println(j ==nil)
	no1 := &i
	no2 := &j
	if reflect.DeepEqual(no1, no2) {
		fmt.Println("equal")
		return
	}
	fmt.Println("not equal")
}
