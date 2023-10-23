package main

import "fmt"

func main() {
	fmt.Println(minimumHammingDistance([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}}))
}

// leetcode1722_执行交换操作后的最小汉明距离
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	fa = Init(n)
	for i := 0; i < len(allowedSwaps); i++ {
		a, b := allowedSwaps[i][0], allowedSwaps[i][1]
		union(a, b)
	}
	m := make(map[int]map[int]int)
	res := 0
	for i := 0; i < len(source); i++ {
		v := find(i)
		if m[v] == nil {
			m[v] = make(map[int]int)
		}
		m[v][source[i]]++
	}
	for i := 0; i < len(target); i++ {
		v := find(i)
		if m[v][target[i]] == 0 {
			res++
		} else {
			m[v][target[i]]--
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
