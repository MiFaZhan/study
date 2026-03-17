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
	// 实现接口方法
	fmt.Println(d.name, "正在吃饭")
}

// 接口组合
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() string {
	return "Reading data"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}

func main() {
	// 创建一个 dog 实例
	d := dog{name: "旺财", age: 3}
	// 调用 returnAge 方法
	fmt.Printf("%s的年龄:%d\n", d.name, d.age)
	// 调用 eat 方法
	d.eat()

	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello, Go!")
}
