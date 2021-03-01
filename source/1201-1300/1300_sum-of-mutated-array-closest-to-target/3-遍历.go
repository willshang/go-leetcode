package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findBestValue([]int{4, 9, 3}, 10))
}

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	res := target / n
	diff := target + 1
	for {
		a := getSum(arr, res)
		if diff <= abs(target-a) {
			return res - 1
		}
		diff = abs(target - a)
		res = res + 1
	}
	return res
}

func getSum(nums []int, target int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= target {
			res = res + nums[i]
		} else {
			res = res + target
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
