package main

func main() {

}

func findCircleNum(M [][]int) int {
	n := len(M)
	fa = Init(n)
	count = n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				union(i, j)
			}
		}
	}
	return getCount()
}

var fa []int
var count int

// 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	count = n
	return arr
}

// 查询
func find(x int) int {
	if fa[x] == x {
		return x
	}
	// 路径压缩
	fa[x] = find(fa[x])
	return fa[x]
}

// 合并
func union(i, j int) {
	x, y := find(i), find(j)
	if x != y {
		fa[x] = y
		count--
	}
}

func query(i, j int) bool {
	return find(i) == find(j)
}

func getCount() int {
	return count
}
