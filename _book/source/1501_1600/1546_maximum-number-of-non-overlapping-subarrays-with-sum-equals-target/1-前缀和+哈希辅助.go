package main

import "fmt"

func main() {
	// fmt.Println(maxNonOverlapping([]int{2, 5, 7, 1, 6, 7}, 7))
	// fmt.Println(maxNonOverlapping([]int{-1, 3, 5, 1, 4, 2, -9}, 6))
	fmt.Println(maxNonOverlapping([]int{1, 2, 3, 4, 5, 6}, 7))
}

// leetcode1546_和为目标值的最大数目不重叠非空子数组数目
func maxNonOverlapping(nums []int, target int) int {
	res := 0
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	prev := -1
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if index, ok := m[sum-target]; ok && index >= prev {
			res++
			prev = i
		}
		m[sum] = i
	}
	return res
}
