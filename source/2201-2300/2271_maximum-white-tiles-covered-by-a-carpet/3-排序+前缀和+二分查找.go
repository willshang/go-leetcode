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
	n := len(tiles)
	arr := make([]int, n+1)

	for i := 0; i < n; i++ {
		v := tiles[i][1] - tiles[i][0] + 1
		arr[i+1] = arr[i] + v
	}
	for i := 0; i < n; i++ {
		right := tiles[i][0] + carpetLen - 1       // 右边界
		index := sort.Search(n, func(j int) bool { // 右边界>=right的下标
			return tiles[j][1] >= right
		})
		if index >= n {
			res = max(res, arr[n]-arr[i])
		} else {
			rightNum := right - tiles[index][0] + 1 // 右边覆盖的数量：右边在瓷砖内部的数量
			if rightNum <= 0 {                      // 毯子不够长，覆盖不到右边
				rightNum = 0
			}
			res = max(res, arr[index]-arr[i]+rightNum)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
