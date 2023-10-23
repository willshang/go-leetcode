package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minNumber([]int{10, 2}))
}

// 剑指offer_面试题45_把数组排成最小的数
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		if fmt.Sprintf("%d%d", nums[i], nums[j]) <
			fmt.Sprintf("%d%d", nums[j], nums[i]) {
			return true
		}
		return false
	})
	str := ""
	for i := 0; i < len(nums); i++ {
		str = str + fmt.Sprintf("%d", nums[i])
	}
	return str
}
