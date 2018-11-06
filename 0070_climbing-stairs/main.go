package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}
func climbStairs(n int) int {
	if n < 2{
		return 1
	}
	result := make([]int, n+1)
	result[0],result[1] = 1,1
	for i := 2; i <= n; i++{
		result[i] = result[i-1] + result[i-2]
	}

	return result[n]
}