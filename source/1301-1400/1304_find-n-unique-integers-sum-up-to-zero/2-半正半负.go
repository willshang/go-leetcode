package main

import "fmt"

func main() {
	fmt.Println(sumZero(10))
}

// leetcode1304_和为零的N个唯一整数
func sumZero(n int) []int {
	res := make([]int, 0)
	if n%2 == 1 {
		res = append(res, 0)
	}
	for i := 1; i <= n/2; i++ {
		res = append(res, i)
		res = append(res, -i)
	}
	return res
}
