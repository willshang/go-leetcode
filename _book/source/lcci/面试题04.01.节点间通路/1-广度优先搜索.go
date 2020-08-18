package main

func main() {

}

// 程序员面试金典04.01_节点间通路
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	edges := make([][]int, n)
	// 邻接表
	for i := 0; i < len(graph); i++ {
		a := graph[i][0]
		b := graph[i][1]
		edges[a] = append(edges[a], b)
	}
	queue := make([]int, 0)
	queue = append(queue, start)
	visited := make([]bool, n)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node] = true
		if node == target {
			return true
		}
		for i := 0; i < len(edges[node]); i++ {
			if visited[edges[node][i]] == false {
				if edges[node][i] == target {
					return true
				}
				queue = append(queue, edges[node][i])
			}
		}
	}
	return false
}
