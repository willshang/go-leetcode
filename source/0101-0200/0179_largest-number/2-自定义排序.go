package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return fmt.Sprintf("%d%d", nums[i], nums[j]) >=
			fmt.Sprintf("%d%d", nums[j], nums[i])
	})
	res := ""
	for i := 0; i < len(nums); i++ {
		res = res + strconv.Itoa(nums[i])
	}
	if res[0] == '0' {
		return "0"
	}
	return res
}
