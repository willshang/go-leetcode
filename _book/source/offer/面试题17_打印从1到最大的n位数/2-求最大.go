package main

import "fmt"

func main() {
	fmt.Println(printNumbers(3))
}

func printNumbers(n int) []int {
	res := make([]int, 0)
	maxValue := 1
	for i := 1; i <= n; i++ {
		maxValue = maxValue * 10
	}
	maxValue = maxValue - 1
	for i := 1; i <= maxValue; i++ {
		res = append(res, i)
	}
	return res
}
