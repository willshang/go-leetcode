package main

func main() {

}

var res int
var arr [][][2]int
var visited []bool

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	arr = make([][][2]int, n) // 邻接表
	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1], edges[i][2]
		arr[a] = append(arr[a], [2]int{b, c})
		arr[b] = append(arr[b], [2]int{a, c})
	}
	visited = make([]bool, n)
	visited[0] = true
	res = 0
	dfs(values, maxTime, 0, 0, values[0]) // 初始带着0点的价值
	return res
}

func dfs(values []int, maxTime int, start int, t int, sum int) {
	if start == 0 {
		res = max(res, sum)
	}
	for i := 0; i < len(arr[start]); i++ {
		next, c := arr[start][i][0], arr[start][i][1]
		if t+c <= maxTime { // 时间在范围内
			if visited[next] == false { // 该点没有出现过，加上该点的价值
				visited[next] = true
				dfs(values, maxTime, next, t+c, sum+values[next])
				visited[next] = false
			} else { // 该点出现过，不加价值，只加时间
				dfs(values, maxTime, next, t+c, sum)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
