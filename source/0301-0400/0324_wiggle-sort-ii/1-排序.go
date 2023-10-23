package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{1, 2, 1, 2, 1, 2, 1}
	arr := []int{4, 5, 5, 6}
	wiggleSort(arr)
	fmt.Println(arr)
}

// leetcode324_摆动排序II
func wiggleSort(nums []int) {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	a := 1
	for i := 0; i < len(arr)/2; i++ {
		nums[a] = arr[i]
		a = a + 2
	}
	a = 0
	for i := len(arr) / 2; i < len(arr); i++ {
		nums[a] = arr[i]
		a = a + 2
	}
}
