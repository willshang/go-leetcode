package main

import "fmt"

func main() {
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
}

func maxUncrossedLines(A []int, B []int) int {
	n, m := len(A), len(B)
	cur := make([]int, m+1)
	for i := 1; i <= n; i++ {
		pre := cur[0]
		for j := 1; j <= m; j++ {
			temp := cur[j]
			if A[i-1] == B[j-1] {
				cur[j] = pre + 1
			} else {
				cur[j] = max(cur[j], cur[j-1])
			}
			pre = temp
		}
	}
	return cur[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
