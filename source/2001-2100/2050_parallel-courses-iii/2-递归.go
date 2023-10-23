package main

func main() {

}

var m map[int]int
var arr [][]int
var temp []int

func minimumTime(n int, relations [][]int, time []int) int {
	arr = make([][]int, n+1)
	degree := make([]int, n+1) // 出度
	temp = make([]int, n+1)
	m = make(map[int]int)
	copy(temp, time)
	for i := 0; i < len(relations); i++ {
		a, b := relations[i][0], relations[i][1] // a => b
		arr[b] = append(arr[b], a)
		degree[a]++
	}
	res := 0
	for i := 1; i <= n; i++ {
		if degree[i] == 0 { // 从出度为0的节点往前递归
			res = max(res, dfs(i))
		}
	}
	return res
}

func dfs(root int) int {
	if _, ok := m[root]; ok {
		return m[root]
	}
	sum := 0
	for i := 0; i < len(arr[root]); i++ {
		sum = max(sum, dfs(arr[root][i]))
	}
	sum = sum + temp[root-1]
	m[root] = sum
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
