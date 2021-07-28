package main

func main() {

}

// leetcode1462_课程表IV
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	res := make([]bool, len(queries))
	m := make(map[int]map[int]bool)
	for i := 0; i < numCourses; i++ {
		m[i] = make(map[int]bool)
	}
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		m[a][b] = true // a=>b
	}
	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				// ik + kj =>ij
				if m[i][k] == true && m[k][j] == true {
					m[i][j] = true
				}
			}
		}
	}
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		res[i] = m[a][b]
	}
	return res
}
