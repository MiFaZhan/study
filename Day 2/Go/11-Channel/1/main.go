package main

import "fmt"

// c chan int 是 Go 语言中声明通道（channel）的一种方式
func sum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func main() {
	s := []int{-3, -2, -1, 0, 1, 2, 3}

	//创建一个通道，用于接收 sum 函数计算的结果
	c := make(chan int)
	//计算切片 s 的和
	go sum(s, c)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//从通道 c 中接收计算结果
	s1, s2, s3 := <-c, <-c, <-c
	fmt.Println(s1, s2, s3)
}

//通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
//ch := make(chan int, 100)
