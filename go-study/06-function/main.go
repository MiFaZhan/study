package main

import "fmt"

/* 声明全局变量 */
var g int

func main() {
	a := 10
	b := 20
	g = Max(a, b)
	// 调用Max函数
	fmt.Println(g)

	// 调用swap函数
	var c string = "hello"
	var d string = "world"
	c, d = swap(c, d)
	fmt.Println("\n交换后的字符串:", c, d)
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func swap(a, b string) (string, string) {
	return b, a
}
