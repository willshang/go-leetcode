package main

import "sort"

func main() {

}

func intersection(nums [][]int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			m[nums[i][j]]++
		}
	}
	res := make([]int, 0)
	for k, v := range m {
		if v == len(nums) {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}
