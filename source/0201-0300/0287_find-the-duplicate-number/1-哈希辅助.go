package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > 0 {
			return nums[i]
		}
		m[nums[i]] = 1
	}
	return -1
}
