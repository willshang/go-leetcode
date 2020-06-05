package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := m[target-b]; ok {
			return []int{nums[j], nums[i]}
		}
		m[b] = i
	}
	return nil
}
