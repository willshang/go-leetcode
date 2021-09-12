package main

import (
	"math"
	"sort"
)

func main() {

}

// leetcode1889_装包裹的最小浪费空间
var mod = 1000000007

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	n := len(packages)
	res := math.MaxInt64
	for i := 0; i < len(boxes); i++ {
		temp := boxes[i]
		sort.Ints(temp)
		if temp[len(temp)-1] < packages[n-1] { // 最大箱子无法装最大包裹
			continue
		}
		sum := 0
		left := 0
		for j := 0; j < len(temp); j++ { // 枚举箱子
			// 找到第一个大于target的位置
			right := sort.SearchInts(packages, temp[j]+1) // 选择当前箱子能装下的位置
			sum = sum + (right-left)*temp[j]              // 累加：个数*箱子大小
			left = right
		}
		res = min(res, sum)
	}
	if res == math.MaxInt64 {
		return -1
	}
	for i := 0; i < n; i++ {
		res = res - packages[i]
	}
	return res % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
