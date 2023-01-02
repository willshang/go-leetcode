package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, l int
	fmt.Scanf("%d %d\n", &n, &l)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		arr[i] = temp
	}
	res := getResult(arr, l)
	fmt.Printf("%.2f\n", res)
}

func getResult(arr []int, l int) float64 {
	sort.Ints(arr)
	maxValue := 0
	for i := 1; i < len(arr); i++ {
		maxValue = max(maxValue, arr[i]-arr[i-1])
	}
	left := arr[0]
	right := l - arr[len(arr)-1]
	res := max(maxValue, left*2)
	res = max(res, right*2)
	return float64(res) / float64(2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
