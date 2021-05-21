package main

import "fmt"

func main() {
	fmt.Println(beautifulArray(10))
}

// leetcode932_漂亮数组
func beautifulArray(N int) []int {
	if N == 1 {
		return []int{1}
	}
	res := make([]int, 0)
	a := beautifulArray((N + 1) / 2)
	b := beautifulArray(N / 2)
	for i := 0; i < len(a); i++ {
		res = append(res, a[i]*2-1)
	}
	for i := 0; i < len(b); i++ {
		res = append(res, b[i]*2)
	}
	return res
}
