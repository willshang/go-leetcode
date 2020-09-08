package main

import "fmt"

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
}

// leetcode1306_跳跃游戏III
var m map[int]bool

func canReach(arr []int, start int) bool {
	m = make(map[int]bool)
	return dfs(arr, start)
}

func dfs(arr []int, i int) bool {
	if i < 0 || i > len(arr)-1 || m[i] == true {
		return false
	}
	m[i] = true
	return arr[i] == 0 || dfs(arr, i+arr[i]) || dfs(arr, i-arr[i])
}
