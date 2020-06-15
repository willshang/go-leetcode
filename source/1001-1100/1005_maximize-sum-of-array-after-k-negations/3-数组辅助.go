package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	// fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	//fmt.Println(largestSumAfterKNegations([]int{-2, 0, -2}, 7))
}

// leetcode1005_K次取反后最大化的数组和
func largestSumAfterKNegations(A []int, K int) int {
	arr := make([]int, 201)
	for i := 0; i < len(A); i++ {
		arr[A[i]+100]++
	}
	i := 0
	for K > 0 {
		for arr[i] == 0 {
			i++
		}
		if i > 100 {
			break
		}
		arr[i]--
		arr[200-i]++
		K--
	}
	if K%2 == 1 && i != 100 {
		for j := i; j < len(arr); j++ {
			if arr[j] > 0 {
				arr[j]--
				arr[200-j]++
				break
			}
		}
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		res = res + (i-100)*arr[i]
	}
	return res
}
