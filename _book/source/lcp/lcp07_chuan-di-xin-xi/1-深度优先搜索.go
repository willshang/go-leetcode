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

var ways [][]bool

func numWays(n int, relation [][]int, k int) int {
	ways = make([][]bool, n)
	for i := range ways {
		ways[i] = make([]bool, n)
	}
	sum := 0
	for i := 0; i < len(relation); i++ {
		ways[relation[i][0]][relation[i][1]] = true
	}
	for i := 0; i < n; i++ {
		if ways[0][i] == true {
			sum = sum + dfs(i, n, k-1)
		}
	}
	return sum
}

func dfs(i, n, k int) int {
	sum := 0
	if k < 0 || i < 0 || i >= n {
		return 0
	}
	if k == 0 && i == n-1 {
		return 1
	} else {
		for j := 0; j < n; j++ {
			if ways[i][j] == true {
				sum += dfs(j, n, k-1)
			}
		}
	}
	return sum
}
