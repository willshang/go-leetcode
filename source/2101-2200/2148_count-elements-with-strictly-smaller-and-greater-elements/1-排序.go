package main

import "sort"

func main() {

}

func countElements(nums []int) int {
	sort.Ints(nums)
	minValue, maxValue := nums[0], nums[len(nums)-1]
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != minValue && nums[i] != maxValue {
			res++
		}
	}
	return res
}
