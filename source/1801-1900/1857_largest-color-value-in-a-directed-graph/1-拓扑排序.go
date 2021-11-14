package main

func main() {

}

// leetcode1857_有向图中最大颜色值
func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	arr := make([][]int, n)  // 邻接表
	degree := make([]int, n) // 入度
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1] // a=>b
		arr[a] = append(arr[a], b)
		degree[b]++
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 0 { // 入度为0
			queue = append(queue, i)
		}
	}
	dp := make([][26]int, n) // dp[i][j] => 到节点i颜色j出现的次数
	count := 0
	for len(queue) > 0 {
		count++ // 节点+1
		node := queue[0]
		queue = queue[1:]
		dp[node][int(colors[node]-'a')]++ // 该节点颜色+1
		for i := 0; i < len(arr[node]); i++ {
			next := arr[node][i]
			degree[next]--
			if degree[next] == 0 {
				queue = append(queue, next)
			}
			// 更新次数
			for j := 0; j < 26; j++ {
				dp[next][j] = max(dp[next][j], dp[node][j])
			}
		}
	}
	if count != n { // 判断环
		return -1
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			res = max(res, dp[i][j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
