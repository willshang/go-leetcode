package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

var arr []int

func climbStairs(n int) int {
	arr = make([]int, n+1)
	return climbStart(0, n)
}

func climbStart(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if arr[i] > 0 {
		return arr[i]
	}
	arr[i] = climbStart(i+1, n) + climbStart(i+2, n)
	return arr[i]
}
