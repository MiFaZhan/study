package main

/*
Go 语言的切片（Slice）主要用于以下场景：

动态数组： 切片是动态数组的抽象，长度可变，可以灵活地追加元素，非常适合需要动态调整数组大小的场景。
数组切片： 可以从数组中创建切片，实现对数组部分的引用和操作，而无需复制整个数组。
数据集处理： 切片可以用于处理各种数据集，例如列表、队列、栈等，简化数据操作。
函数参数： 切片可以作为函数参数传递，实现数组的“传值”功能，避免数组复制带来的性能损耗。
并发安全： 切片是引用类型，可以安全地在多个协程之间共享，支持并发访问。
*/

import "fmt"

func main() {
	//一、分步声明和初始化
	//var numbers []int
	//numbers = make([]int, 3, 5)

	//二、使用 := 赋值形式
	//var numbers := make([]int, 3, 5)

	var numbers = make([]int, 3, 5)

	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300

	printSlice(numbers)

	numbers1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("\nnumbers1[1:4]:", numbers1[1:4])
	fmt.Println("numbers1[:4]:", numbers1[:4])
	fmt.Println("numbers1[4:]:", numbers1[4:])
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
