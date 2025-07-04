//Go programm for finding lowest number
package main



import "fmt"

func find_lowest(myarr []int) int {
	var low int = myarr[0]
	for i := 0; i < len(myarr); i++ {
		if myarr[i] < low {
			low = myarr[i]
		}
	}

	return low

}

func main() {
	var size int
	fmt.Scanf("%d", &size)
	var myarr []int = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &myarr[i])
	}

	fmt.Printf("smallest element is : %d", find_lowest(myarr))

}
