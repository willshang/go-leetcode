package main

import "fmt"

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}))
}

func minDominoRotations(A []int, B []int) int {
	a, b := A[0], B[0]
	resA := check(A, B, a)
	resB := check(A, B, b)
	if resA > 0 && resB > 0 { // 都行选最少
		return min(resA, resB)
	}
	return max(resA, resB) // 不行选最多
}

func check(A []int, B []int, target int) int {
	a, b := 0, 0
	for i := 0; i < len(A); i++ {
		if A[i] != target && B[i] != target { // 都不满足直接返回-1
			return -1
		} else if A[i] != target { // A[i]不满足，需要交换
			a++
		} else if B[i] != target { // B[i]不满足，需要交换
			b++
		}
	}
	return min(a, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
