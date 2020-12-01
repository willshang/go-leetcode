package main

import "sort"

func main() {

}

// leetcode1665_完成所有任务的最少初始能量
func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})
	total := 0
	res := 0
	for i := 0; i < len(tasks); i++ {
		res = max(res, total+tasks[i][1])
		total = total + tasks[i][0]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
