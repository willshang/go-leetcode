package main

import "fmt"

func main() {
	arr := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(arr))
}

// leetcode200_岛屿数量
func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	fa = Init(n*m + 1)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				grid[i][j] = '0'
				if i >= 1 && grid[i-1][j] == '1' {
					union(i*m+j, (i-1)*m+j)
				}
				if i < n-1 && grid[i+1][j] == '1' {
					union(i*m+j, (i+1)*m+j)
				}
				if j >= 1 && grid[i][j-1] == '1' {
					union(i*m+j, i*m+j-1)
				}
				if j < m-1 && grid[i][j+1] == '1' {
					union(i*m+j, i*m+j+1)
				}
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
	count = 0
	for i := 0; i < n; i++ {
		arr[i] = i
	}
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
