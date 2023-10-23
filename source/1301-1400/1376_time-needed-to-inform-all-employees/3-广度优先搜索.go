package main

import "fmt"

func main() {
	fmt.Println(numOfMinutes(4, 2, []int{3, 3, -1, 2}, []int{0, 0, 162, 914}))
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	res := 0
	m := make(map[int][]int)
	for i := 0; i < len(manager); i++ {
		if _, ok := m[manager[i]]; ok {
			m[manager[i]] = append(m[manager[i]], i)
		} else {
			m[manager[i]] = []int{i}
		}
	}
	queue := make([]int, 0)
	queue = append(queue, headID)
	costM := make(map[int]int)
	costM[headID] = 0
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		res = max(res, costM[id])
		for i := 0; i < len(m[id]); i++ {
			costM[m[id][i]] = informTime[id] + costM[id]
			queue = append(queue, m[id][i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
