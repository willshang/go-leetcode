package main

import (
	"math"
)

func main() {

}

func networkDelayTime(times [][]int, n int, k int) int {
	maxValue := math.MaxInt32 / 10
	arr := make([][]int, n) // 邻接表：i=>j的集合
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			arr[i][j] = maxValue
		}
	}
	for i := 0; i < len(times); i++ {
		a, b, c := times[i][0]-1, times[i][1]-1, times[i][2] // a=>b
		arr[a][b] = c
	}
	for p := 0; p < n; p++ { // floyd
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				arr[i][j] = min(arr[i][j], arr[i][p]+arr[p][j])
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if arr[k-1][i] == maxValue {
			return -1
		}
		res = max(res, arr[k-1][i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
