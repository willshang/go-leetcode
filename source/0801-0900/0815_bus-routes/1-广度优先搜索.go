package main

func main() {

}

// leetcode815_公交路线
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	n := len(routes)
	m := make(map[int][]int) // 该公交车经过第几条线路的数组
	for i := 0; i < n; i++ {
		for j := 0; j < len(routes[i]); j++ {
			node := routes[i][j]
			m[node] = append(m[node], i)
		}
	}
	count := 1
	queue := make([]int, 0)
	queue = append(queue, source)
	visited := make(map[int]bool) // 第几条线路访问情况
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			for j := 0; j < len(m[node]); j++ {
				if visited[m[node][j]] == false {
					visited[m[node][j]] = true
					for k := 0; k < len(routes[m[node][j]]); k++ {
						if routes[m[node][j]][k] == target {
							return count
						}
						queue = append(queue, routes[m[node][j]][k])
					}
				}
			}
		}
		count++
		queue = queue[length:]
	}
	return -1
}
