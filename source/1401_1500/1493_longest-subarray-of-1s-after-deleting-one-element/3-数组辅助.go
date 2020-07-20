package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 0, 0, 1}))
}

func longestSubarray(nums []int) int {
	arr := make([]int, 0)
	count := 0
	for _, v := range nums {
		if v == 0 {
			arr = append(arr, count)
			count = 0
			continue
		}
		count++
	}
	arr = append(arr, count)
	if len(arr) == 1 {
		return arr[0] - 1
	}
	res := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+arr[i+1] > res {
			res = arr[i] + arr[i+1]
		}
	}
	return res
}
