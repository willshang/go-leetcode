package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{3, 2, 0},
		{3, 1, 0},
		{1, 2, 0},
	}
	pairs := [][]int{
		{0, 1},
		{2, 3},
	}
	fmt.Println(unhappyFriends(4, arr, pairs))
}

// leetcode1583_统计不开心的朋友
func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	arr := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		arr[i] = make(map[int]int)
	}
	for i := 0; i < len(preferences); i++ {
		total := n
		for j := 0; j < len(preferences[i]); j++ {
			arr[i][preferences[i][j]] = total
			total = total - 1
		}
	}
	m := make(map[int]bool)
	for i := 0; i < len(pairs); i++ {
		x, y := pairs[i][0], pairs[i][1]
		x1, y1 := pairs[i][1], pairs[i][0]
		for j := 0; j < len(pairs); j++ {
			if i == j {
				continue
			}
			u, v := pairs[j][0], pairs[j][1]
			u1, v1 := pairs[j][1], pairs[j][0]
			if arr[x][u] > arr[x][y] && arr[u][x] > arr[u][v] {
				m[x] = true
			}
			if arr[x1][u] > arr[x1][y1] && arr[u][x1] > arr[u][v] {
				m[x1] = true
			}
			if arr[x][u1] > arr[x][y] && arr[u1][x] > arr[u1][v1] {
				m[x] = true
			}
			if arr[x1][u1] > arr[x1][y1] && arr[u1][x1] > arr[u1][v1] {
				m[x1] = true
			}
		}
	}
	return len(m)
}
