package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}

func singleNumbers(nums []int) []int {
	res := make([]int, 2)
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = temp ^ nums[i]
	}
	index := firstBit(temp)
	for i := 0; i < len(nums); i++ {
		if isBit(nums[i], index) {
			res[0] = res[0] ^ nums[i]
		} else {
			res[1] = res[1] ^ nums[i]
		}
	}
	return res
}

func firstBit(num int) int {
	res := 0
	for num&1 == 0 {
		res++
		num = num >> 1
	}
	return res
}

func isBit(num int, index int) bool {
	num = num >> index
	return num&1 == 1
}
