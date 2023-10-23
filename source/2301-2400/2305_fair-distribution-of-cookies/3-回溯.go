package main

import "math"

func main() {

}

// leetcode2305_公平分发饼干
var res int

func distributeCookies(cookies []int, k int) int {
	res = math.MaxInt32
	dfs(cookies, make([]int, k), 0, 0, 0)
	return res
}

// index => cookies的下标；count => 已经分配的个数
func dfs(cookies []int, arr []int, index int, maxValue int, count int) {
	if maxValue > res { // 剪枝
		return
	}
	if index == len(cookies) {
		res = maxValue
		return
	}
	if count < len(arr) {
		arr[count] = cookies[index]
		dfs(cookies, arr, index+1, max(maxValue, arr[count]), count+1)
		arr[count] = 0
	}
	for i := 0; i < count; i++ {
		arr[i] = arr[i] + cookies[index]
		dfs(cookies, arr, index+1, max(maxValue, arr[i]), count)
		arr[i] = arr[i] - cookies[index]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
