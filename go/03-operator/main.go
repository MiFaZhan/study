package main

import "fmt"

func main() {
	a := 1
	b := 2
	var c int = a + b
	var d int = a - b
	var e int = a * b
	var f int = a / b
	var g int = a % b

	fmt.Println(c, d, e, f, g)

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
}
