package main

func main() {

}

func validPath(n int, edges [][]int, source int, destination int) bool {
	fa = Init(n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1] // a<=>b
		union(a, b)
	}
	return query(source, destination)
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
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}

// 合并
func union(i, j int) {
	fa[find(i)] = find(j)
}

func query(i, j int) bool {
	return find(i) == find(j)
}
