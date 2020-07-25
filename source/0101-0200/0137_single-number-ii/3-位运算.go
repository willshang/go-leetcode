package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, -3, 2}))
}

func singleNumber(nums []int) int {
	var res int
	for i := 0; i < 64; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if (nums[j]>>i)&1 == 1 {
				count++
			}
		}
		res |= count % 3 << i
	}
	return res
}
