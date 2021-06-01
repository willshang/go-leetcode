package main

import "sort"

func main() {

}

func getBiggestThree(grid [][]int) []int {
	arr := make([]int, 0)
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr = append(arr, grid[i][j]) // 单独一个
			for k := 1; k < n; k++ {      // 枚举增加的长度
				left := i - k
				right := i + k
				button := j + 2*k
				mid := j + k
				if left < 0 || right > n-1 || button > m-1 {
					break
				}
				sum := 0
				b := mid
				for a := left; a < i; a++ { // 左到上
					sum = sum + grid[a][b]
					b--
				}
				for a := i; a < right; a++ { // 上到右
					sum = sum + grid[a][b]
					b++
				}
				for a := right; a > i; a-- { // 右到下
					sum = sum + grid[a][b]
					b++
				}
				for a := i; a > left; a-- { // 下到左
					sum = sum + grid[a][b]
					b--
				}
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
