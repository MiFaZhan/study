package main

import "fmt"

func quickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	pivot := arr[n/2]
	left, right := make([]int, 0, n/2), make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		if i == n/2 {
			continue
		}
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{6, 2, 5, 12, 3, 1}
	arr = quickSort(arr)
	fmt.Println(arr)
}
