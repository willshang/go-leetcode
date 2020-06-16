package main

import "fmt"

func main() {
	A := []int{1, 2, 5}
	B := []int{2, 4}

	fmt.Println(fairCandySwap(A, B))
}

// leetcode888_公平的糖果交换
func fairCandySwap(A []int, B []int) []int {
	m := make(map[int]bool)
	sumA := 0
	sumB := 0
	for _, v := range A {
		sumA = sumA + v
		m[v] = true
	}
	for _, v := range B {
		sumB = sumB + v
	}
	half := (sumA - sumB) / 2
	a, b := 0, 0
	// sumA-A[i]+B[j] == sumB-B[j]+A[i]
	// sumA-sumB=2(A[i]-B[j])
	// (sumA-sumB)/2 = A[i]-B[j]
	for _, b = range B {
		a = b + half
		if m[a] == true {
			return []int{a, b}
		}
	}
	return nil
}
