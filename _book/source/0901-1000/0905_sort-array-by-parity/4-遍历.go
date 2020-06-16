package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

// leetcode905_按奇偶排序数组
func sortArrayByParity(A []int) []int {
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[count], A[i] = A[i], A[count]
			count++
		}
	}
	return A
}
