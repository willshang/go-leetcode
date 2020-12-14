package main

func main() {

}

// leetcode684_冗余连接
func findRedundantConnection(edges [][]int) []int {
	n := len(edges) + 1
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		if find(fa, a) == find(fa, b) {
			return edges[i]
		}
		union(fa, a, b)
	}
	return nil
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
