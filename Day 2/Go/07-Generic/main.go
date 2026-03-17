package main

import "fmt"

// 定义一个泛型函数，用于打印任意类型的切片
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// 定义一个泛型结构体
type Container[T any] struct {
	value T
}

// 为泛型结构体定义方法
func (c *Container[T]) SetValue(v T) {
	c.value = v
}

func (c *Container[T]) GetValue() T {
	return c.value
}

// 定义一个带有类型约束的泛型函数，只接受数字类型
func AddNumbers[T int | float64](a, b T) T {
	return a + b
}

func main() {
	// 使用泛型函数打印整数切片
	ints := []int{1, 2, 3, 4, 5}
	fmt.Print("Integers: ")
	PrintSlice(ints)

	// 使用泛型函数打印字符串切片
	strings := []string{"Hello", "World", "Go", "Generics"}
	fmt.Print("Strings: ")
	PrintSlice(strings)

	// 使用泛型结构体
	intContainer := Container[int]{value: 100}
	fmt.Println("Container int value:", intContainer.GetValue())

	strContainer := Container[string]{}
	strContainer.SetValue("Generic Container")
	fmt.Println("Container string value:", strContainer.GetValue())

	// 使用带约束的泛型函数
	sumInt := AddNumbers(10, 20)
	sumFloat := AddNumbers(10.5, 20.3)
	fmt.Printf("Sum of ints: %d, Sum of floats: %.1f\n", sumInt, sumFloat)
}
