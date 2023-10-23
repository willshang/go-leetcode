package main

func main() {

}

func minReorder(n int, connections [][]int) int {
	m := make(map[int]map[int]int)
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
	res := 0
	visited := make(map[int]bool)
	visited[0] = true
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for k, v := range m[node] {
			if visited[k] == true {
				continue
			}
			visited[k] = true
			if v == 1 {
				res++
			}
			queue = append(queue, k)
		}
	}
	return res
}
