package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{3, 4, 3, 3}))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			return nums[i]
		}
	}
	return 0
}
