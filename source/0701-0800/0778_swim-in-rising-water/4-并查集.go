package main

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func swimInWater(grid [][]int) int {
	n := len(grid)
	arr := make([][2]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[grid[i][j]] = [2]int{i, j} // 高度对应的位置
		}
	}
	fa = Init(n * n)
	for i := 0; i < len(arr); i++ {
		a, b := arr[i][0], arr[i][1]
		for j := 0; j < 4; j++ {
			x, y := a+dx[j], b+dy[j]
			if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] <= i {
				union(x*n+y, a*n+b)
			}
		}
		if query(0, n*n-1) {
			return i
		}
	}
	return 0
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
