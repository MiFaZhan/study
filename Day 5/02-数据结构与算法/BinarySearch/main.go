package main

import "fmt"

// 二分查找（迭代版本）：在有序数组中查找目标值
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2 // 防止溢出

		if arr[mid] == target {
			return mid // 找到目标，返回索引
		} else if arr[mid] < target {
			left = mid + 1 // 目标在右半部分
		} else {
			right = mid - 1 // 目标在左半部分
		}
	}
	return -1 // 未找到
}

// 二分查找（递归版本）
func binarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1 // 未找到
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, right)
	} else {
		return binarySearchRecursive(arr, target, left, mid-1)
	}
}

// 查找第一个出现的位置（处理重复元素）
func binarySearchFirst(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			right = mid - 1 // 继续在左半部分查找
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	fmt.Println("=== 二分查找示例 ===")

	// 注意：二分查找要求数组必须是有序的
	numbers := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// 示例1: 迭代版本
	target := 7
	result := binarySearch(numbers, target)
	if result != -1 {
		fmt.Printf("迭代版本: 在索引 %d 处找到 %d\n", result, target)
	} else {
		fmt.Printf("迭代版本: 未找到 %d\n", target)
	}

	// 示例2: 递归版本
	target2 := 15
	result2 := binarySearchRecursive(numbers, target2, 0, len(numbers)-1)
	if result2 != -1 {
		fmt.Printf("递归版本: 在索引 %d 处找到 %d\n", result2, target2)
	} else {
		fmt.Printf("递归版本: 未找到 %d\n", target2)
	}

	// 示例3: 查找不存在的元素
	target3 := 8
	result3 := binarySearch(numbers, target3)
	if result3 != -1 {
		fmt.Printf("在索引 %d 处找到 %d\n", result3, target3)
	} else {
		fmt.Printf("未找到 %d\n", target3)
	}

	// 示例4: 处理重复元素
	fmt.Println("\n=== 处理重复元素 ===")
	duplicates := []int{1, 2, 2, 2, 3, 4, 5, 5, 5, 6}
	target4 := 5
	result4 := binarySearchFirst(duplicates, target4)
	if result4 != -1 {
		fmt.Printf("第一个 %d 出现在索引 %d\n", target4, result4)
	}

	// 时间复杂度说明
	fmt.Println("\n=== 时间复杂度 ===")
	fmt.Println("最好情况: O(1) - 中间元素就是目标")
	fmt.Println("最坏情况: O(log n) - 需要不断二分")
	fmt.Println("平均情况: O(log n)")
	fmt.Println("\n前提条件: 数组必须是有序的！")

	// 查找过程演示
	fmt.Println("\n=== 查找过程演示 ===")
	demonstrateSearch(numbers, 13)
}

// 演示二分查找的过程
func demonstrateSearch(arr []int, target int) {
	fmt.Printf("在数组 %v 中查找 %d\n", arr, target)
	left := 0
	right := len(arr) - 1
	step := 1

	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("步骤 %d: left=%d, right=%d, mid=%d, arr[mid]=%d\n",
			step, left, right, mid, arr[mid])

		if arr[mid] == target {
			fmt.Printf("找到目标 %d 在索引 %d\n", target, mid)
			return
		} else if arr[mid] < target {
			fmt.Println("  目标在右半部分")
			left = mid + 1
		} else {
			fmt.Println("  目标在左半部分")
			right = mid - 1
		}
		step++
	}
	fmt.Println("未找到目标")
}
