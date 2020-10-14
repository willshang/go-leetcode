package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
}

// leetcode881_救生艇
func numRescueBoats(people []int, limit int) int {
	res := 0
	sort.Ints(people)
	i, j := 0, len(people)-1
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
			j--
		} else {
			j--
		}
		res++
	}
	return res
}
