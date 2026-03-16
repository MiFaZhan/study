package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("\n变量a的地址: %x\n", &a)

	var p *int = &a
	fmt.Printf("指针p的地址: %x\n", &p)
	fmt.Printf("指针p指向的地址: %x\n", p)
	fmt.Printf("指针p指向的地址的值: %d\n", *p)
}
