package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
}

func numberOfSubarrays(nums []int, k int) int {
	res := 0
	dp := make([]int, 0)
	count := 0
	for i := 0; i < len(nums); i++ {
		count++
		if nums[i]%2 == 1 {
			dp = append(dp, count)
			count = 0
		}
		if len(dp) >= k {
			res = res + dp[len(dp)-k]
		}
	}
	return res
}
