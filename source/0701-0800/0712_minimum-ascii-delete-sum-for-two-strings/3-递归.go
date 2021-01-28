package main

func main() {

}

var m [][]int

func minimumDeleteSum(s1 string, s2 string) int {
	a, b := len(s1), len(s2)
	m = make([][]int, a+1)
	for i := 0; i <= a; i++ {
		m[i] = make([]int, b+1)
		for j := 0; j <= b; j++ {
			m[i][j] = -1
		}
	}

	total := dfs(s1, s2, 0, 0)
	return sumAscii(s1) + sumAscii(s2) - 2*total
}

func dfs(s1 string, s2 string, i, j int) int {
	if len(s1) == i || len(s2) == j {
		return 0
	}
	if m[i][j] > -1 {
		return m[i][j]
	}
	if s1[i] == s2[j] {
		m[i][j] = dfs(s1, s2, i+1, j+1) + int(s1[i])
	} else {
		m[i][j] = max(dfs(s1, s2, i, j+1), dfs(s1, s2, i+1, j))
	}
	return m[i][j]
}

func sumAscii(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res + int(s[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
