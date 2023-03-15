/*
请完成以下任务：
    在此Pad中完成代码，不要将代码复制到外部IDE中。
    在完成面试后，请点击右下角的「End Interview」按钮。
    输入字符的规则：
        每行表示一条记录，字段之间以逗号（,）分隔。
        如果字段内容包含逗号（,），则该字段需要用双引号（"）包裹。
        如果字段内容包含双引号（"），则该字段需要用双引号包裹，并且字段内的每个双引号需要转义为两个双引号（""）。
    编写解析程序，将解析后的内容按行输出，字段之间以制表符（\t）分隔。
    例子：
        输入：Cindy,47,"收藏,""爬山",New Task
        输出：Cindy 47 收藏,"爬山 New Task
*/
package main

import (
  "fmt"
  "strings"
)

func main() {
  rows := `2,Tina,37,"足球,""篮球",Old Job
3,Alice Job,66,"""看电影"",旅游","上海,上海市"
4,John,44,"洗衣机101,""","LA""CITY"""
5,"Jane,li",55,Hiking,Canada`
  execute(rows)
}

func execute(rows string) {
  sli := strings.Split(rows, ",")

  fmt.Printf("%v \n", sli)

  //println(rows)
}
