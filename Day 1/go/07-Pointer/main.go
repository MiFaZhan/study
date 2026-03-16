package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("\n变量a的地址: %x\n", &a)

	var p *int = &a
	fmt.Printf("指针p的地址: %x\n", &p)
	fmt.Printf("指针p指向的地址: %x\n", p)
	fmt.Printf("指针p指向的地址的值: %d\n", *p)

	//空指针
	var p1 *int
	//判断空指针
	if p1 == nil {
		fmt.Println("\np1是一个空指针")
		fmt.Printf("%x\n", p1)
	}

	//指针遍历数组
	var p2 []*int
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		p2 = append(p2, &arr[i])
	}
	for i := 0; i < len(p2); i++ {
		fmt.Printf("p2[%d] = %x, 指向的值: %d\n", i, p2[i], *p2[i])
	}

	//指向指针的指针
	var ptr *int = &a
	//访问指向指针的指针变量值需要使用两个 * 号
	var pptr **int = &ptr

	//打印指向指针的指针变量值
	fmt.Println("\n变量a的值:", a)
	fmt.Printf("变量a的地址: %x\n", &a)
	fmt.Printf("\n指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指针变量 ptr 的地址: %x\n", &ptr)
	fmt.Printf("\n指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Printf("指向指针的指针变量 pptr 的地址: %x\n", &pptr)

	b := 0
	c := 1
	fmt.Printf("\n交换前: b = %d, c = %d\n", b, c)
	swap(&b, &c)
	fmt.Printf("交换后: b = %d, c = %d\n", b, c)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
