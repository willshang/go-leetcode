package main

import "fmt"

func main() {
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
}

func maxUncrossedLines(A []int, B []int) int {
	n, m := len(A), len(B)
	prev := make([]int, m+1)
	cur := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if A[i-1] == B[j-1] {
				cur[j] = prev[j-1] + 1
			} else {
				cur[j] = max(prev[j], cur[j-1])
			}
		}
		copy(prev, cur)
	}
	return cur[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
