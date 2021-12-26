package main

func main() {

}

// leetcode2097_合法重新排列数对
var arr map[int][]int
var path []int

func validArrangement(pairs [][]int) [][]int {
	path = make([]int, 0)
	arr = make(map[int][]int) // 有向图邻接表
	m := make(map[int]int)
	for i := 0; i < len(pairs); i++ {
		a, b := pairs[i][0], pairs[i][1]
		arr[a] = append(arr[a], b)
		m[b]++ // 入度+1
		m[a]-- // 出度-1
	}
	start := pairs[0][0] // 寻找起始节点
	for k, v := range m {
		if v == -1 {
			start = k
			break
		}
	}
	dfs(start)
	res := make([][]int, 0)
	for i := len(path) - 1; i > 0; i-- {
		res = append(res, []int{path[i], path[i-1]})
	}
	return res
}

func dfs(start int) {
	for len(arr[start]) > 0 {
		next := arr[start][0]
		arr[start] = arr[start][1:]
		dfs(next)
	}
	path = append(path, start)
}
