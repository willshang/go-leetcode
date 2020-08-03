package main

import (
	"fmt"
)

func main() {
	//fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex([]int{100}))
}

func hIndex(citations []int) int {
	left := 0
	right := len(citations) - 1
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] == len(citations)-mid {
			return len(citations) - mid
		} else if citations[mid] > len(citations)-mid {
			right = mid - 1
		} else if citations[mid] < len(citations)-mid {
			left = mid + 1
		}
	}
	return len(citations) - left
}
