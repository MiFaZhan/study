package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间，now的类型就是 time.Time
	now := time.Now()
	fmt.Printf("当前时间对象: %v\n", now)
	fmt.Printf("它的类型是: %T\n", now) // 输出: time.Time

	// 从时间对象中提取年、月、日等信息
	year, month, day := now.Date()
	fmt.Printf("日期: %d-%02d-%02d\n", year, month, day)
	fmt.Printf("星期: %s\n", now.Weekday())
	fmt.Printf("时间戳（秒）: %d\n", now.Unix())

	fmt.Println("\n格式化输出: " + now.Format("2006-01-02 15:04:05"))

	fmt.Println("\n格式化日期部分输出: " + now.Format("2006/01/02"))

	start := time.Now() // 记录开始时间点

	// 模拟程序运行，让程序暂停2秒
	time.Sleep(2 * time.Second)

	end := time.Now() // 记录结束时间点

	// 计算时间差，duration 的类型就是 time.Duration
	duration := end.Sub(start)
	fmt.Printf("\n程序执行耗时: %v\n", duration)
	fmt.Printf("具体秒数: %.2f 秒\n", duration.Seconds())

}
