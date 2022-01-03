package main

import "fmt"

func main() {
	a := &UnionFind{}
	a.Init(10)
	fmt.Println(a.count)
	fmt.Println(a.fa)
}

type UnionFind struct {
	fa    []int
	count int
}

// 初始化
func (u *UnionFind) Init(n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	u.count = n
	u.fa = arr
}

// 查询
func (u UnionFind) find(x int) int {
	if u.fa[x] == x {
		return x
	}
	// 路径压缩
	u.fa[x] = u.find(u.fa[x])
	return u.fa[x]
}

// 合并
func (u *UnionFind) union(i, j int) {
	x, y := u.find(i), u.find(j)
	if x != y {
		u.fa[x] = y
		u.count--
	}
}

func (u UnionFind) query(i, j int) bool {
	return u.find(i) == u.find(j)
}
