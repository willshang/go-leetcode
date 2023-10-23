package main

import "fmt"

func main() {
	fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
}

func possibleBipartition(n int, dislikes [][]int) bool {
	fa = Init(n + 1)
	arr := make([][]int, n+1)
	for i := 0; i < len(dislikes); i++ {
		a, b := dislikes[i][0], dislikes[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < len(arr[i]); j++ {
			target := arr[i][j]
			if find(i) == find(target) { // 和不喜欢的人在相同组，失败
				return false
			}
			union(arr[i][0], target) // 不喜欢的人在同一组
		}
	}
	return true
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
	for x != fa[x] {
		fa[x] = fa[fa[x]]
		x = fa[x]
	}
	return x
}

// 合并
func union(i, j int) {
	fa[find(i)] = find(j)
}
