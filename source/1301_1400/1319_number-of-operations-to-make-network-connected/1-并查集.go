package main

func main() {

}

// leetcode1319_连通网络的操作次数
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	res := n - 1
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	for i := 0; i < len(connections); i++ {
		a, b := connections[i][0], connections[i][1]
		if find(fa, a) != find(fa, b) {
			union(fa, a, b)
			res--
		}
	}
	return res
}

func union(fa []int, a, b int) {
	fa[find(fa, a)] = find(fa, b)
}

func find(fa []int, a int) int {
	for fa[a] != a {
		fa[a] = fa[fa[a]]
		a = fa[a]
	}
	return a
}
