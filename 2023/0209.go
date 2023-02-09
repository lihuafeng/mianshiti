package main
//题一
func main() {
	var a int8 = -1
	var b int8 = -128 / a

	println(b)
}



/**
//题二
package main

func main() {
    const a int8 = -1
    var b int8 = -128 / a

    println(b)
}
 */

/**
答案解析：
这是站长在去年双节期间公众号「polarisxu」上发布的题目。

答对的人真不多（半数以上答错了），特别是题一，一半以上竟然是 128，难道不知道 int8 能表示的范围吗？[-128, 127]。不过为什么答案是：题一 -128，题二编译错误？

其实这是一道计算机基础题。

先看看网友 Jayce 的解释：第一题是 -128（untyped const）/ -1 (int8 var)，untyped 隐式转换为 int8，刚好在范围内，结果是 128 ，溢出 int8 的范围。因为结果不是常量，允许溢出，最高位为符号位，变成了补码，刚好又是 -128。 第二题 -128 和 -1 都是 const，直接在编译时求值，untyped 的 -128 隐式转 int8，结果为 128，仍然是一个 const。const 转换时不允许溢出，编译错误。 其实差别就是表达式的值，题一不是常量题二是，常量类型转换不允许溢出后 truncate。
 */