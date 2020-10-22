package main

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}

// leetcode1052_爱生气的书店老板
func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	total := 0
	res := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total = total + customers[i]
		}
	}
	window := 0
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			window = window + customers[i]
		}
	}
	res = max(res, total+window)
	for i := X; i < n; i++ {
		window = window + customers[i]*grumpy[i] - customers[i-X]*grumpy[i-X]
		res = max(res, total+window)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
