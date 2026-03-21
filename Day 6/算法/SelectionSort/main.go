package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := []int{6, 2, 5, 12, 3, 1}
	selectionSort(arr)
	fmt.Println(arr)
}
