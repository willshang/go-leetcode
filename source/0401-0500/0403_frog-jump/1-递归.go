package main

func main() {

}

var m [][]int

func canCross(stones []int) bool {
	n := len(stones)
	m = make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = -1
		}
	}
	return dfs(stones, 0, 0) == 1
}

func dfs(stones []int, index int, k int) int {
	if m[index][k] >= 0 {
		return m[index][k]
	}
	for i := index + 1; i < len(stones); i++ {
		next := stones[i] - stones[index]
		if k-1 <= next && next <= k+1 { // k-1、k或 k+1
			if dfs(stones, i, next) == 1 {
				m[index][k] = 1
				return 1
			}
		}
	}
	if index == len(stones)-1 {
		m[index][k] = 1
	} else {
		m[index][k] = 0
	}
	return m[index][k]
}
