package main

import "fmt"

/*
[1 2 3 1 5 6 7] 4
[1 2 3 1 5 6 4] 7
[1 2 7 1 5 6 4] 3
[1 2 7 1 5 3 4] 6
[1 6 7 1 5 3 4] 2
[1 6 7 1 2 3 4] 5
[5 6 7 1 2 3 4] 1
[5 6 7 1 2 3 4]
*/
func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate(nums, 2)
	fmt.Println(nums)
}

// leetcode 189 旋转数组
func rotate(nums []int, k int) {
	n := len(nums)

	if k > n {
		k = k % n
	}
	if k == 0 || k == n {
		return
	}
	count := 0

	for i := 0; count < len(nums); i++ {
		current := i
		prev := nums[i]
		for {
			next := (current + k) % len(nums)
			nums[next], prev = prev, nums[next]
			current = next
			// fmt.Println(nums, prev)
			count++
			if i == current {
				break
			}
		}
	}
}
