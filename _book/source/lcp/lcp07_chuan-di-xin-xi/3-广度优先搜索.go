package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 2},
		{2, 1},
		{3, 4},
		{2, 3},
		{1, 4},
		{2, 0},
		{0, 4},
	}
	fmt.Println(numWays(5, arr, 3))
}

func numWays(n int, relation [][]int, k int) int {
	m := make(map[int][]int)
	for i := 0; i < len(relation); i++ {
		m[relation[i][0]] = append(m[relation[i][0]], relation[i][1])
	}
	start := []int{0}
	for i := 0; i < k; i++ {
		arr := make([]int, 0)
		for j := 0; j < len(start); j++ {
			for k := 0; k < len(m[start[j]]); k++ {
				arr = append(arr, m[start[j]][k])
			}
		}
		start = arr
	}
	res := 0
	for i := 0; i < len(start); i++ {
		if start[i] == n-1 {
			res++
		}
	}
	return res
}
