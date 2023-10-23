package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

// leetcode207_课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	path := make([]int, 0)
	inEdges := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		// prev->cur
		prev := prerequisites[i][1]
		cur := prerequisites[i][0]
		edges[prev] = append(edges[prev], cur)
		inEdges[cur]++ // 入度
	}
	// 入度为0
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inEdges[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		path = append(path, start)
		for i := 0; i < len(edges[start]); i++ {
			out := edges[start][i]
			inEdges[out]--
			if inEdges[out] == 0 {
				queue = append(queue, out)
			}
		}
	}
	return len(path) == numCourses
}
