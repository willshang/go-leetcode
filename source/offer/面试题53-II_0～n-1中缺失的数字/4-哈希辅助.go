package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}

func missingNumber(nums []int) int {
	m := make(map[int]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	for i := 0; i <= len(nums); i++ {
		if m[i] == false {
			return i
		}
	}
	return 0
}
