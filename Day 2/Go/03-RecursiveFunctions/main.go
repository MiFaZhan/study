package main

import "fmt"

func main() {
	fmt.Println("\n5的阶乘:", factorial(5))

	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	// 递归调用
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
