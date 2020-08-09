package main

import "fmt"

func main() {
	fmt.Println(findMagicIndex([]int{0, 2, 3, 4, 5}))
}

func findMagicIndex(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	res := search(nums, left, mid-1)
	if res != -1 {
		return res
	} else if nums[mid] == mid {
		return mid
	}
	return search(nums, mid+1, right)
}
