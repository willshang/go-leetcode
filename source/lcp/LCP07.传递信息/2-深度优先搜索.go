package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 2},
		{2, 1},
		{3, 4},
		{2, 3},
		{1, 4},
		{2, 0},
		{0, 4},
	}
	fmt.Println(numWays(5, arr, 3))
}

var res, K, N int

func numWays(n int, relation [][]int, k int) int {
	res = 0
	K = k
	N = n
	dfs(0, relation, 0)
	return res
}

func dfs(n int, relation [][]int, k int) {
	if n == N-1 && k == K {
		res++
	}
	if k > K {
		return
	}
	for i := 0; i < len(relation); i++ {
		if relation[i][0] == n {
			dfs(relation[i][1], relation, k+1)
		}
	}
}
