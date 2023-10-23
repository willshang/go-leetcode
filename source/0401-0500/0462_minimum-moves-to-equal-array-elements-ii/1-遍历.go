package main

import "sort"

func main() {

}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	target := nums[len(nums)/2]
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res + abs(target-nums[i])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
