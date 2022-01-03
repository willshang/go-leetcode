package main

import "fmt"

func main() {
	fa, rank = Init(10)
	union(3, 1)
	union(1, 4)
	fmt.Println(find(3))
	fmt.Println(query(3, 4))
	fmt.Println(query(3, 5))
	fmt.Println(fa, rank)
}

var fa []int
var rank []int

// 初始化
func Init(n int) ([]int, []int) {
	arr := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		r[i] = 1
	}
	return arr, r
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
	// 按秩合并
	x, y := find(i), find(j)
	if rank[x] <= rank[y] {
		fa[x] = y
	} else {
		fa[y] = x
	}
	if rank[x] == rank[y] && x != y {
		rank[y]++ // 如果深度相同且根节点不同，则新的根节点的深度+1
	}
}

func query(i, j int) bool {
	return find(i) == find(j)
}
