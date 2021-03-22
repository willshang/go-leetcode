package main

func main() {

}

// leetcode1765_地图中的最高点
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func highestPeak(isWater [][]int) [][]int {
	n, m := len(isWater), len(isWater[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	queue := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if isWater[i][j] == 1 {
				res[i][j] = -1
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	level := 0
	for len(queue) > 0 {
		level++
		length := len(queue)
		for i := 0; i < length; i++ {
			for j := 0; j < 4; j++ {
				x := queue[i][0] + dx[j]
				y := queue[i][1] + dy[j]
				if 0 <= x && x < n && 0 <= y && y < m && res[x][y] == 0 {
					res[x][y] = level
					queue = append(queue, [2]int{x, y})
				}
			}
		}
		queue = queue[length:]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if res[i][j] == -1 {
				res[i][j] = 0
			}
		}
	}
	return res
}
