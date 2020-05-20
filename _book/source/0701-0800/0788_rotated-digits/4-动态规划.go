package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotatedDigits(10))
}

// 每个数字由(i/10)和(i%10)组成
// dp[i]={dp[i/10],dp[i%10]}
func rotatedDigits(N int) int {
	dp := []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	if N >= 10 {
		dp = append(dp, make([]int, N-9)...)
	}
	res := 0
	for i := 0; i <= N; i++ {
		if dp[i/10] == -1 || dp[i%10] == -1 {
			dp[i] = -1
		} else if dp[i] = dp[i/10] | dp[i%10]; dp[i] == 1 {
			// arr[i/10] = 1/0 arr[i%10] == 1/0
			// 异或操作，确保把0，1，8组成的数字剔除
			// 0|0 == 0
			// 0|1 == 1
			// 1|0 == 1
			// 1|1 == 1
			res++
		}
	}
	return res
}
