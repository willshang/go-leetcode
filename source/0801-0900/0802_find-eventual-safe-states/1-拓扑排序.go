package main

func main() {

}

// leetcode802_找到最终的安全状态
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	safeArr := make([]bool, n) // 安全节点
	arr := make([][]int, n)
	m := make(map[int]map[int]bool)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		m[i] = make(map[int]bool)
		if len(graph[i]) == 0 { // 没有出节点，是安全节点
			queue = append(queue, i)
		}
		for j := 0; j < len(graph[i]); j++ {
			a, b := i, graph[i][j]     // a=>b
			arr[b] = append(arr[b], a) // 反向b=>a
			m[a][b] = true             // a=>b
		}
	}
	res := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		safeArr[node] = true
		for i := 0; i < len(arr[node]); i++ { // 反向边遍历
			index := arr[node][i]
			delete(m[index], node)  // 删除
			if len(m[index]) == 0 { // 没有出节点
				queue = append(queue, index)
			}
		}
	}
	for i := 0; i < n; i++ {
		if safeArr[i] == true {
			res = append(res, i)
		}
	}
	return res
}
