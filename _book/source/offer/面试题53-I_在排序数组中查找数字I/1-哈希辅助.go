package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			m[target]++
		}
	}
	return m[target]
}
