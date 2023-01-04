package main

func main() {
	var a  int = 3
	// 以下有额外内存分配吗？
	var i interface{} = a


	//在 Go 中，接口被实现为一对指针（请参阅 Russ Cox 的 Go 数据结构：接口）：指向有关类型信息的指针和指向值的指针。可以简单的表示为：
	//
	//type iface struct {
	//    tab  *itab
	//    data unsafe.Pointer
	//}
	//其中 tab 是指向类型信息的指针；data 是指向值的指针。因此，一般来说接口意味着必须在堆中动态分配该值。
	//
	//然而，**Go 1.15 发行说明**在 runtime 部分中提到了一个有趣的改进：
	//
	//Converting a small integer value into an interface value no longer causes allocation.
	//
	//意思是说，将小整数转换为接口值不再需要进行内存分配。小整数是指 0 到 255 之间的数。
	//
	//我们实际简单测试一下。
	//
	//创建一个包 smallint，在包中创建文件 smallint.go，加上如下代码：
	//
	//package smallint
	//
	//func Convert(val int) []interface{} {
	//    var slice = make([]interface{}, 100)
	//    for i := 0; i < 100; i++ {
	//        slice[i] = val
	//    }
	//
	//    return slice
	//}
	//为了更好的看到效果，函数中进行了 100 次 int 到 interface 的转换。写个基准测试 smallint_test.go：
	//
	//package smallint_test
	//
	//import (
	//    "testing"
	//    "test/smallint"
	//)
	//
	//func BenchmarkConvert(b *testing.B) {
	//    for i := 0; i < b.N; i++ {
	//        result := smallint.Convert(12)
	//        _ = result
	//    }
	//}
}
