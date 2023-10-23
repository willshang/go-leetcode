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
	m := make(map[int]int, 0)
	for i := 1; i < total; i++ {
		if judge(s, i) {
			m[i] = bits.OnesCount(uint(i))
		}
	}
	for i := 1; i < total; i++ { // 遍历状态
		for j := i; j > 0; j = (j - 1) & i { // 枚举子集
			res = max(res, m[j]*m[j^i]) // 子集 * 子集的补集
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
