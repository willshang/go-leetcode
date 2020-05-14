package main

import (
	"fmt"
)

func main() {
	//fmt.Println(checkPossibility([]int{4,2,3}))
	//fmt.Println(checkPossibility([]int{4,2,1}))
	fmt.Println(checkPossibility([]int{1, 2, 3, 5, 4, 6}))
}

func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		res := make([]int, 0)
		res = append(res, nums[0:i]...)
		res = append(res, nums[i+1:]...)
		if isSort(res) {
			return true
		}
	}
	return false
}

func isSort(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
