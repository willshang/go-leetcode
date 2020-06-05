package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}

func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res ^ (i + 1) ^ nums[i]
	}
	return res
}
