package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

// leetcode228_汇总区间
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[j]-nums[j-1] != 1 {
			str := ""
			if j-i > 1 {
				str = strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j-1])
			} else {
				str = strconv.Itoa(nums[i])
			}
			res = append(res, str)
			i = j
		}
		j++
	}
	if j == len(nums) {
		str := ""
		if j-i > 1 {
			str = strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j-1])
		} else {
			str = strconv.Itoa(nums[i])
		}
		res = append(res, str)
	}
	return res
}
