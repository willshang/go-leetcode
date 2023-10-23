package main

import "sort"

func main() {

}

// leetcode721_账户合并
func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	fa = Init(n)
	res := make([][]string, 0)
	m := make(map[string]int)
	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			email := accounts[i][j]
			if id, ok := m[email]; ok { // 邮箱重复出现，合并账户
				union(i, id)
			} else {
				m[email] = i
			}
		}
	}
	temp := make([][]string, n)
	for k, v := range m {
		target := find(v)
		temp[target] = append(temp[target], k)
	}
	for i := 0; i < len(temp); i++ {
		if len(temp[i]) > 0 {
			arr := temp[i]
			sort.Strings(arr)
			arr = append([]string{accounts[i][0]}, arr...)
			res = append(res, arr)
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
