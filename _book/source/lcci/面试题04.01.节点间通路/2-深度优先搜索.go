package main

func main() {

}

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	edges := make([][]int, n)
	// 邻接表
	for i := 0; i < len(graph); i++ {
		a := graph[i][0]
		b := graph[i][1]
		edges[a] = append(edges[a], b)
	}

	visited := make([]bool, n)
	return dfs(edges, visited, start, target)
}

func dfs(edges [][]int, visited []bool, start, target int) bool {
	if start == target {
		return true
	}
	visited[start] = true
	for i := 0; i < len(edges[start]); i++ {
		if visited[edges[start][i]] == false {
			if edges[start][i] == target {
				return true
			}
			if dfs(edges, visited, edges[start][i], target) {
				return true
			}
		}
	}
	return false
}
