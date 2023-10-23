package main

import "fmt"

func main() {
	arr := [][]byte{
		{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '*', '#', '*'},
	}
	fmt.Println(rotateTheBox(arr))
}

func rotateTheBox(box [][]byte) [][]byte {
	n, m := len(box), len(box[0])
	res := make([][]byte, m)
	for i := 0; i < m; i++ {
		res[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			res[i][j] = '.'
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := 0 // 统计一行中#的个数
			for ; j < m && box[i][j] != '*'; j++ {
				if box[i][j] == '#' {
					count++
				}
			}
			if j < m {
				res[j][n-1-i] = '*' // 填充*
			}
			// 移动# => 填充#
			for k := j - 1; count > 0; k-- {
				res[k][n-1-i] = '#'
				count--
			}
		}
	}
	return res
}
