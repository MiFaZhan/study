package main

import "fmt"

func main() {
	a := [...][]int{
		{1, 2, 3},
		{4, 5, 6, 7},
	}

	//遍历数组
	for i, v := range a {
		for j, vv := range v {
			fmt.Printf("a[%d][%d] = %d\n", i, j, vv)
		}
		fmt.Println()
	}

	// 遍历数组（无索引）
	for _, v := range a {
		for _, vv := range v {
			fmt.Printf("%v ", vv)
		}
		fmt.Println()
	}

	fmt.Println()
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}

}
