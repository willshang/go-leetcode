package main

import "fmt"

func main() {
	arr := [][]int{{2,2,2},{2,1,2},{2,2,2}}
	fmt.Println(surfaceArea(arr))
}

//第一步：总面试是个数*6
//第二步：同一位置，从2层以上开始，每升高一层，减少2个面
//第三部：左右位置，每相邻一个，减少两个面
//第四步：前后位置，每相邻一个，减少2个面
func surfaceArea(grid [][]int) int {
	sum := 0
	for i, rows := range grid{
		for j, _ := range rows{
			sum = sum + grid[i][j] * 6

			if grid[i][j] > 1{
				sum = sum - (grid[i][j]-1)*2
			}

			if j > 0{
				sum = sum - (min(grid[i][j],grid[i][j-1])*2)
			}

			if i > 0{
				sum = sum - (min(grid[i][j],grid[i-1][j])*2)
			}
		}
	}
	return sum
}

func min(a, b int)int  {
	if a > b {
		return b
	}
	return a
}