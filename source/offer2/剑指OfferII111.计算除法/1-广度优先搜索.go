package main

func main() {

}

type Node struct {
	to    int
	value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]int) // 计算对应的id
	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]
		if _, ok := m[a]; ok == false {
			m[a] = len(m)
		}
		if _, ok := m[b]; ok == false {
			m[b] = len(m)
		}
	}
	arr := make([][]Node, len(m)) // 邻接表
	for i := 0; i < len(equations); i++ {
		a, b := m[equations[i][0]], m[equations[i][1]]
		arr[a] = append(arr[a], Node{to: b, value: values[i]})
		arr[b] = append(arr[b], Node{to: a, value: 1 / values[i]}) // 除以
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		a, okA := m[queries[i][0]]
		b, okB := m[queries[i][1]]
		if okA == false || okB == false {
			res[i] = -1
		} else {
			res[i] = bfs(arr, a, b) // 广度优先查找
		}
	}
	return res
}

func bfs(arr [][]Node, start, end int) float64 {
	temp := make([]float64, len(arr)) // 结果的比例
	temp[start] = 1
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == end {
			return temp[node]
		}
		for i := 0; i < len(arr[node]); i++ {
			next := arr[node][i].to
			if temp[next] == 0 {
				temp[next] = temp[node] * arr[node][i].value
				queue = append(queue, next)
			}
		}
	}
	return -1
}
