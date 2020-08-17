package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)

	if k > n {
		k = k % n
	}
	if k == 0 || k == n {
		return
	}
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			nums[j], last = last, nums[j]
		}
	}
}
