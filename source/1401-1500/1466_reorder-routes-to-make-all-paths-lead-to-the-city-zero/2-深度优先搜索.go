package main

func main() {

}

// leetcode1466_重新规划路线
var res int
var m map[int]map[int]int

func minReorder(n int, connections [][]int) int {
	m = make(map[int]map[int]int)
	for i := 0; i < len(connections); i++ {
		a, b := connections[i][0], connections[i][1]
		if _, ok := m[a]; ok == false {
			m[a] = make(map[int]int)
		}
		if _, ok := m[b]; ok == false {
			m[b] = make(map[int]int)
		}
		m[a][b] = 1  // a->b
		m[b][a] = -1 // b->a
	}
	res = 0
	visited := make(map[int]bool)
	visited[0] = true
	dfs(0, visited)
	return res
}

func dfs(start int, visited map[int]bool) {
	for k, v := range m[start] {
		if visited[k] == true {
			continue
		}
		visited[k] = true
		if v == 1 {
			res++
		}
		dfs(k, visited)
	}
}
