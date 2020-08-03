package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	nums = append(nums, nums[0])
	i, j := 0, 1
	index := 0
	res = append(res, strconv.Itoa(nums[i]))
	for j = 1; j < len(nums); j++ {
		if nums[j]-nums[j-1] != 1 {
			if j-i > 1 {
				str := strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j-1])
				res[index] = str
			}
			res = append(res, strconv.Itoa(nums[j]))
			i = j
			index++
		}
	}
	return res[:index]
}
