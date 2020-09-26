package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(fourSum([]int{0, 0}, 0))
	//fmt.Println(fourSum([]int{0, 1, 5, 0, 1, 5, 5, -4}, 11))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	m := make(map[[3]int]int)
	p := make(map[int]int)
	sort.Ints(nums)
	for k, v := range nums {
		p[v] = k
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if value, ok := p[target-sum]; ok && value > k {
					if _, ok2 := m[[3]int{nums[i], nums[j], nums[k]}]; !ok2 {
						res = append(res, []int{nums[i], nums[j], nums[k], target - nums[i] - nums[j] - nums[k]})
						m[[3]int{nums[i], nums[j], nums[k]}] = 1
					}
				}
			}
		}
	}
	return res
}
