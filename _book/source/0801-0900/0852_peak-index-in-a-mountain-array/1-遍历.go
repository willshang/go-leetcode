package main

import "fmt"

func main() {
	arr := []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(A []int) int {
	n := len(A)
	for i := 0; i < n-1; i++ {
		if A[i] > A[i+1] {
			return i
		}
	}
	return 0
}
