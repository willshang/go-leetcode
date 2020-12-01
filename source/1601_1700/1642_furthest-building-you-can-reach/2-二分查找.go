package main

import "sort"

func main() {

}

// leetcode1642_可以到达的最远建筑
func furthestBuilding(heights []int, bricks int, ladders int) int {
	left, right := 0, len(heights)
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		if check(heights[0:mid], bricks, ladders) {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}
	return res - 1
}

func check(heights []int, bricks int, ladders int) bool {
	arr := make([]int, 0)
	for i := 1; i < len(heights); i++ {
		need := heights[i] - heights[i-1]
		if need > 0 {
			arr = append(arr, need)
		}
	}
	sort.Ints(arr)
	i := 0
	for ; i < len(arr); i++ {
		if bricks-arr[i] >= 0 {
			bricks = bricks - arr[i]
			continue
		}
		if ladders > 0 {
			ladders--
			continue
		}
		break
	}
	return i == len(arr)
}
