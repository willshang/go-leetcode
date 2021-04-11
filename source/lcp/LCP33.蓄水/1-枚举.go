package main

import "math"

func main() {

}

// leetcode_lcp33_蓄水
func storeWater(bucket []int, vat []int) int {
	n := len(vat)
	maxValue := 0
	for i := 0; i < n; i++ {
		maxValue = max(maxValue, vat[i])
	}
	if maxValue == 0 {
		return 0
	}
	res := math.MaxInt32
	for k := 1; k <= maxValue; k++ { // 枚举蓄水的次数
		temp := k
		for i := 0; i < n; i++ {
			begin := vat[i] / k // 需要升级到的目的容量
			if vat[i]%k > 0 {
				begin++
			}
			if begin > bucket[i] {
				temp = temp + begin - bucket[i] // 升级次数
			}
		}
		res = min(res, temp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
