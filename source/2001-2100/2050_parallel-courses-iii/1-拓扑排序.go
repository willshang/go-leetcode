package main

func main() {

}

// leetcode2050_并行课程III
func minimumTime(n int, relations [][]int, time []int) int {
	res := 0
	degree := make([]int, n+1)
	dis := make([]int, n+1)   // 计算入度
	arr := make([][]int, n+1) // 邻接表
	for i := 0; i < len(relations); i++ {
		a, b := relations[i][0], relations[i][1] // a => b
		arr[a] = append(arr[a], b)
		degree[b]++ // 入度+1
	}
	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if degree[i] == 0 { // 入度为0：起点
			queue = append(queue, i)
			dis[i] = time[i-1]
			res = max(res, dis[i])
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < len(arr[cur]); i++ {
			next := arr[cur][i]
			dis[next] = max(dis[next], dis[cur]+time[next-1]) // 更新为较大的结果
			degree[next]--                                    // next节点入度-1
			if degree[next] == 0 {                            // 入度=0
				queue = append(queue, next)
			}
			res = max(res, dis[next])
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
