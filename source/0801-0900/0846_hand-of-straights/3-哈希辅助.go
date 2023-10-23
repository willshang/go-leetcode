package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	//fmt.Println(isNStraightHand([]int{1, 2, 3}, 1))
	//fmt.Println(isNStraightHand([]int{8, 10, 12}, 3))
}

func isNStraightHand(hand []int, W int) bool {
	n := len(hand)
	if n%W != 0 {
		return false
	}
	if W == 1 {
		return true
	}
	m := make(map[int]int)
	for i := 0; i < len(hand); i++ {
		m[hand[i]]++
	}
	sort.Ints(hand)
	for i := 0; i < len(hand); i++ {
		value := m[hand[i]]
		if value > 0 {
			for j := 0; j < W; j++ {
				if m[hand[i]+j] < value {
					return false
				}
				m[hand[i]+j] = m[hand[i]+j] - value
			}
		}
	}
	return true
}
