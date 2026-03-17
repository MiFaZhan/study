package main

import "fmt"

func main() {
	var i interface{} = "hello"
	// 断言 i 为 string 类型
	s := i.(string)
	fmt.Println(s)

	// 断言 i 为 int 类型
	// 会触发 panic 异常
	//fmt.Println(i.(int))

	// 使用类型断言的安全方式
	v, ok := i.(int)
	// 如果 ok 为 true，则断言成功
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("断言失败")
	}
}
