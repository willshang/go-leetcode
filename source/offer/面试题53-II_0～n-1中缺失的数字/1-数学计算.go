package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum = sum - nums[i]
	}
	return sum
}
