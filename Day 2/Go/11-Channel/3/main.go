package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	//创建一个通道，用于接收 fibonacci 函数计算的结果
	c := make(chan int, 10)
	//遍历通道
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
