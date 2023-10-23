package main

func main() {

}

// leetcode797_所有可能的路径
var res [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = make([][]int, 0)
	dfs(graph, 0, len(graph)-1, make([]int, 0))
	return res
}

func dfs(graph [][]int, cur, target int, path []int) {
	if cur == target {
		path = append(path, cur)
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(graph[cur]); i++ {
		dfs(graph, graph[cur][i], target, append(path, cur))
	}
}
