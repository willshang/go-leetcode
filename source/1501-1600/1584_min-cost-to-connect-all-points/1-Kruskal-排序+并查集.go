package main

import (
	"sort"
)

func main() {

}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	arr := make([][3]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			arr = append(arr, [3]int{i, j, dis(points[i], points[j])}) // a=>b c
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][2] < arr[j][2]
	})
	return Kruskal(n, arr)
}

func Kruskal(n int, arr [][3]int) int {
	res := 0
	fa = Init(n)
	count := 0
	for i := 0; i < len(arr); i++ {
		a, b, c := arr[i][0], arr[i][1], arr[i][2]
		if query(a, b) == false {
			union(a, b)
			res = res + c
			count++
			if count == n-1 {
				break
			}
		}
	}
	return res
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

func dis(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
