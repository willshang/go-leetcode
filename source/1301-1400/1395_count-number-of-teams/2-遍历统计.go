package main

import "fmt"

func main() {
	//fmt.Println(numTeams([]int{2, 5, 3, 4, 1}))
	fmt.Println(numTeams([]int{2, 1, 3}))
}

func numTeams(rating []int) int {
	res := 0
	n := len(rating)
	for i := 0; i < n; i++ {
		leftMax, leftMin := 0, 0
		rightMax, rightMin := 0, 0
		for j := 0; j < i; j++ {
			if rating[j] > rating[i] {
				leftMax++
			}
			if rating[j] < rating[i] {
				leftMin++
			}
		}
		for j := i + 1; j < len(rating); j++ {
			if rating[j] > rating[i] {
				rightMin++
			}
			if rating[j] < rating[i] {
				rightMax++
			}
		}
		res = res + leftMin*rightMin + leftMax*rightMax
	}
	return res
}
