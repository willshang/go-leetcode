package main

import "fmt"

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

// 剑指OfferII113.课程顺序
func findOrder(numCourses int, prerequisites [][]int) []int {
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
	if len(path) != numCourses {
		return nil
	}
	return path
}
