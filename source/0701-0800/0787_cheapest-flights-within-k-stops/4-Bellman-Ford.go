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
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[src] = 0 // 到开始地为0
	for i := 1; i <= k+1; i++ {
		temp := make([]int, n)
		copy(temp, dis)
		for j := 0; j < len(flights); j++ {
			a, b, c := flights[j][0], flights[j][1], flights[j][2] // a=>b c
			temp[b] = min(temp[b], dis[a]+c)
		}
		dis = temp
	}
	if dis[dst] == maxValue {
		return -1
	}
	return dis[dst]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
