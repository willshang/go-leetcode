package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumWhiteTiles([][]int{
		{1, 5},
		{10, 11},
		{12, 18},
		{20, 25},
		{30, 32},
	}, 10))
}

// leetcode2271_毯子覆盖的最多白色砖块数
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	res := 0
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	count := 0
	left := 0
	for right := 0; right < len(tiles); right++ {
		l, r := tiles[right][0], tiles[right][1]
		count = count + r - l + 1
		for tiles[left][1]+carpetLen-1 < r { // 左边瓷砖右边界+毯子 无法覆盖右边瓷砖右边界：left指针右移
			count = count - (tiles[left][1] - tiles[left][0] + 1)
			left++
		}
		leftNum := r - tiles[left][0] - carpetLen + 1 // 未覆盖的数量：左边不在瓷砖内部的数量
		if leftNum < 0 {                              // 毯子够长：毯子左侧有多余
			leftNum = 0
		}
		res = max(res, count-leftNum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
