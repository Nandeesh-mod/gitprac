package main

import (
	"fmt"
)

func prime(n int) {

	for i := 0; i <= n; i++ {
		var flag int = 0
		if i == 0 || i == 1 {
			continue
		}
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Printf("%d ", i)
		}
	}

}
func main() {
	var n int = 100
	prime(n)
}
