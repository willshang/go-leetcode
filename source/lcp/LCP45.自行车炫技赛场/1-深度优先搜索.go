package main

import "sort"

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var res [][]int
var visited map[[3]int]bool
var n, m int
var originX, originY int

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	res = make([][]int, 0)
	n, m = len(terrain), len(terrain[0])
	originX, originY = position[0], position[1]
	visited = make(map[[3]int]bool)
	visited[[3]int{originX, originY, 1}] = true
	dfs(terrain, obstacle, originX, originY, 1)
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		}
		return res[i][0] < res[j][0]
	})
	return res
}

func dfs(terrain [][]int, obstacle [][]int, i, j int, speed int) {
	if speed == 1 && (i == originX && j == originY) == false {
		res = append(res, []int{i, j})
	}
	for k := 0; k < 4; k++ {
		x, y := i+dx[k], j+dy[k]
		if 0 <= x && x < n && 0 <= y && y < m {
			// next = speed + h1-h2-o2
			next := speed + (terrain[i][j] - terrain[x][y] - obstacle[x][y])
			if next > 0 && visited[[3]int{x, y, next}] == false {
				visited[[3]int{x, y, next}] = true
				dfs(terrain, obstacle, x, y, next)
			}
		}
	}
}
