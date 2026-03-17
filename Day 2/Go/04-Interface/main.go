package main

import "fmt"

type animal interface {
	// 定义一个接口方法
	eat()
}

type dog struct {
	name string
	age  int
}

func (d dog) eat() {
	fmt.Println(d.name, "正在吃")
}

func (d dog) sleep() {
	fmt.Println(d.name, "正在睡")
}

func main() {
	// 创建一个 dog 实例
	d := dog{name: "旺财", age: 3}
	// 调用 eat 方法
	d.eat()
	// 调用 sleep 方法
	d.sleep()
}
