package main

func main() {

}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	fa = Init(n * n * 4)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			index := 4 * (i*n + j) // 扩大4倍，每个格子拆分为4个：0、1、2、3（顺时针）
			if grid[i][j] == '/' {
				union(index, index+3)   // 合并0、3
				union(index+1, index+2) // 合并1、2
			} else if grid[i][j] == '\\' {
				union(index, index+1)   // 合并0、1
				union(index+2, index+3) // 合并2、3
			} else {
				union(index, index+1) // 合并0、1、2、3
				union(index+1, index+2)
				union(index+2, index+3)
			}
			if j < n-1 {
				union(index+1, index+7) // 向右合并
			}
			if i < n-1 {
				union(index+2, index+4*n) // 向下合并
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
