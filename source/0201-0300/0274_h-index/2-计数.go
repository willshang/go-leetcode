package main

import (
	"fmt"
)

func main() {
	//fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{0, 0}))
}

func hIndex(citations []int) int {
	arr := make([]int, len(citations)+1)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations) {
			arr[len(citations)]++
		} else {
			arr[citations[i]]++
		}
	}
	count := 0
	for i := len(citations); i >= 0; i-- {
		count = count + arr[i]
		if count >= i {
			return i
		}
	}
	return 0
}
