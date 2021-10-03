package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// 剑指OfferII007.数组中和为0的三个数
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	m := make(map[[2]int]int)
	p := make(map[int]int)
	sort.Ints(nums)
	for k, v := range nums {
		p[v] = k
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j]
			if sum > 0 {
				break
			}
			if value, ok := p[-sum]; ok && value > j {
				if _, ok2 := m[[2]int{nums[i], nums[j]}]; !ok2 {
					res = append(res, []int{nums[i], nums[j], 0 - nums[i] - nums[j]})
					m[[2]int{nums[i], nums[j]}] = 1
				}
			}
		}
	}
	return res
}
