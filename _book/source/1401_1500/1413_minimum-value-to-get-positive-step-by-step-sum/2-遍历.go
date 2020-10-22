package main

import "fmt"

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
	fmt.Println(minStartValue([]int{1, 2}))
}

func minStartValue(nums []int) int {
	res := 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum+res < 1 {
			res = 1 - sum
		}
	}
	return res
}
