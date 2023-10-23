package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 3, 1, 2, 3}
	wiggleSort(arr)
	fmt.Println(arr)
}

// 程序员面试金典10.11_峰与谷
func wiggleSort(nums []int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i = i + 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}
