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

func wiggleSort(nums []int) {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)
	j := len(nums)
	k := (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		if i%2 == 1 {
			j--
			nums[i] = arr[j]
		} else {
			k--
			nums[i] = arr[k]
		}
	}
}
