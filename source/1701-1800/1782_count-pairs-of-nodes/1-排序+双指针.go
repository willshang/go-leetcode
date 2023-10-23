package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countPairs(4, [][]int{
		{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1},
	}, []int{2, 3}))
}

// leetcode1782_统计点对的数目
func countPairs(n int, edges [][]int, queries []int) []int {
	degree := make([]int, n+1) // 点=>相连边的次数
	m := make(map[[2]int]int)  // 边的出现次数
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		if a > b { // 调整为a<b
			a, b = b, a
		}
		m[[2]int{a, b}]++
		degree[a]++
		degree[b]++
	}
	arr := make([]int, n+1) // 使用临时数组来排序
	copy(arr, degree)
	sort.Ints(arr)
	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		target := queries[i]
		left, right := 1, n // 双指针=>2数之和
		for left < right {
			if target < arr[left]+arr[right] {
				res[i] = res[i] + right - left // 计算对数
				right--
			} else {
				left++
			}
		}
		for k, v := range m { // 遍历边：减去2点之间的重复边的数量后不满足要求，数目-1
			a, b := k[0], k[1]
			if target < degree[a]+degree[b] && target >= degree[a]+degree[b]-v {
				res[i]--
			}
		}
	}
	return res
}
