package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
}

// leetcode1202_交换字符串中的元素
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	for i := 0; i < len(pairs); i++ {
		a, b := pairs[i][0], pairs[i][1]
		union(fa, a, b)
	}
	m := make(map[int][]int)
	for i := 0; i < len(s); i++ {
		target := find(fa, i)
		m[target] = append(m[target], i)
	}
	res := []byte(s)
	for _, v := range m {
		arr := make([]int, 0)
		for i := 0; i < len(v); i++ {
			arr = append(arr, int(s[v[i]]))
		}
		sort.Ints(arr)
		for i := 0; i < len(v); i++ {
			res[v[i]] = byte(arr[i])
		}
	}
	return string(res)
}

func union(fa []int, a, b int) {
	fa[find(fa, a)] = find(fa, b)
}

func find(fa []int, a int) int {
	for fa[a] != a {
		fa[a] = fa[fa[a]]
		a = fa[a]
	}
	return a
}
