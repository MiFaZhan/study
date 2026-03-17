package main

import "fmt"

func printType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("int类型:", v)
	case string:
		fmt.Println("string类型:", v)
	default:
		fmt.Println("未知类型")
	}
}
func main() {
	// 测试 printType 函数
	printType(10)
	printType("hello")
	printType(true)
}
