package main

func main() {

}

// leetcode_lcp20_快速公交
var mod = 1000000007

var visited map[int]int

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	visited = make(map[int]int)
	return dfs(inc, dec, jump, cost, target) % mod
}

func dfs(inc int, dec int, jump []int, cost []int, cur int) int {
	if cur == 0 {
		return 0
	}
	if cur == 1 {
		return inc
	}
	if visited[cur] > 0 {
		return visited[cur]
	}
	res := cur * inc // 最坏的情况：全+1
	for i := 0; i < len(jump); i++ {
		diff, next := cur%jump[i], cur/jump[i]
		if diff == 0 { // 直接坐公交
			res = min(res, dfs(inc, dec, jump, cost, next)+cost[i])
		} else {
			// 向左走坐公交
			res = min(res, dfs(inc, dec, jump, cost, next)+cost[i]+diff*inc)
			// 向右走坐公交
			res = min(res, dfs(inc, dec, jump, cost, next+1)+cost[i]+(jump[i]-diff)*dec)
		}
	}
	visited[cur] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
