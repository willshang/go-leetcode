package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	//fmt.Println(subsetsWithDup([]int{1, 4, 4, 4, 4}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	res = append(res, []int{})
	start := 1
	for i := 0; i < len(nums); i++ {
		temp := make([][]int, len(res))
		for j := 0; j < len(res); j++ {
			if i > 0 && nums[i] == nums[i-1] && j < start {
				continue
			}
			value = append(value, nums[i])
			temp[key] = append(temp[key], value...)
		}
		start = len(res)
		fmt.Println(temp, start)
		for _, v := range temp {
			res = append(res, v)
		}
	}
	return res
}
