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

var redArr [][]int
var blueArr [][]int
var res []int

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	redArr = make([][]int, n)
	blueArr = make([][]int, n)
	for i := 0; i < len(red_edges); i++ {
		a, b := red_edges[i][0], red_edges[i][1]
		redArr[a] = append(redArr[a], b)
	}
	for i := 0; i < len(blue_edges); i++ {
		a, b := blue_edges[i][0], blue_edges[i][1]
		blueArr[a] = append(blueArr[a], b)
	}
	res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	res[0] = 0
	bfs(n, 0) // 红
	bfs(n, 1) // 蓝
	return res
}

func bfs(n int, color int) {
	visited := make([][2]bool, n)
	visited[0][color] = true
	queue := make([]int, 0)
	queue = append(queue, 0)
	count := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			targetColor := count % 2
			if targetColor == color { //偶数次和当前颜色一致
				for i := 0; i < len(redArr[node]); i++ {
					next := redArr[node][i]
					if visited[next][targetColor] == false {
						visited[next][targetColor] = true
						if res[next] == -1 || res[next] > count+1 {
							res[next] = count + 1
						}
						queue = append(queue, next)
					}
				}
			} else {
				for i := 0; i < len(blueArr[node]); i++ {
					next := blueArr[node][i]
					if visited[next][targetColor] == false {
						visited[next][targetColor] = true
						if res[next] == -1 || res[next] > count+1 {
							res[next] = count + 1
						}
						queue = append(queue, next)
					}
				}
			}
		}
		queue = queue[length:]
		count++
	}
}
