package main

import "fmt"

func main() {
	fmt.Println(edgeScore([]int{1, 2, 0}))
}

func edgeScore(edges []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		v := edges[i]
		m[v] = m[v] + i
		if m[v] > m[res] || m[v] == m[res] && v < res {
			res = v
		}
	}
	return res
}
