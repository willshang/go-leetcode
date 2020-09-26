package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoSum(2))
}

var arr []int
var start int

func twoSum(n int) []float64 {
	res := make([]float64, 0)
	if n < 1 {
		return res
	}
	start = n
	arr = make([]int, 5*n+1)
	for i := 1; i <= 6; i++ {
		dfs(n, i)
	}
	total := math.Pow(float64(6), float64(n))
	for i := n; i <= 6*n; i++ {
		ratio := float64(arr[i-n]) / total
		res = append(res, ratio)
	}
	return res
}

func dfs(current, sum int) {
	if current == 1 {
		arr[sum-start]++
	} else {
		for i := 1; i <= 6; i++ {
			dfs(current-1, sum+i)
		}
	}
}
