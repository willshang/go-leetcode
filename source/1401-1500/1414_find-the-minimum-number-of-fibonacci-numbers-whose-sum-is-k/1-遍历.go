package main

import "fmt"

func main() {
	fmt.Println(findMinFibonacciNumbers(19))
}

// leetcode1414_和为K的最少斐波那契数字数目
func findMinFibonacciNumbers(k int) int {
	arr := make([]int, 2)
	arr[0] = 1
	arr[1] = 1
	for arr[len(arr)-2]+arr[len(arr)-1] <= k {
		arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
	}
	res := 0
	for i := len(arr) - 1; i >= 0; i-- {
		for arr[i] <= k {
			k = k - arr[i]
			res++
		}
	}
	return res
}
