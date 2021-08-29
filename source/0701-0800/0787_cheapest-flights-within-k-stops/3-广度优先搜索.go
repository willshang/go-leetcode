package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice(3, arr, 0, 2, 1))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	maxValue := math.MaxInt32 / 10
	prices := make([]int, n)
	arr := make([][][2]int, n)
	for i := 0; i < n; i++ {
		prices[i] = maxValue
	}
	prices[src] = 0
	for i := 0; i < len(flights); i++ {
		a, b, c := flights[i][0], flights[i][1], flights[i][2] // a=>b c
		arr[a] = append(arr[a], [2]int{b, c})
	}
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{1, src, prices[src]}) // 次数，起点，价格
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node[0] > k+1 { // 大于k+1次退出
			break
		}
		cur, value := node[1], node[2]
		for i := 0; i < len(arr[cur]); i++ {
			b, c := arr[cur][i][0], arr[cur][i][1]
			if prices[b] > c+value {
				prices[b] = c + value
				queue = append(queue, [3]int{node[0] + 1, b, prices[b]})
			}
		}
	}
	if prices[dst] == maxValue {
		return -1
	}
	return prices[dst]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
