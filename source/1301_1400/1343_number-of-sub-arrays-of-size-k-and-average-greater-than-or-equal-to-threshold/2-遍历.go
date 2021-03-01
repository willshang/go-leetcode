package main

import "fmt"

func main() {
	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))
}

// leetcode1343_大小为K且平均值大于等于阈值的子数组数目
func numOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	sum := 0
	for i := 0; i < k; i++ {
		sum = sum + arr[i]
	}
	res := 0
	target := k * threshold
	if sum >= target {
		res++
	}
	for i := k; i < n; i++ {
		sum = sum + arr[i] - arr[i-k]
		if sum >= target {
			res++
		}
	}
	return res
}
