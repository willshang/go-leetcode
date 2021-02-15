package main

import "fmt"

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
}

// leetcode978_最长湍流子数组
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	up := 1
	down := 1
	res := 1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			up, down = down+1, 1
		} else if arr[i] < arr[i-1] {
			up, down = 1, up+1
		} else {
			up, down = 1, 1
		}
		res = max(res, max(up, down))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
