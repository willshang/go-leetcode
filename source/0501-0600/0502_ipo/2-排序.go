package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}

// leetcode502_IPO
type Node struct {
	profit  int
	capital int
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	arr := make([]Node, 0)
	for i := 0; i < len(Profits); i++ {
		arr = append(arr, Node{
			profit:  Profits[i],
			capital: Capital[i],
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].profit > arr[j].profit
	})
	index := 0
	for k > 0 {
		if index == len(arr) {
			return W
		}
		// 挑选一个满足条件的项目，利润最大即可
		if arr[index].capital <= W {
			k--
			W = W + arr[index].profit
			arr = append(arr[:index], arr[index+1:]...)
			index = 0
			continue
		}
		index++
	}
	return W
}
