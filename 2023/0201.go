package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	/**
	A. str := 'abc' + '123'
	B. str := "abc" + "123"
	C. str := '123' + "abc"
	D. fmt.Sprintf("abc%d", 123)

	参考答案：BD
	参考解析：考的知识点是字符串连接。除了以上两种连接方式，还有 strings.Join()、buffer.WriteString() 等。
	 */

	str1 := "abc" + "123"
	fmt.Printf("str1=%s \n", str1)

	str2 := fmt.Sprintf("abc%d",123)
	fmt.Printf("str2=%s \n", str2)

	str3 := strings.Join([]string{"abc","123"},"")
	fmt.Printf("str3=%s \n", str3)

	var str4 bytes.Buffer
	str4.WriteString("abc")
	str4.WriteString("123")
	fmt.Printf("str4=%s \n", str4.String())
}
