package main

import "fmt"

func main() {
	//这是一个 bug，说正确答案应该是 []uint8{} []uint8 才对，因为 byte 是 uint8 的别名，但现在即使是 x := []uint8{}，结果也还是 []byte{} []uint8。不过为了兼容性，该 bug 目前不会修复，影响不大
	x := []byte{}
	fmt.Printf("%#v %T\n", x, x)
	//如果要将数字 65 转为字符串，不能使用 string(num)，如果使用这种方式转，得到的是一个 rune 的字符串表示，因为字面 A 的 ASCII 码是 65，因此这里输出结果是 A,string。
	num := 65
	str := string(num)
	fmt.Printf("%v, %T\n", str, str)
}
