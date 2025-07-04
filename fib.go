// this is the  programm for fibbonacci sequnce 
// done by dev 2
package main

import (
	"fmt"
)

// finding fib by using loops print all the fib number

func fibLoops(n int) {
	var first int = 0
	var second int = 1

	fmt.Printf("%d %d ", first, second)
	for i := 2; i < n; i++ {
		fib := first + second
		fmt.Printf("fib : %d ", fib)
		first = second
		second = fib
	}

}

// using recursion

func fibn(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		n_fib := fibn(n-1) + fibn(n-2)
		return n_fib
	}
}

func main() {
	var n int = 10
	for i := 0; i < n; i++ {
		fmt.Printf(" %d ", fibn(i))
	}
}
