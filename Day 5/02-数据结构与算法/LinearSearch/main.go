package main

import "fmt"

// 线性查找：从头到尾遍历数组，逐个比较元素
func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i // 返回找到的索引
		}
	}
	return -1 // 未找到返回-1
}

// 线性查找字符串数组
func linearSearchString(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("=== 线性查找示例 ===")

	// 示例1: 查找整数
	numbers := []int{10, 23, 45, 70, 11, 15}
	target := 70
	result := linearSearch(numbers, target)
	if result != -1 {
		fmt.Printf("在索引 %d 处找到 %d\n", result, target)
	} else {
		fmt.Printf("未找到 %d\n", target)
	}

	// 示例2: 查找不存在的元素
	target2 := 100
	result2 := linearSearch(numbers, target2)
	if result2 != -1 {
		fmt.Printf("在索引 %d 处找到 %d\n", result2, target2)
	} else {
		fmt.Printf("未找到 %d\n", target2)
	}

	// 示例3: 查找字符串
	fmt.Println("\n=== 字符串查找 ===")
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	targetName := "Charlie"
	result3 := linearSearchString(names, targetName)
	if result3 != -1 {
		fmt.Printf("在索引 %d 处找到 %s\n", result3, targetName)
	} else {
		fmt.Printf("未找到 %s\n", targetName)
	}

	// 时间复杂度说明
	fmt.Println("\n=== 时间复杂度 ===")
	fmt.Println("最好情况: O(1) - 第一个元素就是目标")
	fmt.Println("最坏情况: O(n) - 目标在最后或不存在")
	fmt.Println("平均情况: O(n) - 需要遍历一半的元素")
}
