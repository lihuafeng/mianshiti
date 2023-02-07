package main

import (
	"fmt"
)

type People1 interface {
	Show()
}

type Student1 struct{}

func (stu *Student1) Show() {

}

func live() People1 {
	var stu *Student1
	return stu
}

/*
以下代码打印出来什么内容，说出为什么。
答案解析：
结果

BBBBBBB
分析：

我们需要了解interface的内部结构，才能理解这个题目的含义。（源码基于 Go1.17）

interface 在使用的过程中，共有两种表现形式

一种为空接口(empty interface)，定义如下：

var MyInterface interface{}
另一种为非空接口(non-empty interface), 定义如下：

type MyInterface interface {
	function()
}
这两种 interface 类型在底层分别用两种struct表示，空接口为eface, 非空接口为iface。
空接口 eface
空接口 eface 结构，由两个属性构成，一个是类型信息 _type，一个是数据信息。其数据结构声明如下：

type eface struct {      // 空接口
    _type *_type         // 类型信息
    data  unsafe.Pointer // 指向数据的指针(go 语言中特殊的指针类型 unsafe.Pointer 类似于 c 语言中的void*)
}
_type 属性：是 Go 语言中所有类型的公共描述，Go 语言几乎所有的数据结构都可以抽象成 _type，是所有类型的公共描述，_type 负责决定 data 应该如何解释和操作， _type 的结构如下：

type _type struct {
	size       uintptr // 类型大小
	ptrdata    uintptr // 前缀持有所有指针的内存大小
	hash       uint32  // 数据 hash 值
	tflag      tflag
	align      uint8   // 对齐
	fieldalign uint8   // 嵌入结构体时的对齐
	kind       uint8   // kind 有些枚举值 kind 等于 0 是无效的
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
data 属性： 表示指向具体的实例数据的指针，它是一个unsafe.Pointer类型，相当于一个 C 的万能指针void*。
非空接口 iface
iface 表示 non-empty interface 的数据结构，非空接口初始化的过程就是初始化一个 iface 类型的结构，其中data的作用与 eface 的相同，这里不再多加描述。

type iface struct {
  tab  *itab
  data unsafe.Pointer
}
iface 结构中最重要的是 itab 结构（结构如下），每一个 itab 都占 32 字节的空间。itab 可以理解为pair<interface type, concrete type> 。itab 里面包含了 interface 的一些关键信息，比如 method 的具体实现。

type itab struct {
  inter  *interfacetype   // 接口自身的元信息
  _type  *_type           // 具体类型的元信息
  hash   int32            // _type 里也有一个同样的 hash，此处多放一个是为了方便运行接口断言
  _      [4]byte
  fun    [1]uintptr       // 函数指针，指向具体类型所实现的方法
}
其中值得注意的字段，个人理解如下：

interface type包含了一些关于 interface 本身的信息，比如package path，包含的method。这里的interfacetype 是定义 interface 的一种抽象表示。
_type表示具体化的类型，与 eface 的 _type 类型相同。
hash字段其实是对_type.hash的拷贝，它会在 interface 的实例化时，用于快速判断目标类型和接口中的类型是否一致。另，Go 的 interface 的 Duck-typing 机制也是依赖这个字段来实现。
fun字段其实是一个动态大小的数组，虽然声明时是固定大小为 1，但在使用时会直接通过 fun 指针获取其中的数据，并且不会检查数组的边界，所以该数组中保存的元素数量是不确定的。
所以，People 拥有一个 Show 方法，属于非空接口，People 的内部定义是一个iface结构体

type People interface {
    Show()
}

func live() People {
    var stu *Student
    return stu
}
stu 是一个指向 nil 的空指针，但是最后return stu 会触发匿名变量 People = stu 值拷贝动作，所以最后live()放回给上层的是一个People insterface{}类型，也就是一个iface struct{}类型。 stu 为 nil，只是iface中的 data 为 nil 而已。 但是iface struct{}本身并不为 nil.
 */
func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB") //正确的输出
	}
}
