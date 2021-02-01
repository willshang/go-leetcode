package main

func main() {

}

// leetcode1743_从相邻元素对还原数组
var res []int
var m map[int][]int

func restoreArray(adjacentPairs [][]int) []int {
	m = make(map[int][]int)
	for i := 0; i < len(adjacentPairs); i++ {
		a, b := adjacentPairs[i][0], adjacentPairs[i][1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	arr := make([]int, 0)
	for k, v := range m {
		if len(v) == 1 {
			arr = append(arr, k)
		}
	}
	res = make([]int, 0)
	dfs(arr[0], make(map[int]bool))
	return res
}

func dfs(cur int, visited map[int]bool) {
	res = append(res, cur)
	visited[cur] = true
	for i := range m[cur] {
		if visited[m[cur][i]] == false {
			dfs(m[cur][i], visited)
		}
	}
}
