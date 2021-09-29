package main

import "sort"

func main() {

}

// leetcode_lcp45_自行车炫技赛场
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	res := make([][]int, 0)
	n, m := len(terrain), len(terrain[0])
	originX, originY := position[0], position[1]
	visited := make(map[[3]int]bool)
	visited[[3]int{originX, originY, 1}] = true
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{originX, originY, 1})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		i, j, speed := node[0], node[1], node[2]
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
					queue = append(queue, [3]int{x, y, next})
				}
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		}
		return res[i][0] < res[j][0]
	})
	return res
}
