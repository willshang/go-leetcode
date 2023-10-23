package main

import "fmt"

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

func kLengthApart(nums []int, k int) bool {
	last := -(k + 1) // 兼容第0个元素是1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if i-last <= k {
				return false
			}
			last = i
		}
	}
	return true
}
