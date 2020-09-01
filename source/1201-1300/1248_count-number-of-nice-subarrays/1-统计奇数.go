package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
}

// leetcode1248_统计「优美子数组」
func numberOfSubarrays(nums []int, k int) int {
	res := 0
	arr := make([]int, 0)
	arr = append(arr, -1)
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			arr = append(arr, i)
		}
	}
	arr = append(arr, len(nums))
	for i := 1; i+k < len(arr); i++ {
		res = res + (arr[i]-arr[i-1])*(arr[i+k]-arr[i+k-1])
	}
	return res
}
