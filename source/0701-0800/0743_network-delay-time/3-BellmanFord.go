package main

import "math"

func main() {

}

func networkDelayTime(times [][]int, n int, k int) int {
	maxValue := math.MaxInt32 / 10
	dis := make([]int, n) // k到其他点的距离
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[k-1] = 0
	for k := 0; k < n-1; k++ {
		flag := true
		for i := 0; i < len(times); i++ {
			a, b, c := times[i][0]-1, times[i][1]-1, times[i][2] // a=>b
			if dis[a]+c < dis[b] {
				flag = false
				dis[b] = dis[a] + c
			}
		}
		if flag == true {
			break
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if dis[i] == maxValue {
			return -1
		}
		res = max(res, dis[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
