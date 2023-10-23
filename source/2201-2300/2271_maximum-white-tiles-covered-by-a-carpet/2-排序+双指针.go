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

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	res := 0
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	count := 0
	right := 0
	n := len(tiles)
	for left := 0; left < n; left++ {
		for right < n && tiles[right][1]-tiles[left][0]+1 < carpetLen { // 左边瓷砖左边界+毯子 能覆盖右边瓷砖右边界：right指针右移
			count = count + (tiles[right][1] - tiles[right][0] + 1)
			right++
		}
		if right < n {
			rightNum := tiles[left][0] + carpetLen - tiles[right][0] // 右边覆盖的数量：右边在瓷砖内部的数量
			if rightNum < 0 {                                        // 毯子不够长，覆盖不到右边
				rightNum = 0
			}
			res = max(res, count+rightNum)
		} else {
			res = max(res, count)
		}
		count = count - (tiles[left][1] - tiles[left][0] + 1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
