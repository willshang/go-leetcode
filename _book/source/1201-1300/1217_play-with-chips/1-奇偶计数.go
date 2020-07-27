package main

import "fmt"

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))
}

// leetcode1217_玩筹码
/*
1、所有偶数移动到同一偶数位置，花费0
2、所有奇数移动到同一奇数位置，花费0
3、将较小移动到较多的位置。
*/
func minCostToMoveChips(chips []int) int {
	odd := 0
	even := 0
	for i := 0; i < len(chips); i++ {
		if chips[i]%2 == 1 {
			odd++
		} else {
			even++
		}
	}
	if odd > even {
		return even
	}
	return odd
}
