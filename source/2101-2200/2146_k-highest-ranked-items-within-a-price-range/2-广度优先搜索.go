package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 2, 0, 1},
		{1, 3, 0, 1},
		{0, 2, 5, 1},
	}
	fmt.Println(highestRankedKItems(arr, []int{2, 5}, []int{0, 0}, 3))
}

// leetcode2146_价格范围内最高排名的K样物品
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	arr := make([][]int, 0) // [dis, price, x, y]
	queue := make([][]int, 0)
	queue = append(queue, []int{start[0], start[1], 0})  // [x,y,dis]
	grid[start[0]][start[1]] = -grid[start[0]][start[1]] // 置反
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x, y, dis := node[0], node[1], node[2]
		if pricing[0] <= -grid[x][y] && -grid[x][y] <= pricing[1] { // 起点判断
			arr = append(arr, []int{dis, -grid[x][y], x, y})
		}
		for i := 0; i < 4; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if 0 <= newX && newX < n && 0 <= newY && newY < m && grid[newX][newY] > 0 {
				queue = append(queue, []int{newX, newY, dis + 1})
				grid[newX][newY] = -grid[newX][newY]
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			if arr[i][1] == arr[j][1] {
				if arr[i][2] == arr[j][2] {
					return arr[i][3] < arr[j][3]
				}
				return arr[i][2] < arr[j][2]
			}
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})
	res := make([][]int, 0)
	for i := 0; i < k && i < len(arr); i++ {
		res = append(res, []int{arr[i][2], arr[i][3]})
	}
	return res
}
