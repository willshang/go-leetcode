package main

import (
	"fmt"
)

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}

// leetcode1042_不邻接植花
func gardenNoAdj(N int, paths [][]int) []int {
	res := make([]int, N+1)
	arr := make([][]int, N+1)
	for i := 0; i < len(paths); i++ {
		arr[paths[i][0]] = append(arr[paths[i][0]], paths[i][1])
		arr[paths[i][1]] = append(arr[paths[i][1]], paths[i][0])
	}
	for i := 1; i <= N; i++ {
		m := map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
		}
		for j := 0; j < len(arr[i]); j++ {
			if res[arr[i][j]] > 0 {
				delete(m, res[arr[i][j]])
			}
		}
		for k := range m {
			res[i] = k
			break
		}
	}
	return res[1:]
}
