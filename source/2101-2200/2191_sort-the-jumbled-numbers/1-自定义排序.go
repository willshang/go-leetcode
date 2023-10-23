package main

import "sort"

func main() {

}

// leetcode2191_将杂乱无章的数字排序
func sortJumbled(mapping []int, nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool {
		return transfer(mapping, nums[i]) < transfer(mapping, nums[j])
	})
	return nums
}

func transfer(mapping []int, num int) int {
	res := 0
	if num == 0 {
		return mapping[0]
	}
	v := 1
	for ; num > 0; num = num / 10 {
		res = res + mapping[num%10]*v
		v = v * 10
	}
	return res
}
