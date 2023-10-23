package main

import "sort"

func main() {

}

// leetcode1833_雪糕的最大数量
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	for i := 0; i < len(costs); i++ {
		if costs[i] <= coins {
			coins = coins - costs[i]
		} else {
			return i
		}
	}
	return len(costs)
}
