package main

func main() {

}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	visited := make([]int, n) // 0：未访问、1：在访问（有环）、2：安全
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if dfs(graph, i, visited) == true {
			res = append(res, i)
		}
	}
	return res
}

func dfs(graph [][]int, index int, visited []int) bool {
	if visited[index] > 0 {
		return visited[index] == 2
	}
	visited[index] = 1
	for i := 0; i < len(graph[index]); i++ {
		next := graph[index][i]
		if visited[next] == 1 || dfs(graph, next, visited) == false {
			return false
		}
	}
	visited[index] = 2
	return true
}
