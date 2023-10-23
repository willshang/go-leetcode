package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 4}))
}

func findDuplicate(nums []int) int {
	arrV := [32]int{}
	arrI := [32]int{}
	res := 0
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		index := i
		for j := 0; j < 32; j++ {
			if value&(1<<j) > 0 {
				arrV[j]++
			}
			if index > 0 && (index&(1<<j) > 0) {
				arrI[j]++
			}
		}
	}

	for i := 0; i < len(arrV); i++ {
		if arrV[i] > arrI[i] {
			res = res ^ (1 << i)
		}
	}
	return res
}
