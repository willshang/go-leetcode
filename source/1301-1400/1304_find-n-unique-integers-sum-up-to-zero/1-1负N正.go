package main

import "fmt"

func main() {
	fmt.Println(sumZero(10))
}

func sumZero(n int) []int {
	res := make([]int, n)
	sum := 0
	for i := 0; i < n-1; i++ {
		res[i] = i + 1
		sum = sum + i + 1
	}
	res[n-1] = -sum
	return res
}
