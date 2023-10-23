package main

func main() {

}

// leetcode785_判断二分图
func isBipartite(graph [][]int) bool {
	n := len(graph)
	m := make(map[int]int) // 分组： 0一组，1一组
	for i := 0; i < n; i++ {
		// 没有被分配过，分配到0一组
		if _, ok := m[i]; ok == true {
			continue
		}
		m[i] = 0
		queue := make([]int, 0)
		queue = append(queue, i)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for i := 0; i < len(graph[node]); i++ {
				target := graph[node][i]
				if _, ok := m[target]; ok == false {
					m[target] = 1 - m[node] // 相反一组
					queue = append(queue, target)
				} else if m[node] == m[target] { // 已经分配，查看是否同一组
					return false
				}
			}
		}
	}
	return true
}
