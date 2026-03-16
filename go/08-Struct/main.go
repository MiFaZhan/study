package main

import "fmt"

type Books struct {
	id     int
	titile string
	author string
}

func main() {
	fmt.Println(Books{1, "Go 语言", "张三"})
	fmt.Println(Books{id: 2, titile: "Java 语言", author: "李四"})
	fmt.Println(Books{id: 3, titile: "Python 语言"})

	var Book1 Books
	Book1.id = 1001
	Book1.titile = "活着"
	Book1.author = "余华"
	fmt.Println("\nBook1 Title:", Book1.titile)
	fmt.Println("Book1 Author:", Book1.author)
}
