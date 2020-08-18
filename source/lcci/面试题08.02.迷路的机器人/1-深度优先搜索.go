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

var res [][]int

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	res = make([][]int, 0)
	//dfs(obstacleGrid)
	return res
}

func dfs() {

}
