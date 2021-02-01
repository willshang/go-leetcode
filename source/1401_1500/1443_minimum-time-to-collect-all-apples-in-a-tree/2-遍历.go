package main

import "fmt"

func main() {
	fmt.Println(minTime(4, [][]int{
		{0, 2},
		{0, 3},
		{1, 2},
	}, []bool{false, true, false, false}))
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	ans := 0
	m := make([]bool, n)
	m[0] = true
	for i := 0; i < len(edges); i++ {
		from, to := edges[i][0], edges[i][1]
		if m[from] {
			m[to] = true
		} else {
			m[from] = true
			// 改变数据
			// [[0 2] [0 3] [1 2]] => [[0 2] [0 3] [2 1]]
			edges[i][0], edges[i][1] = edges[i][1], edges[i][0]
		}
	}
	for i := len(edges) - 1; i >= 0; i-- {
		from, to := edges[i][0], edges[i][1]
		if hasApple[to] {
			hasApple[from] = true
			ans += 2
		}
	}
	return ans
}
