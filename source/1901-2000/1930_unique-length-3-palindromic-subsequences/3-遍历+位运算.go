package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countPalindromicSubsequence("aabca"))
}

func countPalindromicSubsequence(s string) int {
	n := len(s)
	res := 0
	pre, suf := make([]int, n+1), make([]int, n+1)
	arr := make([]int, 26)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] | (1 << int(s[i]-'a'))
	}
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] | (1 << int(s[i]-'a'))
	}
	for i := 1; i < n-1; i++ {
		index := int(s[i] - 'a')
		arr[index] = arr[index] | (pre[i] & suf[i+1])
	}

	for i := 0; i < 26; i++ {
		res = res + bits.OnesCount(uint(arr[i]))
	}
	return res
}
