package main

func main() {

}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	fa = Init(n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(graph[i]); j++ {
			target := graph[i][j]
			if find(i) == find(target) { // 和不喜欢的人在相同组，失败
				return false
			}
			union(graph[i][0], target) // 不喜欢的人在同一组
		}
	}
	return true
}

var fa []int

// 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// 查询
func find(x int) int {
	for x != fa[x] {
		fa[x] = fa[fa[x]]
		x = fa[x]
	}
	return x
}

// 合并
func union(i, j int) {
	fa[find(i)] = find(j)
}
