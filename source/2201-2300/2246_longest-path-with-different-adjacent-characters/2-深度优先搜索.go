package main

func main() {

}

var arr [][]int
var res int

func longestPath(parent []int, s string) int {
	res = 0
	n := len(parent)
	arr = make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		arr[p] = append(arr[p], i)
	}
	dfs(0, s)
	return res
}

func dfs(x int, s string) int {
	var a, b int // 2个节点值 a>b
	for i := 0; i < len(arr[x]); i++ {
		v := dfs(arr[x][i], s)
		if s[x] != s[arr[x][i]] {
			if v > a {
				a, b = v, a
			} else if v > b {
				b = v
			}
		}
	}
	res = max(res, 1+a+b)
	return a + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
