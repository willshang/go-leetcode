package main

import "fmt"

func main() {
	arr := [][]int{{0,0,0,0,0,0,0,0}}
	fmt.Println(maxAreaOfIsland(arr))
}
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := range grid{
		for j := range grid[i]{
			maxArea = max(maxArea,getArea(grid,i,j))
		}
	}
	return maxArea
}

func getArea(grid [][]int, i, j int)int  {
	if grid[i][j] == 0{
		return 0
	}

	grid[i][j] = 0
	area := 1
	if i != 0{
		area = area + getArea(grid,i-1,j)
	}
	if j != 0{
		area = area + getArea(grid,i,j-1)
	}
	if i != len(grid)-1{
		area = area + getArea(grid,i+1,j)
	}
	if j != len(grid[0])-1{
		area = area + getArea(grid,i,j+1)
	}
	return area
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}