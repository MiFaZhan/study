package main

import "fmt"

func main() {
	// 定义整型变量
	var a int
	a = 100
	fmt.Println("\n整型变量:", a)

	// 定义字符串变量
	b := "字符串变量"
	fmt.Println("\n字符串变量:", b)

	// 格式化输出
	//%d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	fmt.Printf("\n格式化输出:%s\n", fmt.Sprintf(url, stockcode, enddate))

	// 定义整型数组
	var c []int
	c = []int{1, 2, 3}
	fmt.Println("\n整型数组:", c)

	var d int
	var e float64
	var f bool
	var g string
	fmt.Printf("\n默认值:%v %v %v %q\n", d, e, f, g)

	//定义常量
	const name = "米法展"
	fmt.Println("\n常量:", name)

	//特殊常量
	const (
		h = iota
		i = iota
		j = "特殊常量"
		k = 3 << iota
	)

	fmt.Println("\n特殊常量:", h, i, j, k)

	//浮点数
	var pi float64 = 3.14
	fmt.Println("\n浮点数:", pi)

	//类型转换
	fmt.Printf("\n浮点数转换为整型:%v %T\n", int(pi), int(pi))
}
