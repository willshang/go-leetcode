package main

func main() {

}

var arr map[int][]int
var res [][]int

func validArrangement(pairs [][]int) [][]int {
	res = make([][]int, 0)
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
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func dfs(start int) {
	for len(arr[start]) > 0 {
		next := arr[start][0]
		arr[start] = arr[start][1:]
		dfs(next)
		res = append(res, []int{start, next})
	}
}
