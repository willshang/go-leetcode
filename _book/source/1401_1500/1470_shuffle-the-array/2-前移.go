package main

import (
	"fmt"
)

func main() {
	//fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4))
}

func shuffle(nums []int, n int) []int {
	for i := n; i < 2*n; i++ {
		temp := i
		for j := 0; j < 2*n-1-i; j++ {
			nums[temp], nums[temp-1] = nums[temp-1], nums[temp]
			temp--
		}
	}
	return nums
}
