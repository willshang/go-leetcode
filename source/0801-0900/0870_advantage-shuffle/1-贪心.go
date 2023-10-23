package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
}

// leetcode870_优势洗牌
func advantageCount(A []int, B []int) []int {
	res := make([]int, len(A))
	sort.Ints(A)
	arr := make([][2]int, 0)
	for i := 0; i < len(B); i++ {
		arr = append(arr, [2]int{i, B[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	left, right := 0, len(A)-1
	for i := 0; i < len(A); i++ {
		if A[i] > arr[left][1] { // 满足条件放前面
			index := arr[left][0]
			left++
			res[index] = A[i]
		} else { // 不满足条件放后面
			index := arr[right][0]
			right--
			res[index] = A[i]
		}
	}
	return res
}
