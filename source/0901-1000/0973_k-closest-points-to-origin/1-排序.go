package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 3},
		{-2, 2},
	}
	fmt.Println(kClosest(arr, 1))
}

// leetcode973_最接近原点的K个点
func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] <
			points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:K]
}
