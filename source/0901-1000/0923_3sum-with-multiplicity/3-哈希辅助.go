package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}

// leetcode923_三数之和的多种可能
func threeSumMulti(arr []int, target int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		res = (res + m[target-arr[i]]) % 1000000007 // 以当前i为k，求有多少i+j
		for j := 0; j < i; j++ {
			m[arr[i]+arr[j]]++
		}
	}
	return res % 1000000007
}
