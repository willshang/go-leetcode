package main

func main() {

}

// leetcode1971_寻找图中是否存在路径
var m map[int][]int
var visited map[int]bool

func validPath(n int, edges [][]int, source int, destination int) bool {
	m = make(map[int][]int)
	visited = make(map[int]bool)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1] // a<=>b
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	return dfs(source, destination)
}

func dfs(source int, destination int) bool {
	visited[source] = true
	if source == destination {
		return true
	}
	for i := 0; i < len(m[source]); i++ {
		next := m[source][i]
		if visited[next] == false && dfs(next, destination) {
			return true
		}
	}
	return false
}
