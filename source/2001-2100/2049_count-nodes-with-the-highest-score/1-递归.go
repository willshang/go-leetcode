package main

func main() {

}

var arr [][]int
var sum []int

func countHighestScoreNodes(parents []int) int {
	res := 0
	maxValue := 0
	n := len(parents)
	sum = make([]int, n)
	arr = make([][]int, n)
	for i := 1; i < n; i++ { // 邻接表
		a, b := i, parents[i] // b=>a
		arr[b] = append(arr[b], a)
	}
	dfs(0)                   // 先统计
	for i := 0; i < n; i++ { // 后计算
		value := 1
		for j := 0; j < len(arr[i]); j++ { // 计算左右子树乘积
			if sum[arr[i][j]] > 0 {
				value = value * sum[arr[i][j]]
			}
		}
		if parents[i] != -1 { // 计算去掉以根节点位i的剩余部分
			if sum[parents[i]]-sum[i] > 0 {
				value = value * (sum[0] - sum[i]) // 根节点总数-节点i总数
			}
		}
		if value > maxValue {
			maxValue = value
			res = 1
		} else if value == maxValue {
			res++
		}
	}
	return res
}

func dfs(root int) int {
	count := 1
	for i := 0; i < len(arr[root]); i++ {
		count = count + dfs(arr[root][i])
	}
	sum[root] = count
	return count
}
