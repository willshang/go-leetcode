package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(peopleIndexes([][]string{
		{"leetcode", "google", "facebook"},
		{"google", "microsoft"},
		{"google", "facebook"},
		{"google"},
		{"amazon"},
	}))
}

// leetcode1452_收藏清单
type Node struct {
	index int
	str   []string
}

func peopleIndexes(favoriteCompanies [][]string) []int {
	n := len(favoriteCompanies)
	arr := make([]Node, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, Node{
			index: i,
			str:   favoriteCompanies[i],
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i].str) < len(arr[j].str)
	})
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		flag := true
		for j := i + 1; j < n; j++ {
			if judge(arr[i].str, arr[j].str) == true {
				flag = false
				break
			}
		}
		if flag == true {
			res = append(res, arr[i].index)
		}
	}
	sort.Ints(res)
	return res
}

func judge(a, b []string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(a); i++ {
		m[a[i]] = true
	}
	for i := 0; i < len(b); i++ {
		if _, ok := m[b[i]]; ok {
			delete(m, b[i])
		}
	}
	return len(m) == 0
}
