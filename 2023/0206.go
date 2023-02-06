package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
/**
答案解析：
继承与多态的特点
在 golang 中对多态的特点体现从语法上并不是很明显。

我们知道发生多态的几个要素：

1、有interface接口，并且有接口定义的方法。

2、有子类去重写interface的接口。

3、有父类指针指向子类的具体对象

那么，满足上述 3 个条件，就可以产生多态效果，就是，父类指针可以调用子类的具体方法。

所以上述代码报错的地方在 var peo People = Student{} 这条语句， Student{} 已经重写了父类 People{} 中的 Speak(string) string 方法，那么只需要用父类指针指向子类对象即可。（Go 中不叫父类，这里是为了好理解）

所以应该改成 var peo People = &Student{} 即可编译通过。（People 为 interface 类型，就是指针类型）
 */
func main() {
	var peo People = Student{} // 错的
	//peo := &Student{} //对的
	think := "love"
	fmt.Println(peo.Speak(think))
}