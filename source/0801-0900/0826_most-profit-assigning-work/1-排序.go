package main

import "sort"

func main() {

}

// leetcode826_安排工作以达到最大收益
type Node struct {
	difficulty int
	profit     int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	arr := make([]Node, 0)
	for i := 0; i < len(difficulty); i++ {
		arr = append(arr, Node{
			difficulty: difficulty[i],
			profit:     profit[i],
		})
	}
	sort.Ints(worker)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].difficulty < arr[j].difficulty
	})
	res := 0
	index := 0
	maxProfit := 0
	for i := 0; i < len(worker); i++ {
		// 找到工人收益最大
		for index < len(arr) && worker[i] >= arr[index].difficulty {
			maxProfit = max(maxProfit, arr[index].profit)
			index++
		}
		res = res + maxProfit
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
