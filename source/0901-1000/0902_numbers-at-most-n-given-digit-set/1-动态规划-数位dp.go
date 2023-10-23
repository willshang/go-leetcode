package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
	fmt.Println(atMostNGivenDigitSet([]string{"3", "5"}, 4))
}

// leetcode902_最大为N的数字组合
func atMostNGivenDigitSet(digits []string, n int) int {
	str := fmt.Sprintf("%d", n)
	m := len(digits)
	k := len(str)
	dp := make([]int, k+1) // dp[i] 表示 长度为i的N（最后i位）的个数
	dp[0] = 1
	// 1、考虑k位数的可能
	for i := 1; i <= k; i++ {
		value := str[k-i] - '0'
		for j := 0; j < len(digits); j++ {
			v := digits[j][0] - '0'
			if v < value { // 小于，剩下i-1位随便用：+  m^(i-1)
				dp[i] = dp[i] + int(math.Pow(float64(m), float64(i-1)))
			} else if v == value { // 等于：考虑i-1位
				dp[i] = dp[i] + dp[i-1]
			}
		}
	}
	// 2、考虑非k位数的可能
	for i := 1; i < k; i++ {
		dp[k] = dp[k] + int(math.Pow(float64(m), float64(i))) // + m^i
	}
	return dp[k]
}
