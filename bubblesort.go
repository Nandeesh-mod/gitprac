package main

import (
	"fmt"
)

// optimised bubble sort
func bubble_sort(arr []int) {
	var length int = len(arr)
	total_iteration := 0

	for i := 0; i < length-1; i++ {
		total_iteration = (total_iteration + 1) + (length - 1 - i)
		swap_count := 0
		for j := 0; j < length-1-i; j++ {
			if arr[j] < arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swap_count++
			}
		}
		if swap_count == 0 {
			break
		}
	}
	fmt.Printf("Total iteration is : %d", total_iteration)
}

func main() {
	var n int
	fmt.Scan(&n)
	my_arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&my_arr[i])
	}
	bubble_sort(my_arr)
	fmt.Println(my_arr)
}
