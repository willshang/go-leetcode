package main

import "fmt"

func main() {
	//fmt.Println(maximalNetworkRank(8, [][]int{
	//	{0, 1},
	//	{1, 2},
	//	{2, 3},
	//	{2, 4},
	//	{5, 6},
	//	{5, 7},
	//}))

	fmt.Println(maximalNetworkRank(4, [][]int{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 3},
	}))
}

// leetcode1615_最大网络秩
func maximalNetworkRank(n int, roads [][]int) int {
	arr := make([]int, n)
	m := make(map[int]int)
	for i := 0; i < len(roads); i++ {
		a, b := roads[i][0], roads[i][1]
		arr[a]++
		arr[b]++
		m[a*100+b]++
		m[b*100+a]++
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			value := arr[i] + arr[j] - m[i*100+j]
			if value > res {
				res = value
			}
		}
	}
	return res
}
