package main

import "fmt"

func main() {
	fmt.Println(countBits(10))
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
