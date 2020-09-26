package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
	fmt.Println(sumOddLengthSubarrays([]int{10, 11, 12}))
}

// leetcode1588_所有奇数长度子数组的和
func sumOddLengthSubarrays(arr []int) int {
	sum := make([]int, len(arr)+1)
	sum[0] = 0
	for i := 1; i <= len(arr); i++ {
		sum[i] = sum[i-1] + arr[i-1]
	}
	res := 0
	for i := 1; i <= len(arr); i = i + 2 {
		for j := 0; j < len(arr); j++ {
			k := j + i
			if k <= len(arr) {
				total := sum[k] - sum[j]
				res = res + total
			}
		}
	}
	return res
}
