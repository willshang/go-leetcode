package main

import "fmt"

func main() {

}

var arr []map[int]bool
var m map[string]bool

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	res := make([]bool, len(queries))
	m = make(map[string]bool)
	arr = make([]map[int]bool, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		if arr[a] == nil {
			arr[a] = make(map[int]bool)
		}
		arr[a][b] = true // a=>b
	}
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		res[i] = dfs(a, b) // a=>b
	}
	return res
}

func dfs(i, target int) bool {
	status := fmt.Sprintf("%d,%d", i, target)
	if value, ok := m[status]; ok {
		return value
	}
	res := false
	if arr[i][target] == true {
		res = true
	} else {
		for k := range arr[i] {
			if dfs(k, target) == true {
				res = true
				break
			}
		}
	}
	m[status] = res
	return res
}
