package main

import "fmt"

func main() {
	//ch := make(chan int)
	ch := make(chan int, 2)

	// 无缓存区的通道会阻塞
	//ch 缓冲区大小为2，可以同时发送两个数据
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
