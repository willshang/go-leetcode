package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

// leetcode1046_最后一块石头的重量
func lastStoneWeight(stones []int) int {
	length := len(stones)
	if length == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	for stones[length-2] != 0 {
		stones[length-1] = stones[length-1] - stones[length-2]
		stones[length-2] = 0
		sort.Ints(stones)
	}
	return stones[length-1]
}
