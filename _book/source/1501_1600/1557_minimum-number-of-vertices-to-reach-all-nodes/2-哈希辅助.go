package main

import "fmt"

func main() {
	fmt.Println(findSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	res := make([]int, 0)
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[i] = true
	}
	for i := 0; i < len(edges); i++ {
		b := edges[i][1]
		delete(m, b)
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}
