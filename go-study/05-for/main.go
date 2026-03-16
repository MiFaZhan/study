package main

import "fmt"

func main() {
	sum1 := 0

	for i := 0; i <= 10; i++ {
		sum1 += i
	}
	fmt.Println("0-10 累加和:", sum1)

	sum2 := 1
	for sum2 <= 10 {
		sum2 += sum2
		if sum2 > 10 {
			break // 满足条件时退出循环
		}
	}
	fmt.Println("翻倍后结果:", sum2)

	a := []int{1, 2, 3, 4, 5}
	fmt.Println("\n--- 遍历切片 ---")
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	fmt.Println("\n--- 使用 continue 跳过 i==2 ---")
	for i := 0; i < len(a); i++ {
		if i == 2 {
			continue
		}
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	fmt.Println("\n--- 使用break 当i==3时跳出循环  ---")
	for i := 0; i < len(a); i++ {
		if i == 3 {
			break
		}
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	fmt.Println("\n--- 使用 goto 跳过 b==15  ---")
	/* 定义局部变量 */
	var b int = 10

	/* 循环 */
LOOP:
	for b < 20 {
		if b == 15 {
			/* 跳过迭代 */
			b = b + 1
			goto LOOP
		}
		fmt.Printf("b的值为 : %d\n", b)
		b++
	}
}
