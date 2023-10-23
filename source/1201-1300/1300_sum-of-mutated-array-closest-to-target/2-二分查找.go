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
	temp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		temp[i] = temp[i-1] + arr[i-1]
	}
	left, right := 0, arr[n-1]
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		index := sort.SearchInts(arr, mid)
		if index < 0 {
			index = abs(index) - 1
		}
		total := temp[index] + (n-index)*mid
		if total <= target {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	a := getSum(arr, res)
	b := getSum(arr, res+1)
	if abs(a-target) > abs(b-target) {
		return res + 1
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
