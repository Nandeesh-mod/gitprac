package main

import "fmt"

func insertion_sort(arr []int) {
	var length int = len(arr)

	for i := 1; i < length; i++ {
		var ele int = arr[i]
		j := i - 1
		for j >= 0 && ele < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = ele
	}
}

func main() {
	var myarr []int = []int{50, -4, 39, -200, 121, 11}
	insertion_sort(myarr)
	fmt.Println(myarr)
}
