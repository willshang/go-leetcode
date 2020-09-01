package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(pathWithObstacles(arr))
}

// 程序员面试金典08.02_迷路的机器人
var res [][]int

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	res = make([][]int, 0)
	path := make([][]int, 0)
	path = append(path, []int{0, 0})
	dfs(obstacleGrid, path)
	return res
}

func dfs(arr [][]int, path [][]int) {
	if len(res) == 0 {
		x, y := path[len(path)-1][0], path[len(path)-1][1]
		if arr[x][y] == 0 {
			arr[x][y] = 1
			if x < len(arr)-1 {
				dfs(arr, append(path, []int{x + 1, y}))
			}
			if y < len(arr[0])-1 {
				dfs(arr, append(path, []int{x, y + 1}))
			}
			if x == len(arr)-1 && y == len(arr[0])-1 {
				res = make([][]int, len(path))
				copy(res, path)
			}
		}
	}
}
