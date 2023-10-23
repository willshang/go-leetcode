package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(maxProduct("leetcodecom"))
}

func maxProduct(s string) int {
	res := 0
	n := len(s)
	total := 1 << n
	arr := make([]int, 0)
	for i := 1; i < total; i++ {
		if judge(s, i) {
			arr = append(arr, i)
		}
	}
	for i := 0; i < len(arr); i++ { // 枚举回文状态
		for j := i + 1; j < len(arr); j++ {
			if arr[i]&arr[j] == 0 {
				a, b := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
				res = max(res, a*b)
			}
		}
	}
	return res
}

func judge(s string, status int) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && (status&(1<<left)) == 0 {
			left++
		}
		for left < right && (status&(1<<right)) == 0 {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
