package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}

// leetcode1514_概率最大的路径
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	arr := make([][]Node, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], Node{index: b, Value: succProb[i]})
		arr[b] = append(arr[b], Node{index: a, Value: succProb[i]})
	}
	maxValue := make([]float64, n)
	maxValue[start] = 1.0
	queue := make([]Node, 0)
	queue = append(queue, Node{
		Value: 1.0,
		index: start,
	})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if maxValue[node.index] > node.Value {
			continue
		}
		for i := 0; i < len(arr[node.index]); i++ {
			next := arr[node.index][i]
			// 注意使用<=，即=号要过滤，不然容易超时
			if maxValue[node.index]*next.Value <= maxValue[next.index] {
				continue
			}
			maxValue[next.index] = maxValue[node.index] * next.Value
			queue = append(queue, Node{
				Value: maxValue[next.index],
				index: next.index,
			})
		}
	}
	return maxValue[end]
}

type Node struct {
	Value float64
	index int
}
