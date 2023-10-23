package main

import "fmt"

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
}

// leetcode997_找到小镇的法官
func findJudge(N int, trust [][]int) int {
	out := make([]int, N+1)
	in := make([]int, N+1)
	for i := range trust {
		out[trust[i][0]] = -1
		in[trust[i][1]]++
	}
	for i := 1; i <= N; i++ {
		// 出度为0，入度为N-1
		if out[i] == 0 && in[i] == N-1 {
			return i
		}
	}
	return -1
}
