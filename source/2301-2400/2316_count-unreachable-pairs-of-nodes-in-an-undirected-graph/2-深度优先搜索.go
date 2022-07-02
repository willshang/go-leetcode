package main

func main() {

}

// leetcode2316_统计无向图中无法互相到达点对数
var arr [][]int
var visited []bool
var count int

func countPairs(n int, edges [][]int) int64 {
	res := int64(0)
	arr = make([][]int, n)
	visited = make([]bool, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}

	for i := 0; i < n; i++ {
		if visited[i] == false {
			count = 0
			dfs(i)
			res = res + int64(count)*int64(n-count)
		}
	}
	return res / 2
}

func dfs(start int) {
	visited[start] = true
	count++
	for i := 0; i < len(arr[start]); i++ {
		next := arr[start][i]
		if visited[next] == false {
			dfs(next)
		}
	}
}
