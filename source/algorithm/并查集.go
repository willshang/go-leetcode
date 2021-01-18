package main

import "fmt"

func main() {
	fa = Init(10)
	union(3, 1)
	union(1, 4)
	fmt.Println(find(3))
	fmt.Println(query(3, 4))
	fmt.Println(query(3, 5))
	fmt.Println(fa)
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
