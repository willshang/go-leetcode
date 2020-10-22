package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(breakfastNumber([]int{10, 20, 5}, []int{5, 5, 2}, 15))
	fmt.Println(breakfastNumber([]int{2, 1, 1}, []int{8, 9, 5, 1}, 9))
}

// leetcode_lcp18_早餐组合
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	res := 0
	j := len(drinks) - 1
	for i := 0; i < len(staple); i++ {
		for j >= 0 && staple[i]+drinks[j] > x {
			j--
		}
		res = (res + j + 1) % 1000000007
	}
	return res
}
