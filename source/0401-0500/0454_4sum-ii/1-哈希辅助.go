package main

import "fmt"

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}

// leetcode454_四数相加II
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	m := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			res = res + m[0-c-d]
		}
	}
	return res
}
