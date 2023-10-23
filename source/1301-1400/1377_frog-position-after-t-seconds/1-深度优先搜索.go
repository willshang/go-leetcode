package main

func main() {

}

var arr [][]int
var res []float64

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	arr = make([][]int, n+1)
	res = make([]float64, n+1)
	res[1] = 1
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	visited := make([]bool, n+1) // 已经访问过的
	visited[1] = true
	dfs(1, t, visited)
	return res[target]
}

func dfs(start int, t int, visited []bool) {
	if t <= 0 {
		return
	}
	count := 0
	for i := 0; i < len(arr[start]); i++ {
		next := arr[start][i]
		if visited[next] == false {
			count++
		}
	}
	if count == 0 {
		return
	}
	per := res[start] / float64(count) // 每一跳的概率
	for i := 0; i < len(arr[start]); i++ {
		next := arr[start][i]
		if visited[next] == false {
			visited[next] = true
			res[start] = res[start] - per // start-per
			res[next] = res[next] + per   // next+per
			dfs(next, t-1, visited)
			visited[next] = false
		}
	}
}
