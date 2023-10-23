package main

func main() {

}

func hasValidPath(grid [][]int) bool {
	n, m := len(grid), len(grid[0])
	fa = Init(n * m)
	for i := 0; i < n; i++ { // 放大处理
		for j := 0; j < m; j++ {
			index := i*m + j
			if grid[i][j] == 1 || grid[i][j] == 3 || grid[i][j] == 5 {
				if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) {
					union(index, index-1) // 跟左边相连
				}
			}
			if grid[i][j] == 2 || grid[i][j] == 5 || grid[i][j] == 6 {
				if i > 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
					union(index, index-m) // 跟上边相连
				}
			}
		}
	}
	return query(0, n*m-1)
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
