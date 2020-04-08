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
		res := []int{}
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

/*func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++{
		if nums[i-1] > nums[i]{
			pre := deepCopy(nums)
			pre[i-1] = pre[i]
			next := deepCopy(nums)
			next[i] = next[i-1]
			return sort.IsSorted(sort.IntSlice(pre)) || sort.IsSorted(sort.IntSlice(next))
		}
	}
	return true
}

func deepCopy(nums []int) []int  {
	res := make([]int, len(nums))
	copy(res,nums)
	return res
}*/
