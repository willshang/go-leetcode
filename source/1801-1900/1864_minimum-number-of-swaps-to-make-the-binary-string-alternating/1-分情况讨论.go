package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(minSwaps("01011"))
}

// leetcode1864_构成交替字符串需要的最小交换次数
func minSwaps(s string) int {
	a, b := strings.Count(s, "1"), strings.Count(s, "0")
	if a-b > 1 || b-a > 1 {
		return -1
	}
	n := len(s)
	res := math.MaxInt32
	if a == (n+1)/2 && b == n/2 { // 以1开头：1010xxx
		count := 0
		for i := 0; i < n; i++ {
			if int(s[i]-'0') == i%2 { // 0在偶数下标位置，1在奇数下标位置
				count++
			}
		}
		res = min(res, count/2)
	}
	if b == (n+1)/2 && a == n/2 { // 以0开头：0101xxx
		count := 0
		for i := 0; i < n; i++ {
			if int(s[i]-'0') != i%2 { // 1在偶数下标位置，0在奇数下标位置
				count++
			}
		}
		res = min(res, count/2)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
