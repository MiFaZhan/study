package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
			if !swapped {
				break
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Print("原始数组:")
	fmt.Println(arr)

	bubbleSort(arr)

	fmt.Print("排序后的数组:")
	fmt.Println(arr)
}

/*
算法步骤:

比较相邻元素：从列表的第一个元素开始，比较相邻的两个元素。

交换位置：如果前一个元素比后一个元素大，则交换它们的位置。

重复遍历：对列表中的每一对相邻元素重复上述步骤，直到列表的末尾。这样，最大的元素会被"冒泡"到列表的最后。

缩小范围：忽略已经排序好的最后一个元素，重复上述步骤，直到整个列表排序完成。
*/
