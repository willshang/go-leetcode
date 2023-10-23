package main

import "sort"

func main() {

}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	arr := make([][3]int, 0) // 相邻格子可以连一条边，高度差绝对值最为边的权值
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			index := i*m + j // 转化为一维坐标
			if i > 0 {
				value := abs(heights[i][j] - heights[i-1][j])      // 上到下
				arr = append(arr, [3]int{index - m, index, value}) // 之前坐标，当前坐标，绝对值
			}
			if j > 0 {
				value := abs(heights[i][j] - heights[i][j-1])      // 左到右
				arr = append(arr, [3]int{index - 1, index, value}) // 之前坐标，当前坐标，绝对值
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][2] < arr[j][2]
	})
	fa = Init(n * m)
	for i := 0; i < len(arr); i++ { // 从小到大枚举高度差绝对值
		a, b, c := arr[i][0], arr[i][1], arr[i][2]
		union(a, b)
		if query(0, n*m-1) == true {
			return c
		}
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
