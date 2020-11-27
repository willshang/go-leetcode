package main

import "sort"

func main() {

}

// leetcode1630_等差子数组
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, 0)
	for i := 0; i < len(l); i++ {
		flag := true
		arr := make([]int, 0)
		arr = append(arr, nums[l[i]:r[i]+1]...)
		sort.Ints(arr)
		for j := 2; j < len(arr); j++ {
			if arr[j]-arr[j-1] != arr[1]-arr[0] {
				flag = false
				break
			}
		}
		res = append(res, flag)
	}
	return res
}
