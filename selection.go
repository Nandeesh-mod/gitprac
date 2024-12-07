package main

import (
	"fmt"
)

func selection_sort(arr []int) {
	var length int = len(arr)

	for i := 0; i < length-1; i++ {
		min_index := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		min_ele := arr[min_index]
		arr[min_index] = arr[i]
		arr[i] = min_ele
	}

}

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	selection_sort(arr)
	fmt.Println(arr)
}
