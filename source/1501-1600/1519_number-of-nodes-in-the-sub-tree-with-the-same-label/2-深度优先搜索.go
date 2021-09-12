package main

func main() {

}

var arr [][]int   // 邻接表
var res []int     // 结果
var m map[int]int // 临时结果
func countSubTrees(n int, edges [][]int, labels string) []int {
	res = make([]int, n)
	arr = make([][]int, n)
	m = make(map[int]int)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	dfs(labels, 0, 0)
	for i := 0; i < n; i++ {
		value := int(labels[i] - 'a')
		res[i] = m[i*26+value]
	}
	return res
}

// 后序遍历
func dfs(s string, cur int, prev int) {
	m[cur*26+int(s[cur]-'a')]++
	for i := 0; i < len(arr[cur]); i++ {
		next := arr[cur][i]
		if next == prev { // 注意避免重复往回走
			continue
		}
		dfs(s, next, cur)
		for j := 0; j < 26; j++ { // 子节点的结果
			m[cur*26+j] = m[cur*26+j] + m[next*26+j]
		}
	}
}
