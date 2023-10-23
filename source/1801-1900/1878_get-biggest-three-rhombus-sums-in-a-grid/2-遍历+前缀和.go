package main

import "sort"

func main() {

}

// leetcode1878_矩阵中最大的三个菱形和
func getBiggestThree(grid [][]int) []int {
	arr := make([]int, 0)
	n, m := len(grid), len(grid[0])
	sum1, sum2 := make([][]int, n+2), make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		sum1[i], sum2[i] = make([]int, m+2), make([]int, m+2)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sum1[i][j] = sum1[i-1][j-1] + grid[i-1][j-1] // 下到右的斜线
			sum2[i][j] = sum2[i-1][j+1] + grid[i-1][j-1] // 下到左的斜线
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr = append(arr, grid[i][j])      // 单独一个
			for k := i + 2; k < n; k = k + 2 { // 上下长度
				aX, aY := i, j               // 上顶点
				bX, bY := k, j               // 下顶点
				cX, cY := (i+k)/2, j-(k-i)/2 // 左顶点
				dX, dY := (i+k)/2, j+(k-i)/2 // 右顶点
				if cY < 0 || dY >= m {
					break
				}
				sum := (sum2[cX+1][cY+1] - sum2[aX+1-1][aY+1+1]) + // 左到上
					(sum2[bX+1][bY+1] - sum2[dX+1-1][dY+1+1]) + // 下到右
					(sum1[bX+1][bY+1] - sum1[cX+1-1][cY+1-1]) + // 下到左
					(sum1[dX+1][dY+1] - sum1[aX+1-1][aY+1-1]) - // 右到上
					(grid[aX][aY] + grid[bX][bY] + grid[cX][cY] + grid[dX][dY])
				arr = append(arr, sum)
			}
		}
	}
	sort.Ints(arr)
	res := make([]int, 0)
	res = append(res, arr[len(arr)-1])
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] != arr[i+1] && len(res) < 3 {
			res = append(res, arr[i])
		}
	}
	return res
}
