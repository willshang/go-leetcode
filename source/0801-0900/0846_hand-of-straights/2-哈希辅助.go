package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	//fmt.Println(isNStraightHand([]int{1, 2, 3}, 1))
	//fmt.Println(isNStraightHand([]int{8, 10, 12}, 3))
}

// leetcode846_一手顺子
func isNStraightHand(hand []int, W int) bool {
	n := len(hand)
	if n%W != 0 {
		return false
	}
	if W == 1 {
		return true
	}
	arr := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(hand); i++ {
		if m[hand[i]] == 0 {
			arr = append(arr, hand[i])
		}
		m[hand[i]]++
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if m[arr[i]] > 0 {
			for j := 1; j < W; j++ {
				value := arr[i] + j
				m[value] = m[value] - m[arr[i]]
				if m[value] < 0 {
					return false
				}
			}
		}
	}
	return true
}
