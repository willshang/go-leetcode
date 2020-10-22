package main

import "fmt"

func main() {
	fmt.Println(tribonacci(5))
}

// leetcode1137_第N个泰波那契数
func tribonacci(n int) int {
	arr := make([]int, n+3)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
	}
	return arr[n]
}
