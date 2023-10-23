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

func isNStraightHand(hand []int, W int) bool {
	n := len(hand)
	if n%W != 0 {
		return false
	}
	if W == 1 {
		return true
	}
	sort.Ints(hand)
	for i := 0; i < n; i++ {
		if hand[i] >= 0 {
			count := 1
			for j := i + 1; j < n; j++ {
				if hand[j] > hand[i]+count {
					break
				}
				if hand[j] >= 0 && hand[j] == hand[i]+count {
					hand[j] = -1
					count++
					if count == W {
						break
					}
				}
			}
			if count != W {
				return false
			}
			hand[i] = -1
		}
	}
	return true
}
