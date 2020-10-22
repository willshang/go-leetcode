package main

import "fmt"

func main() {
	//fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6))
	fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9))
}

// leetcode1590_使数组和能被P整除
func minSubarray(nums []int, p int) int {
	n := len(nums)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = (arr[i] + nums[i]) % p
	}
	if arr[n] == 0 {
		return 0
	}
	targetValue := arr[n]
	res := n
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < n; i++ {
		target := (arr[i+1] - targetValue + p) % p
		if value, ok := m[target]; ok {
			if res > i-value {
				res = i - value
			}
		}
		m[arr[i+1]] = i
	}
	if res == n {
		return -1
	}
	return res
}
