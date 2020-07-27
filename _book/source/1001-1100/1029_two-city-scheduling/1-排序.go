package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))
}

// leetcode1029_两地调度
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	res := 0
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			res = res + costs[i][0]
		} else {
			res = res + costs[i][1]
		}
	}
	return res
}
