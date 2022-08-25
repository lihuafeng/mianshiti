package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 0)
	<-ctx.Done()
	fmt.Println("timed out")

	//WithValue 底层是 valueCtx 结构体，其中 key、val 两个字段未导出，这里存放 "a" 和 "b"，同时还内嵌了 Context 接口。根据 Marshal 的规则，非导出的不会被序列化。而内嵌 Context 相当于导出了 Context 字段，而它的值是 context.Background()，即 background = new(emptyCtx)，emptyCtx 实际是 int 类型
	data, _ := json.Marshal(context.WithValue(context.Background(), "a", "b"))
	fmt.Println(string(data))
}