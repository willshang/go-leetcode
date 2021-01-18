package main

import "fmt"

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 0))
}

func numsSameConsecDiff(n int, k int) []int {
	dp := make([][]int, n)
	dp[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; i < n; i++ {
		temp := make([]int, 0)
		for j := 0; j < len(dp[i-1]); j++ {
			num := dp[i-1][j]
			a := num%10 - k
			if a >= 0 {
				temp = append(temp, num*10+a)
			}
			b := num%10 + k
			if b <= 9 && k > 0 {
				temp = append(temp, num*10+b)
			}
		}
		dp[i] = temp
	}
	return dp[n-1]
}
