package main

import "fmt"

func main() {
	// fmt.Println(maxNonOverlapping([]int{2, 5, 7, 1, 6, 7}, 7))
	// fmt.Println(maxNonOverlapping([]int{-1, 3, 5, 1, 4, 2, -9}, 6))
	fmt.Println(maxNonOverlapping([]int{1, 2, 3, 4, 5, 6}, 7))
}

func maxNonOverlapping(nums []int, target int) int {
	res := 0
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := m[sum-target]; ok {
			m = make(map[int]int)
			sum = 0
			res++
		}
		m[sum] = i
	}
	return res
}
