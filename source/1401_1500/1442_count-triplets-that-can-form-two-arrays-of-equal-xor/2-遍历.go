package main

import "fmt"

func main() {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
}

// a[i]^...a[j-1]^a[j]^...a[k] = 0，则j可以取i+1、i+2、...、..k
func countTriplets(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		for k := i + 1; k < len(arr); k++ {
			temp = temp ^ arr[k]
			if temp == 0 {
				res = res + k - i
			}
		}
	}
	return res
}
