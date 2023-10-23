package main

func main() {

}

func validPath(n int, edges [][]int, source int, destination int) bool {
	m := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1] // a<=>b
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	queue := make([]int, 0)
	queue = append(queue, source)
	visited := make(map[int]bool)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == destination {
			return true
		}
		for i := 0; i < len(m[cur]); i++ {
			next := m[cur][i]
			if visited[next] == false {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}
	return false
}
