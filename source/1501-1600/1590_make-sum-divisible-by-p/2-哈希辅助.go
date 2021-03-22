package main

import "fmt"

func main() {
	//fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6))
	fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9))
}

func minSubarray(nums []int, p int) int {
	n := len(nums)
	targetValue := 0
	for i := 0; i < n; i++ {
		targetValue = (targetValue + nums[i]) % p
	}
	if targetValue == 0 {
		return 0
	}
	res := n
	m := make(map[int]int)
	m[0] = -1
	total := 0
	for i := 0; i < n; i++ {
		total = (total + nums[i] + p) % p
		target := (total - targetValue + p) % p
		if value, ok := m[target]; ok {
			if res > i-value {
				res = i - value
			}
		}
		m[total] = i
	}
	if res == n {
		return -1
	}
	return res
}
