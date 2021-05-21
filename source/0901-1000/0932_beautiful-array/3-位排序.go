package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(beautifulArray(10))
	fmt.Println(beautifulArray(16))
}

func beautifulArray(N int) []int {
	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = i + 1
	}
	sort.Slice(res, func(i, j int) bool {
		k := 0
		for res[i]&k == res[j]&k {
			k++
		}
		return res[i]&k < res[j]&k
	})
	return res
}
