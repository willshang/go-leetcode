package main

import "fmt"

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

// leetcode1437_是否所有1都至少相隔k个元素
func kLengthApart(nums []int, k int) bool {
	last := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if last != -1 && i-last <= k {
				return false
			}
			last = i
		}
	}
	return true
}
