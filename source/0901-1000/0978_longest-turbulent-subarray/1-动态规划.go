package main

import "fmt"

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
}

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1
	for i := 1; i < n; i++ {
		down[i] = 1
		up[i] = 1
		if arr[i] > arr[i-1] {
			up[i] = down[i-1] + 1
		} else if arr[i] < arr[i-1] {
			down[i] = up[i-1] + 1
		}
	}
	res := 1
	for i := 0; i < n; i++ {
		res = max(res, up[i])
		res = max(res, down[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
