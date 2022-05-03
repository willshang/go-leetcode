package main

func main() {

}

// leetcode2246_相邻字符不同的最长路径
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
	return res + 1 // 加上自身
}

func dfs(x int, s string) int {
	result := 0 // 以该节点单边路径最大长度
	for i := 0; i < len(arr[x]); i++ {
		value := dfs(arr[x][i], s) + 1 // 子节点的最长路径
		if s[x] != s[arr[x][i]] {
			res = max(res, result+value) // 更新：子树的最长路径+子树的最长路径
			result = max(result, value)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
