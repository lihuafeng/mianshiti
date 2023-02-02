package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {
	/**
	结果

	编译失败，cannot assign to struct field list["student"].Name in map

	分析

	map[string]Student 的 value 是一个 Student 结构值，所以当list["student"] = student,是一个值拷贝过程。
	而list["student"]则是一个值引用。那么值引用的特点是只读。所以对list["student"].Name = "LDB"的修改是不允许的。
	 */

	//map
	list = make(map[string]Student)
	//切片
	sli_list := []Student{}

	student := Student{"Aceld"}

	list["student"] = student
	//list["student"].Name = "LDB"
	sli_list = append(sli_list, student)

	fmt.Println(list["student"])
	fmt.Printf("sli_list=%v \n", sli_list)
	fmt.Printf("student=%v \n", student)

	student.Name = "Stanfu"
	fmt.Println(list["student"])
	fmt.Printf("sli_list=%v \n", sli_list)
	fmt.Printf("student=%v \n", student)
}