package main

import "sort"

func main() {

}

// leetcode1686_石子游戏VI
func stoneGameVI(aliceValues []int, bobValues []int) int {
	arr := make([][2]int, len(aliceValues))
	for i := 0; i < len(arr); i++ {
		arr[i] = [2]int{i, aliceValues[i] + bobValues[i]}
	}
	// 贪心策略：将两组石头的价值相加，每次取价值最大的那一组
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	a, b := 0, 0
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			a = a + aliceValues[arr[i][0]]
		} else {
			b = b + bobValues[arr[i][0]]
		}
	}
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}
