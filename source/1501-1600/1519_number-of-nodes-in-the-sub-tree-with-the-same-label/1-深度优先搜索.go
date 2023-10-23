package main

func main() {

}

// leetcode1519_子树中标签相同的节点数
var arr [][]int // 邻接表
var res []int   // 结果

func countSubTrees(n int, edges [][]int, labels string) []int {
	res = make([]int, n)
	arr = make([][]int, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	dfs(labels, 0, 0)
	return res
}

// 后序遍历
func dfs(s string, cur int, prev int) [26]int {
	count := [26]int{}
	for i := 0; i < len(arr[cur]); i++ {
		next := arr[cur][i]
		if next == prev { // 注意避免重复往回走
			continue
		}
		res := dfs(s, next, cur)
		for j := 0; j < len(res); j++ {
			count[j] = count[j] + res[j]
		}
	}
	value := int(s[cur] - 'a')
	count[value]++
	res[cur] = count[value]
	return count
}
