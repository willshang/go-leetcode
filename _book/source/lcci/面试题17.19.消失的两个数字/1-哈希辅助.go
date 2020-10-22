package main

import "fmt"

func main() {
	fmt.Println(missingTwo([]int{2, 3}))
}

func missingTwo(nums []int) []int {
	res := make([]int, 0)
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for i := 1; i <= len(nums)+2; i++ {
		if m[i] == false {
			res = append(res, i)
		}
	}
	return res
}
