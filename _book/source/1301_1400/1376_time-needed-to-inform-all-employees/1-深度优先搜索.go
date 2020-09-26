package main

import "fmt"

func main() {
	fmt.Println(numOfMinutes(4, 2, []int{3, 3, -1, 2}, []int{0, 0, 162, 914}))
}

// leetcode1376_通知所有员工所需的时间
var res int
var m map[int][]int

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	m = make(map[int][]int)
	for i := 0; i < len(manager); i++ {
		if _, ok := m[manager[i]]; ok {
			m[manager[i]] = append(m[manager[i]], i)
		} else {
			m[manager[i]] = []int{i}
		}
	}
	res = 0
	dfs(headID, 0, informTime)
	return res
}

func dfs(headID int, cost int, informTime []int) {
	arr, ok := m[headID]
	if !ok {
		if cost > res {
			res = cost
		}
		return
	}
	cost = cost + informTime[headID]
	for i := 0; i < len(arr); i++ {
		dfs(arr[i], cost, informTime)
	}
}
