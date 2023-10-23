package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

// leetcode922_按奇偶排序数组 II
func sortArrayByParityII(A []int) []int {
	i := 0
	j := 1
	for i < len(A) || j < len(A) {
		for i < len(A) && A[i]%2 == 0 {
			i = i + 2
		}
		for j < len(A) && A[j]%2 == 1 {
			j = j + 2
		}
		if i >= len(A) || j >= len(A) {
			break
		}
		A[i], A[j] = A[j], A[i]
	}
	return A
}
