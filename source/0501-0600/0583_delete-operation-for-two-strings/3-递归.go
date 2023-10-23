package main

func main() {

}

var m [][]int

func minDistance(word1 string, word2 string) int {
	a, b := len(word1), len(word2)
	m = make([][]int, a+1)
	for i := 0; i <= a; i++ {
		m[i] = make([]int, b+1)
		for j := 0; j <= b; j++ {
			m[i][j] = -1
		}
	}
	total := dfs(word1, word2, 0, 0)
	return a + b - 2*total
}

func dfs(word1 string, word2 string, i, j int) int {
	if len(word1) == i || len(word2) == j {
		return 0
	}
	if m[i][j] > -1 {
		return m[i][j]
	}
	if word1[i] == word2[j] {
		m[i][j] = dfs(word1, word2, i+1, j+1) + 1
	} else {
		m[i][j] = max(dfs(word1, word2, i, j+1), dfs(word1, word2, i+1, j))
	}
	return m[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
