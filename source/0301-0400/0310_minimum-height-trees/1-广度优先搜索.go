package main

func main() {

}

// leetcode310_最小高度树
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}
	m := make(map[int][]int)
	degree := make([]int, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
		degree[a]++
		degree[b]++
	}
	queue := make([]int, 0) // 从叶子节点开始遍历
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		length := len(queue)
		n = n - length
		for i := 0; i < length; i++ {
			node := queue[i]
			degree[node] = 0
			for j := 0; j < len(m[node]); j++ {
				temp := m[node][j]
				if degree[temp] > 0 {
					degree[temp]--
					if degree[temp] == 1 {
						queue = append(queue, temp)
					}
				}
			}
		}
		queue = queue[length:]
	}
	res := make([]int, 0)
	for i := 0; i < len(queue); i++ {
		res = append(res, queue[i])
	}
	return res
}
