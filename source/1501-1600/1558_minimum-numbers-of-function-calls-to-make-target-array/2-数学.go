package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{2, 2}))
}

// leetcode1558_得到目标数组的最少函数调用次数
func minOperations(nums []int) int {
	res := 0
	maxValue := 0 // 保存x2的次数
	for i := 0; i < len(nums); i++ {
		count := 0
		n := nums[i]
		for n != 0 {
			if n%2 == 0 {
				n = n / 2
				count++
			} else {
				n = n - 1
				res = res + 1 // 单独加1
			}
		}
		maxValue = max(maxValue, count)
	}
	return res + maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
