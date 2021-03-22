package main

import "fmt"

func main() {
	fmt.Println(numTeams([]int{2, 5, 3, 4, 1}))
}

// leetcode1395_统计作战单位数
func numTeams(rating []int) int {
	res := 0
	n := len(rating)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (rating[i] < rating[j] && rating[j] < rating[k]) ||
					(rating[i] > rating[j] && rating[j] > rating[k]) {
					res++
				}
			}
		}
	}
	return res
}
