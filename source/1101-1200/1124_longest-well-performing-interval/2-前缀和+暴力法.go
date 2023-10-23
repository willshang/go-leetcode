package main

import "fmt"

func main() {
	//fmt.Println(longestWPI([]int{6, 6, 9}))
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
}

// leetcode1124_表现良好的最长时间段
func longestWPI(hours []int) int {
	arr := make([]int, 0)
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			arr = append(arr, 1)
		} else {
			arr = append(arr, -1)
		}
	}
	temp := make([]int, len(hours)+1)
	for i := 1; i <= len(hours); i++ {
		temp[i] = temp[i-1] + arr[i-1]
	}
	res := 0
	for i := 0; i < len(hours); i++ {
		for j := i; j < len(hours); j++ {
			count := temp[j+1] - temp[i]
			if count > 0 {
				res = max(res, j-i+1)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
