package main

import "fmt"

func main() {
	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))
}

func numOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	temp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		temp[i] = temp[i-1] + arr[i-1]
	}
	res := 0
	target := k * threshold
	for i := k; i <= n; i++ {
		if temp[i]-temp[i-k] >= target {
			res++
		}
	}
	return res
}
