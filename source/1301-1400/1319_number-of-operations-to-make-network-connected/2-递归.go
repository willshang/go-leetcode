package main

var m map[int][]int
var visited []bool

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	m = make(map[int][]int)
	visited = make([]bool, n)
	for i := 0; i < len(connections); i++ {
		a, b := connections[i][0], connections[i][1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	res := 0
	for i := 0; i < n; i++ {
		if visited[i] == false {
			dfs(i)
			res++
		}
	}
	// 连通子图数量-1
	return res - 1
}

func dfs(i int) {
	if visited[i] == true {
		return
	}
	visited[i] = true
	for j := 0; j < len(m[i]); j++ {
		dfs(m[i][j])
	}
}
