package main

import "sort"

func main() {

}

func maxRunTime(n int, batteries []int) int64 {
	sort.Slice(batteries, func(i, j int) bool {
		return batteries[i] > batteries[j]
	})
	sum := 0
	for i := 0; i < len(batteries); i++ {
		sum = sum + batteries[i]
	}
	for i := 0; i < len(batteries); i++ {
		if batteries[i] <= sum/n { // 不够用，返回
			return int64(sum / n)
		}
		sum = sum - batteries[i] // 去除该电池
		n--
	}
	return 0
}
