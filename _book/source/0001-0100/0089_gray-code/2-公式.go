package main

import "fmt"

func main() {
	fmt.Println(grayCode(5))
}

// leetcode89_格雷编码
func grayCode(n int) []int {
	total := 1 << n
	res := make([]int, 0)
	for i := 0; i < total; i++ {
		res = append(res, i^(i>>1))
	}
	return res
}
