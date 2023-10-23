package main

import (
	"sort"
)

func main() {

}

// leetcode2192_有向无环图中一个节点的所有祖先
func getAncestors(n int, edges [][]int) [][]int {
	m := make([]map[int]bool, n) // 祖先节点要去重
	for i := 0; i < n; i++ {
		m[i] = make(map[int]bool)
	}
	arr := make([][]int, n)
	inDegree := make([]int, n) // 入度
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1] // a=>b
		arr[a] = append(arr[a], b)
		inDegree[b]++
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 { // 入度为0的
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < len(arr[cur]); i++ {
			next := arr[cur][i]
			m[next][cur] = true     // 保存父节点
			for k := range m[cur] { // 把父节点的祖先节点也保存下来
				m[next][k] = true
			}
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	// 排序
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		for k, _ := range m[i] {
			res[i] = append(res[i], k)
		}
		sort.Ints(res[i])
	}
	return res
}
