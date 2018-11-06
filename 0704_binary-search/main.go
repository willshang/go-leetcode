package main

import "fmt"

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 9
	fmt.Println(search(nums,target))
}
func search(nums []int, target int) int {
	l, r := 0,len(nums)-1

	for l <= r {
		m := (l+r) /2
		switch  {
		case nums[m] < target:
			l = m + 1
		case nums[m] > target:
			r = m - 1
		default:
			return m
		}
	}

	return -1
}