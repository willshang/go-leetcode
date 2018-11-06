package main

import "fmt"

func main() {
	nums1 := []int{4,1,3,6}
	nums2 := []int{1,2,7,4}

	fmt.Println(nextGreaterElement(nums1,nums2))
}
func nextGreaterElement(findNums []int, nums []int) []int {
	index := make(map[int]int)

	for i, n := range nums{
		index[n] = i
	}

	res := make([]int,len(findNums))
	for i, n := range findNums{
		res[i] = -1
		for j := index[n] + 1; j < len(nums); j++{
			if n < nums[j]{
				res[i] = nums[j]
				break
			}
		}
	}
	return res
}