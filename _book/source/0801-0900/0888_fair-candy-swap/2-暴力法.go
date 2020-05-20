package main

import "fmt"

func main() {
	A := []int{1, 2, 5}
	B := []int{2, 4}

	fmt.Println(fairCandySwap(A, B))
}

func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	for _, v := range A {
		sumA = sumA + v
	}
	for _, v := range B {
		sumB = sumB + v
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if sumA-A[i]+B[j] == sumB-B[j]+A[i] {
				return []int{A[i], B[j]}
			}
		}
	}
	return nil
}
