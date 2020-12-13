package main

import "fmt"

func main() {
	fa = Init(10)
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			count++
		}
	}
	union(3, 1)
	union(1, 4)
	fmt.Println(find(3))
	fmt.Println(query(3, 4))
	fmt.Println(query(3, 5))
	fmt.Println(fa)
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
