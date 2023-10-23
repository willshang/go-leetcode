package main

func main() {

}

// leetcode1926_迷宫中离入口最近的出口
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func nearestExit(maze [][]byte, entrance []int) int {
	n, m := len(maze), len(maze[0])
	queue := make([][2]int, 0)
	visited := make(map[[2]int]bool)
	queue = append(queue, [2]int{entrance[0], entrance[1]})
	visited[[2]int{entrance[0], entrance[1]}] = true
	count := 0
	for len(queue) > 0 {
		count++
		length := len(queue)
		for i := 0; i < length; i++ {
			a, b := queue[i][0], queue[i][1]
			for j := 0; j < 4; j++ {
				x := a + dx[j]
				y := b + dy[j]
				if 0 <= x && x < n && 0 <= y && y < m &&
					maze[x][y] != '+' && visited[[2]int{x, y}] == false {
					if (x == 0 || x == n-1 || y == 0 || y == m-1) && maze[x][y] == '.' {
						return count
					}
					queue = append(queue, [2]int{x, y})
					visited[[2]int{x, y}] = true
				}
			}
		}
		queue = queue[length:]
	}
	return -1
}
