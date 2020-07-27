package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, -3, 2}))
}

func singleNumber(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a = a ^ nums[i]&(^b)
		b = b ^ nums[i]&(^a)
	}
	return a
}
