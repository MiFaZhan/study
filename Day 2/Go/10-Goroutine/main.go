package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
		//time.Second 是 Go 标准库 time 包中预定义的时间常量，它表示 10⁹ 纳秒（10亿纳秒）= 1秒
		time.Sleep(time.Second)
	}
}

func main() {
	//输出的 Main 和 Hello。输出是没有固定先后顺序，因为它们是两个 goroutine 在执行
	go sayHello()
	for i := 0; i < 5; i++ {
		fmt.Println("main:", i)
		time.Sleep(time.Second)
	}
}

/*
Go 标准库 time 包源码（简化版）
const (
    Nanosecond  time.Duration = 1
    Microsecond              = 1000 * Nanosecond
    Millisecond              = 1000 * Microsecond
    Second                   = 1000 * Millisecond  // = 1e9 纳秒
    Minute                   = 60 * Second
    Hour                     = 60 * Minute
)
*/
