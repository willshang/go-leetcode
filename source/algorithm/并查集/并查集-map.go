package main

import (
	"fmt"
	"math"
)

func main() {
	fa = Init([]int{1, 3, 4, 6, 7, 8, 9, 10})
	fmt.Println(fa)
	union(3, 1)
	union(1, 4)
	fmt.Println(find(3))
	fmt.Println(query(3, 4))
	fmt.Println(query(3, 5))
}

var fa map[int]int

// 初始化
func Init(data []int) map[int]int {
	n := len(data)
	arr := make(map[int]int)
	for i := 0; i < n; i++ {
		arr[data[i]] = data[i]
	}
	return arr
}

// 查询
func find(x int) int {
	if _, ok := fa[x]; !ok {
		return math.MinInt32 // 特殊处理
	}
	res := x
	for res != fa[res] {
		res = fa[res]
	}
	return res
}

// 合并
func union(i, j int) {
	x, y := find(i), find(j)
	if x == y {
		return
	} else if x == math.MinInt32 || y == math.MinInt32 {
		return
	}
	fa[x] = y
}

func query(i, j int) bool {
	return find(i) == find(j)
}
