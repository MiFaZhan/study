package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println(a.Name, "say Hello")
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	// 创建一个 Dog 实例
	d := Dog{
		Animal: Animal{Name: "金毛"},
		Breed:  "金毛",
	}
	d.Speak()
	fmt.Println("Breed:", d.Breed)

}
