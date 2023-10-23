package main

import "fmt"

func main() {
	arr := []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}

// 剑指OfferII069.山峰数组的顶部
func peakIndexInMountainArray(A []int) int {
	left, right := 0, len(A)-1
	for {
		mid := left + (right-left)/2
		if A[mid] > A[mid+1] && A[mid] > A[mid-1] {
			return mid
		}
		if A[mid] > A[mid-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
}
