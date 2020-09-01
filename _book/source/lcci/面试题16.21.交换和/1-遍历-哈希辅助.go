package main

import "fmt"

func main() {
	//fmt.Println(findSwapValues([]int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3}))
	fmt.Println(findSwapValues([]int{1, 2, 3}, []int{4, 5, 6}))
}

// 程序员面试金典16.21_交换和
func findSwapValues(array1 []int, array2 []int) []int {
	m := make(map[int]bool)
	sumA, sumB := 0, 0
	for i := 0; i < len(array1); i++ {
		sumA = sumA + array1[i]
		m[array1[i]] = true
	}
	for i := 0; i < len(array2); i++ {
		sumB = sumB + array2[i]
	}
	if (sumA+sumB)%2 == 1 {
		return nil
	}
	half := (sumA - sumB) / 2
	a, b := 0, 0
	// sumA-A[i]+B[j] == sumB-B[j]+A[i]
	// sumA-sumB=2(A[i]-B[j])
	// (sumA-sumB)/2 = A[i]-B[j]
	for _, b = range array2 {
		a = b + half
		if m[a] == true {
			return []int{a, b}
		}
	}
	return nil
}
