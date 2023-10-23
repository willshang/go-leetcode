package main

import "sort"

func main() {

}

func maxSatisfaction(satisfaction []int) int {
	res := 0
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	sum := 0
	temp := 0
	for i := 0; i < len(satisfaction); i++ {
		sum = sum + satisfaction[i]
		temp = temp + sum // 每多遍历1次，前后相邻的2个数，较大数相对较小数多+1次
		if temp > res {
			res = temp
		}
	}
	return res
}
