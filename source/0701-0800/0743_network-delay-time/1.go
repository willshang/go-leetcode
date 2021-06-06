package main

import "math"

func main() {

}

func networkDelayTime(times [][]int, n int, k int) int {
	m := make(map[int]map[int]int)
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt32
		m[i] = make(map[int]int)
	}
	for i := 0; i < len(times); i++ {
		a, b, c := times[i][0], times[i][1], times[i][2] // a=>b
		m[a][b] = c
	}
	res := 0

	return res
}

func dfs() {

}
