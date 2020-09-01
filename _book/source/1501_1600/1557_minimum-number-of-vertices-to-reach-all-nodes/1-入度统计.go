package main

import "fmt"

func main() {
	fmt.Println(findSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
}

// leetcode1557_可以到达所有点的最少点数目
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	res := make([]int, 0)
	inEdges := make([]int, n)
	for i := 0; i < len(edges); i++ {
		// a->b
		b := edges[i][1]
		inEdges[b]++ // 入度
	}
	for i := 0; i < len(inEdges); i++ {
		if inEdges[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
