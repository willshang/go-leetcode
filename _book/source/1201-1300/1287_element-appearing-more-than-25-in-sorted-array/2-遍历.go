package main

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 7, 10}))
}

// leetcode1287_有序数组中出现次数超过25%的元素
func findSpecialInteger(arr []int) int {
	step := len(arr) / 4
	for i := 0; i < len(arr)-step; i++ {
		if arr[i] == arr[i+step] {
			return arr[i]
		}
	}
	return -1
}
