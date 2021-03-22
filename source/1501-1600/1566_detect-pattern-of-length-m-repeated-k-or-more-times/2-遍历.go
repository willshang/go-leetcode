package main

import "fmt"

func main() {
	fmt.Println(containsPattern([]int{1, 2, 4, 4, 4, 4}, 1, 3))
	//fmt.Println(containsPattern([]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2))
	//fmt.Println(containsPattern([]int{1, 2, 1, 2, 1, 3}, 2, 3))
	//fmt.Println(containsPattern([]int{2, 2}, 1, 2))
}

// leetcode1566_重复至少K次且长度为M的模式
func containsPattern(arr []int, m int, k int) bool {
	n := len(arr)
	if n < m*k {
		return false
	}
	i, j := 0, 0
	for i = 0; i <= n-m*k; i++ {
		for j = i + m; j < i+m*k; j++ {
			if arr[j] != arr[j-m] {
				break
			}
		}
		if j == i+m*k {
			return true
		}
	}
	return false
}
