package main

import "fmt"

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
}

func findJudge(N int, trust [][]int) int {
	arr := make([]int, N+1)
	for i := range trust {
		arr[trust[i][0]] = -1
		if arr[trust[i][1]] == -1 {
			continue
		}
		arr[trust[i][1]]++
	}
	for i := 1; i <= N; i++ {
		if arr[i] == N-1 {
			return i
		}
	}
	return -1
}
