package main

import (
	"math"
)

func main() {

}

var res int

func minimumTimeRequired(jobs []int, k int) int {
	res = math.MaxInt32
	dfs(jobs, make([]int, k), 0, 0, 0)
	return res
}

// index => job的下标；count => 已经分配工人数组的个数
func dfs(jobs []int, arr []int, index int, maxValue int, count int) {
	if maxValue > res { // 剪枝
		return
	}
	if index == len(jobs) {
		res = maxValue
		return
	}
	if count < len(arr) { // 分配给空闲的
		arr[count] = jobs[index]
		dfs(jobs, arr, index+1, max(maxValue, arr[count]), count+1)
		arr[count] = 0
	}
	for i := 0; i < count; i++ { // 分配给非空闲的
		arr[i] = arr[i] + jobs[index]
		dfs(jobs, arr, index+1, max(maxValue, arr[i]), count)
		arr[i] = arr[i] - jobs[index]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
