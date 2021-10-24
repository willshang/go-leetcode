package main

import "sort"

func main() {

}

// leetcode2033_获取单值网格的最小操作数
func minOperations(grid [][]int, x int) int {
	n, m := len(grid), len(grid[0])
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr = append(arr, grid[i][j])
		}
	}
	sort.Ints(arr)
	target := arr[len(arr)/2]
	res := 0
	for i := 0; i < len(arr); i++ {
		if abs(target-arr[i])%x != 0 {
			return -1
		}
		res = res + abs(target-arr[i])/x
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
