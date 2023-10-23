package main

import "sort"

func main() {

}

func pairSums(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if target == sum {
			res = append(res, []int{nums[left], nums[right]})
			left++
			right--
		} else if target > sum {
			left++
		} else {
			right--
		}
	}
	return res
}
