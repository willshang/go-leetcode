package main

import "fmt"

func main() {
	fmt.Println(countArrangement(2))
	fmt.Println(countArrangement(15))
}

var res [][]int

func countArrangement(n int) int {
	res = make([][]int, 0)
	dfs(n, make([]int, 0), make([]bool, n+1))
	fmt.Println(res)
	return len(res)
}

func dfs(n int, path []int, visited []bool) {
	if len(path) == n {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := 1; i <= n; i++ {
		index := len(path) + 1
		if visited[i] == false && (i%index == 0 || index%i == 0) {
			visited[i] = true
			dfs(n, append(path, i), visited)
			visited[i] = false
		}
	}
}
