package main

import "fmt"

func main() {
	//fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))
	fmt.Println(shortestAlternatingPaths(5,
		[][]int{
			{3, 2},
			{4, 1},
			{1, 4}},
		[][]int{
			{2, 3},
			{0, 4},
			{4, 3},
			{4, 4},
			{4, 0},
			{1, 0},
		}))
}

// leetcode1129_颜色交替的最短路径
func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	redArr, blueArr := make([][]int, n), make([][]int, n)
	for i := 0; i < len(red_edges); i++ {
		a, b := red_edges[i][0], red_edges[i][1]
		redArr[a] = append(redArr[a], b)
	}
	for i := 0; i < len(blue_edges); i++ {
		a, b := blue_edges[i][0], blue_edges[i][1]
		blueArr[a] = append(blueArr[a], b)
	}
	res := make([]int, n) // res[0] = 0
	for i := 1; i < n; i++ {
		res[i] = -1
	}
	queue, visited := make([][2]int, 0), make([][2]bool, n)
	for i := 0; i < len(redArr[0]); i++ {
		queue = append(queue, [2]int{redArr[0][i], 0})
	}
	for i := 0; i < len(blueArr[0]); i++ {
		queue = append(queue, [2]int{blueArr[0][i], 1})
	}
	count := 1
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node, targetColor := queue[i][0], queue[i][1]
			if res[node] == -1 {
				res[node] = count
			}
			if targetColor == 0 && visited[node][targetColor] == false {
				visited[node][targetColor] = true
				for j := 0; j < len(blueArr[node]); j++ {
					queue = append(queue, [2]int{blueArr[node][j], 1})
				}
			}
			if targetColor == 1 && visited[node][targetColor] == false {
				visited[node][targetColor] = true
				for j := 0; j < len(redArr[node]); j++ {
					queue = append(queue, [2]int{redArr[node][j], 0})
				}
			}
		}
		queue = queue[length:]
		count++
	}
	return res
}
