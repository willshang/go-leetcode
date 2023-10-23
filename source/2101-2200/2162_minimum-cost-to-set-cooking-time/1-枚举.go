package main

import "math"

func main() {

}

// leetcode2162_设置时间的最少代价
func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	m, s := targetSeconds/60, targetSeconds%60
	a := cost(startAt, moveCost, pushCost, m, s)      // 目标1
	b := cost(startAt, moveCost, pushCost, m-1, s+60) // 目标2
	return min(a, b)
}

func cost(startAt int, moveCost int, pushCost int, m, s int) int {
	res := 0
	if m < 0 || m > 99 || s < 0 || s > 99 {
		return math.MaxInt32
	}
	arr := []int{m / 10, m % 10, s / 10, s % 10}
	i := 0 // 需要忽略0后的起始位置（前面0不需要输入）
	for i < 4 && arr[i] == 0 {
		i++
	}
	for j := i; j < 4; j++ { // 从i开始
		if startAt != arr[j] { // 不相同需要移动
			startAt = arr[j]
			res = res + moveCost // 移动
		}
		res = res + pushCost
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
