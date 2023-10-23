package main

import "fmt"

func main() {
	arr := []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}
