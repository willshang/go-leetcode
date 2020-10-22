package main

import "fmt"

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

var res bool
var visited []int
var path []int
var edges [][]int

func findOrder(numCourses int, prerequisites [][]int) []int {
	res = true
	edges = make([][]int, numCourses) // 邻接表
	visited = make([]int, numCourses)
	path = make([]int, 0)
	for i := 0; i < len(prerequisites); i++ {
		// prev->cur
		prev := prerequisites[i][1]
		cur := prerequisites[i][0]
		edges[prev] = append(edges[prev], cur)
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
		if res == false {
			return nil
		}
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}
	return path
}

func dfs(start int) {
	// 0 未搜索
	// 1 搜索中
	// 2 已完成
	visited[start] = 1
	for i := 0; i < len(edges[start]); i++ {
		out := edges[start][i]
		if visited[out] == 0 {
			dfs(out)
			if res == false {
				return
			}
		} else if visited[out] == 1 {
			res = false
			return
		}
	}
	visited[start] = 2
	path = append(path, start)
}
