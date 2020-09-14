package main

import "fmt"

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	mA := make(map[int]int)
	mB := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			mA[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			mB[c+d]++
		}
	}
	for k, v := range mA {
		res = res + v*mB[-k]
	}
	return res
}
